package services

import (
	"context"
	"dr-aditi-backend/internal/config"
	"dr-aditi-backend/internal/db"
	"dr-aditi-backend/internal/models"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// CalendarService handles Google Calendar integration
type CalendarService struct {
	config   *config.CalendarConfig
	service  *calendar.Service
	db       *db.Database
	location *time.Location
}

// NewCalendarService creates a new calendar service
func NewCalendarService(cfg *config.CalendarConfig, database *db.Database, timezone string) (*CalendarService, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.UTC
	}

	cs := &CalendarService{
		config:   cfg,
		db:       database,
		location: loc,
	}

	// If credentials file is provided, initialize Google Calendar
	if cfg.CredentialsFile != "" && cfg.CalendarID != "" {
		if err := cs.initializeGoogleCalendar(); err != nil {
			// Log warning but don't fail - we can work without Google Calendar
			fmt.Printf("Warning: Google Calendar not initialized: %v\n", err)
		}
	}

	return cs, nil
}

// initializeGoogleCalendar sets up the Google Calendar API client
func (cs *CalendarService) initializeGoogleCalendar() error {
	ctx := context.Background()

	creds, err := os.ReadFile(cs.config.CredentialsFile)
	if err != nil {
		return fmt.Errorf("reading credentials: %w", err)
	}

	jwtConfig, err := google.JWTConfigFromJSON(creds,
		calendar.CalendarScope,
		calendar.CalendarEventsScope,
	)
	if err != nil {
		return fmt.Errorf("parsing credentials: %w", err)
	}

	client := jwtConfig.Client(ctx)
	service, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("creating calendar service: %w", err)
	}

	// Verify calendar access
	_, err = service.Calendars.Get(cs.config.CalendarID).Do()
	if err != nil {
		return fmt.Errorf("verifying calendar access: %w", err)
	}

	cs.service = service
	fmt.Printf("Successfully connected to Google Calendar: %s\n", cs.config.CalendarID)
	return nil
}

// GetAvailableSlots returns available slots for a date range
func (cs *CalendarService) GetAvailableSlots(startDate, endDate time.Time) ([]models.DaySlots, error) {
	// Normalize to start of day in local timezone
	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, cs.location)
	end := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, cs.location)

	// Get booked slots from database
	bookedTimes, err := cs.db.GetBookedSlots(start, end)
	if err != nil {
		return nil, fmt.Errorf("getting booked slots: %w", err)
	}

	// Get blocked dates
	blockedDates, err := cs.db.GetBlockedDates()
	if err != nil {
		return nil, fmt.Errorf("getting blocked dates: %w", err)
	}

	// Create blocked dates map
	blockedMap := make(map[string]bool)
	for _, bd := range blockedDates {
		blockedMap[bd.Date] = true
	}

	// Create a map for quick lookup of booked times
	bookedMap := make(map[string]bool)
	for _, t := range bookedTimes {
		key := t.Format("2006-01-02 15:04")
		bookedMap[key] = true
	}

	// If Google Calendar is configured, also check there
	if cs.service != nil {
		gcalEvents, err := cs.getGoogleCalendarEvents(start, end)
		if err != nil {
			fmt.Printf("Warning: Could not fetch Google Calendar events: %v\n", err)
		} else {
			for _, eventTime := range gcalEvents {
				key := eventTime.Format("2006-01-02 15:04")
				bookedMap[key] = true
			}
		}
	}

	// Generate slots for each day
	var result []models.DaySlots
	current := start
	now := time.Now().In(cs.location)

	for !current.After(end) {
		dateStr := current.Format("2006-01-02")
		isBlocked := blockedMap[dateStr]

		daySlots := models.DaySlots{
			Date:      dateStr,
			IsWeekend: current.Weekday() == time.Sunday || isBlocked,
		}

		// Skip Sundays and blocked dates
		if current.Weekday() != time.Sunday && !isBlocked {
			// Generate time slots
			for hour := cs.config.WorkStartHour; hour < cs.config.WorkEndHour; hour++ {
				for _, minute := range []int{0, 30} {
					slotTime := time.Date(current.Year(), current.Month(), current.Day(),
						hour, minute, 0, 0, cs.location)

					// Skip past slots
					if slotTime.Before(now) {
						continue
					}

					// Check if slot is booked
					key := slotTime.Format("2006-01-02 15:04")
					available := !bookedMap[key]

					daySlots.Slots = append(daySlots.Slots, models.TimeSlot{
						Time:      fmt.Sprintf("%02d:%02d", hour, minute),
						Available: available,
					})
				}
			}
		}

		result = append(result, daySlots)
		current = current.AddDate(0, 0, 1)
	}

	return result, nil
}

// getGoogleCalendarEvents fetches events from Google Calendar
func (cs *CalendarService) getGoogleCalendarEvents(start, end time.Time) ([]time.Time, error) {
	events, err := cs.service.Events.List(cs.config.CalendarID).
		TimeMin(start.Format(time.RFC3339)).
		TimeMax(end.Format(time.RFC3339)).
		SingleEvents(true).
		Do()
	if err != nil {
		return nil, err
	}

	var times []time.Time
	for _, event := range events.Items {
		if event.Start.DateTime != "" {
			t, err := time.Parse(time.RFC3339, event.Start.DateTime)
			if err == nil {
				times = append(times, t)
			}
		}
	}

	return times, nil
}

// CreateCalendarEvent creates an event in Google Calendar
func (cs *CalendarService) CreateCalendarEvent(apt *models.Appointment) (string, string, error) {
	// Generate Jitsi Meet link
	roomID := strings.ReplaceAll(apt.ID, "-", "")
	meetLink := fmt.Sprintf("https://meet.jit.si/DrAditi-%s", roomID[:12])

	// If Google Calendar is not configured, just return the meet link
	if cs.service == nil {
		return "", meetLink, nil
	}

	consultationType := "Video"
	if apt.ConsultationType == "voice" {
		consultationType = "Voice"
	}

	event := &calendar.Event{
		Summary: fmt.Sprintf("%s Consultation - %s", consultationType, apt.PatientName),
		Description: fmt.Sprintf(`Patient: %s
Email: %s
Phone: %s
Consultation Type: %s
Join Meeting: %s`,
			apt.PatientName, apt.PatientEmail, apt.PatientPhone,
			consultationType, meetLink),
		Start: &calendar.EventDateTime{
			DateTime: apt.StartTime.Format(time.RFC3339),
			TimeZone: cs.location.String(),
		},
		End: &calendar.EventDateTime{
			DateTime: apt.EndTime.Format(time.RFC3339),
			TimeZone: cs.location.String(),
		},
		Attendees: []*calendar.EventAttendee{
			{Email: apt.PatientEmail},
		},
		Reminders: &calendar.EventReminders{
			UseDefault: false,
			Overrides: []*calendar.EventReminder{
				{Method: "email", Minutes: 60},
				{Method: "popup", Minutes: 15},
			},
		},
	}

	createdEvent, err := cs.service.Events.Insert(cs.config.CalendarID, event).
		SendUpdates("all").
		Do()
	if err != nil {
		return "", meetLink, fmt.Errorf("creating calendar event: %w", err)
	}

	return createdEvent.Id, meetLink, nil
}

// IsSlotAvailable checks if a specific slot is available
func (cs *CalendarService) IsSlotAvailable(slotTime time.Time) (bool, error) {
	// Check in database
	bookedSlots, err := cs.db.GetBookedSlots(slotTime, slotTime.Add(time.Minute))
	if err != nil {
		return false, err
	}

	if len(bookedSlots) > 0 {
		return false, nil
	}

	// Check in Google Calendar if configured
	if cs.service != nil {
		events, err := cs.getGoogleCalendarEvents(slotTime, slotTime.Add(time.Duration(cs.config.SlotDuration)*time.Minute))
		if err != nil {
			fmt.Printf("Warning: Could not check Google Calendar: %v\n", err)
		} else if len(events) > 0 {
			return false, nil
		}
	}

	return true, nil
}
