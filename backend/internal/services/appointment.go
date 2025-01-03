package services

import (
	"Project_PG/Backend/internal/models"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type AppointmentService struct {
	calendarService *CalendarService
}

func NewAppointmentService(calendarService *CalendarService) *AppointmentService {
	return &AppointmentService{
		calendarService: calendarService,
	}
}

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
func (s *AppointmentService) GetAppointments(start, end time.Time) ([]models.Appointment, error) {
	return []models.Appointment{}, nil
}

func (s *AppointmentService) IsTimeSlotAvailable(startTime time.Time) (bool, error) {
	return true, nil
}
