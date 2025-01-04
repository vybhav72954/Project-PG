package services

import (
	"Project_PG/Backend/internal/models"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// AppointmentService orchestrates appointment creation using calendar and email services
type AppointmentService struct {
	calendarService *CalendarService
}

// NewAppointmentService initializes the appointment service
func NewAppointmentService(calendarService *CalendarService) *AppointmentService {
	return &AppointmentService{
		calendarService: calendarService,
	}
}

// CreateAppointment generates a new appointment with UUID and creates calendar event
func (s *AppointmentService) CreateAppointment(req models.BookingRequest) (*models.Appointment, error) {

	appointmentID := uuid.New().String() // Make sure you've imported "github.com/google/uuid"

	// Create new appointment with the ID
	apt := &models.Appointment{
		ID:           appointmentID, // Set the ID here
		PatientName:  req.PatientName,
		PatientEmail: req.PatientEmail,
		StartTime:    req.StartTime,
		EndTime:      req.StartTime.Add(30 * time.Minute),
		Status:       "confirmed",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Create calendar event
	if err := s.calendarService.CreateEvent(apt); err != nil {
		return nil, fmt.Errorf("creating calendar event: %w", err)
	}

	return apt, nil
}

// GetAppointments and IsTimeSlotAvailable are placeholder methods for future implementation
func (s *AppointmentService) GetAppointments(start, end time.Time) ([]models.Appointment, error) {
	return []models.Appointment{}, nil
}

func (s *AppointmentService) IsTimeSlotAvailable(startTime time.Time) (bool, error) {
	return true, nil
}
