package models

import "time"

type Appointment struct {
	ID           string    `json:"id" db:"id"`
	EventID      string    `json:"event_id" db:"event_id"` // Google Calendar Event ID
	PatientName  string    `json:"patient_name" db:"patient_name"`
	PatientEmail string    `json:"patient_email" db:"patient_email"`
	StartTime    time.Time `json:"start_time" db:"start_time"`
	EndTime      time.Time `json:"end_time" db:"end_time"`
	Status       string    `json:"status" db:"status"`       // confirmed, cancelled
	MeetLink     string    `json:"meet_link" db:"meet_link"` // Google Meet link
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type BookingRequest struct {
	PatientName  string    `json:"patient_name" validate:"required"`
	PatientEmail string    `json:"patient_email" validate:"required,email"`
	StartTime    time.Time `json:"start_time" validate:"required"`
}

type TimeSlot struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Available bool      `json:"available"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
