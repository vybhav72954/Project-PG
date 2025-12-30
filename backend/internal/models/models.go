package models

import "time"

// Patient represents a registered patient
type Patient struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Appointment represents a booked consultation
type Appointment struct {
	ID               string    `json:"id" db:"id"`
	PatientID        string    `json:"patient_id" db:"patient_id"`
	PatientName      string    `json:"patient_name" db:"patient_name"`
	PatientEmail     string    `json:"patient_email" db:"patient_email"`
	PatientPhone     string    `json:"patient_phone" db:"patient_phone"`
	ConsultationType string    `json:"consultation_type" db:"consultation_type"` // video or voice
	StartTime        time.Time `json:"start_time" db:"start_time"`
	EndTime          time.Time `json:"end_time" db:"end_time"`
	Status           string    `json:"status" db:"status"` // pending_payment, confirmed, cancelled, completed
	PaymentStatus    string    `json:"payment_status" db:"payment_status"`
	PaymentID        string    `json:"payment_id" db:"payment_id"`
	PaymentOrderID   string    `json:"payment_order_id" db:"payment_order_id"`
	Amount           int       `json:"amount" db:"amount"` // in paise
	MeetLink         string    `json:"meet_link" db:"meet_link"`
	CalendarEventID  string    `json:"calendar_event_id" db:"calendar_event_id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

// BookingRequest is the incoming request to book an appointment
type BookingRequest struct {
	Name             string `json:"name" validate:"required,min=2"`
	Email            string `json:"email" validate:"required,email"`
	Phone            string `json:"phone" validate:"required,min=10"`
	ConsultationType string `json:"consultation_type" validate:"required,oneof=video voice"`
	Date             string `json:"date" validate:"required"`      // YYYY-MM-DD format
	TimeSlot         string `json:"time_slot" validate:"required"` // HH:MM format
}

// TimeSlot represents an available time slot
type TimeSlot struct {
	Time      string `json:"time"` // "09:00", "10:00", etc.
	Available bool   `json:"available"`
}

// DaySlots represents slots for a specific day
type DaySlots struct {
	Date      string     `json:"date"` // YYYY-MM-DD
	Slots     []TimeSlot `json:"slots"`
	IsWeekend bool       `json:"is_weekend"`
}

// PaymentOrder represents Razorpay order creation response
type PaymentOrder struct {
	OrderID       string `json:"order_id"`
	Amount        int    `json:"amount"` // in paise
	Currency      string `json:"currency"`
	AppointmentID string `json:"appointment_id"`
}

// PaymentVerification is the request to verify payment
type PaymentVerification struct {
	RazorpayOrderID   string `json:"razorpay_order_id"`
	RazorpayPaymentID string `json:"razorpay_payment_id"`
	RazorpaySignature string `json:"razorpay_signature"`
	AppointmentID     string `json:"appointment_id"`
}

// Testimonial represents a patient review
type Testimonial struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"` // First name only
	Review    string    `json:"review" db:"review"`
	Rating    int       `json:"rating" db:"rating"`       // 1-5 stars
	Condition string    `json:"condition" db:"condition"` // What they were treated for
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// APIResponse is the standard API response wrapper
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// SlotsRequest is the request to get available slots
type SlotsRequest struct {
	StartDate string `json:"start_date"` // YYYY-MM-DD
	EndDate   string `json:"end_date"`   // YYYY-MM-DD
}

// ConsultationPricing holds the prices for different consultation types
type ConsultationPricing struct {
	Video int `json:"video"` // in paise
	Voice int `json:"voice"` // in paise
}

// Config for pricing
var Pricing = ConsultationPricing{
	Video: 99900, // ₹999
	Voice: 79900, // ₹799
}

// ============================================================
// ADMIN MODELS
// ============================================================

// DashboardStats holds admin dashboard statistics
type DashboardStats struct {
	TodayAppointments    int `json:"today_appointments"`
	WeekAppointments     int `json:"week_appointments"`
	UpcomingAppointments int `json:"upcoming_appointments"`
	PendingAppointments  int `json:"pending_appointments"`
	TotalPatients        int `json:"total_patients"`
	MonthRevenue         int `json:"month_revenue"` // in paise
}

// PatientSummary is a summary of patient info for admin
type PatientSummary struct {
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	Phone            string    `json:"phone"`
	AppointmentCount int       `json:"appointment_count"`
	LastAppointment  time.Time `json:"last_appointment"`
	FirstVisit       time.Time `json:"first_visit"`
}

// BlockedDate represents a date when appointments are not available
type BlockedDate struct {
	ID        string    `json:"id"`
	Date      string    `json:"date"` // YYYY-MM-DD
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}

// AdminLoginRequest is the request for admin login
type AdminLoginRequest struct {
	Password string `json:"password"`
}

// TestimonialRequest for creating/updating testimonials
type TestimonialRequest struct {
	Name      string `json:"name"`
	Review    string `json:"review"`
	Rating    int    `json:"rating"`
	Condition string `json:"condition"`
	IsActive  bool   `json:"is_active"`
}

// BlockDateRequest for blocking a date
type BlockDateRequest struct {
	Date   string `json:"date"` // YYYY-MM-DD
	Reason string `json:"reason"`
}

// UpdateStatusRequest for updating appointment status
type UpdateStatusRequest struct {
	Status string `json:"status"` // confirmed, cancelled, completed
}

type UpdateNotesRequest struct {
	Notes string `json:"notes"`
}
