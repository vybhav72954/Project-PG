package models

import "time"

// Appointment represents a scheduled consultation
type Appointment struct {
	ID           string    `json:"id" db:"id"`                     // Unique identifier for appointment
	EventID      string    `json:"event_id" db:"event_id"`         // Google Calendar Event ID
	PatientName  string    `json:"patient_name" db:"patient_name"` // Name of the patient
	PatientEmail string    `json:"patient_email" db:"patient_email"`
	StartTime    time.Time `json:"start_time" db:"start_time"` // Appointment start time
	EndTime      time.Time `json:"end_time" db:"end_time"`     // Appointment end time
	Status       string    `json:"status" db:"status"`         // Current status (confirmed, cancelled)
	MeetLink     string    `json:"meet_link" db:"meet_link"`   // Jitsi Meet link for consultation
	CreatedAt    time.Time `json:"created_at" db:"created_at"` // Record creation timestamp
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"` // Last update timestamp
}

// BookingRequest represents the incoming appointment booking request
type BookingRequest struct {
	PatientName  string    `json:"patient_name" validate:"required"`
	PatientEmail string    `json:"patient_email" validate:"required,email"`
	StartTime    time.Time `json:"start_time" validate:"required"`
}

// TimeSlot represents an available appointment time slot
type TimeSlot struct {
	StartTime time.Time `json:"start_time"` // Slot start time
	EndTime   time.Time `json:"end_time"`   // Slot end time
	Available bool      `json:"available"`  // Whether slot is available for booking
}

// APIResponse represents the standard API response format
type APIResponse struct {
	Success bool        `json:"success"`         // Whether the request was successful
	Data    interface{} `json:"data,omitempty"`  // Response data (if any)
	Error   string      `json:"error,omitempty"` // Error message (if any)
}
