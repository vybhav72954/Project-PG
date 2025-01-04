package services

import (
	"Project_PG/Backend/internal/config"
	"Project_PG/Backend/internal/models"
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
	"os"
	"strings"
	"time"
)

// CalendarService handles Google Calendar integration and event management
type CalendarService struct {
	config  *config.CalendarConfig
	service *calendar.Service
}

// NewCalendarService initializes calendar service with Google credentials
func NewCalendarService(config *config.CalendarConfig) (*CalendarService, error) {
	ctx := context.Background()

	creds, err := os.ReadFile(config.CredentialsFile)
	if err != nil {
		return nil, fmt.Errorf("reading credentials: %w", err)
	}

	jwtConfig, err := google.JWTConfigFromJSON(creds,
		calendar.CalendarScope,
		calendar.CalendarEventsScope,
		"https://www.googleapis.com/auth/calendar",
		"https://www.googleapis.com/auth/calendar.events",
	)
	if err != nil {
		return nil, fmt.Errorf("parsing credentials: %w", err)
	}

	service, err := calendar.New(jwtConfig.Client(ctx))
	if err != nil {
		return nil, fmt.Errorf("creating calendar service: %w", err)
	}

	// Verify calendar access
	_, err = service.Calendars.Get(config.CalendarID).Do()
	if err != nil {
		return nil, fmt.Errorf("verifying calendar access: %w", err)
	}

	fmt.Printf("Successfully connected to calendar: %s\n", config.CalendarID)
	return &CalendarService{
		config:  config,
		service: service,
	}, nil
}

// CreateEvent creates a calendar event and generates a Jitsi meet link
func (s *CalendarService) CreateEvent(apt *models.Appointment) error {
	// Generate Jitsi Meet room ID based on appointment ID
	roomID := strings.ReplaceAll(apt.ID, "-", "")
	meetLink := fmt.Sprintf("https://meet.jit.si/%s", roomID)

	event := &calendar.Event{
		Summary: fmt.Sprintf("Consultation with %s", apt.PatientName),
		Description: fmt.Sprintf("Online homeopathy consultation\nPatient Email: %s\nJoin meeting: %s",
			apt.PatientEmail, meetLink),
		Start: &calendar.EventDateTime{
			DateTime: apt.StartTime.Format(time.RFC3339),
			TimeZone: "Asia/Kolkata",
		},
		End: &calendar.EventDateTime{
			DateTime: apt.EndTime.Format(time.RFC3339),
			TimeZone: "Asia/Kolkata",
		},
	}

	event, err := s.service.Events.Insert(s.config.CalendarID, event).Do()
	if err != nil {
		return fmt.Errorf("creating calendar event: %w", err)
	}

	apt.EventID = event.Id
	apt.MeetLink = meetLink
	return nil
}

// GetAvailableSlots returns available 30-minute slots between 9 AM to 5 PM
// Considers existing appointments to determine availability - Further Resting Required
func (s *CalendarService) GetAvailableSlots(start, end time.Time) ([]models.TimeSlot, error) {
	fmt.Printf("Fetching events between %v and %v\n", start, end)

	events, err := s.service.Events.List(s.config.CalendarID).
		TimeMin(start.Format(time.RFC3339)).
		TimeMax(end.Format(time.RFC3339)).
		Do()
	if err != nil {
		return nil, fmt.Errorf("fetching events: %w", err)
	}

	fmt.Printf("Found %d existing events\n", len(events.Items))

	var availableSlots []models.TimeSlot
	currentDay := start

	for currentDay.Before(end) {
		dayStart := time.Date(currentDay.Year(), currentDay.Month(), currentDay.Day(), 9, 0, 0, 0, currentDay.Location())
		dayEnd := time.Date(currentDay.Year(), currentDay.Month(), currentDay.Day(), 17, 0, 0, 0, currentDay.Location())

		for slotStart := dayStart; slotStart.Before(dayEnd); slotStart = slotStart.Add(30 * time.Minute) {
			slotEnd := slotStart.Add(30 * time.Minute)
			isAvailable := true

			for _, event := range events.Items {
				eventStart, _ := time.Parse(time.RFC3339, event.Start.DateTime)
				eventEnd, _ := time.Parse(time.RFC3339, event.End.DateTime)

				if (slotStart.Equal(eventStart) || slotStart.After(eventStart)) &&
					slotStart.Before(eventEnd) {
					isAvailable = false
					break
				}
			}

			if isAvailable {
				availableSlots = append(availableSlots, models.TimeSlot{
					StartTime: slotStart,
					EndTime:   slotEnd,
					Available: true,
				})
			}
		}

		currentDay = currentDay.AddDate(0, 0, 1)
	}

	fmt.Printf("Generated %d available slots\n", len(availableSlots))
	return availableSlots, nil
}
