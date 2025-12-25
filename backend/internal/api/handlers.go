package api

import (
	"dr-aditi-backend/internal/config"
	"dr-aditi-backend/internal/db"
	"dr-aditi-backend/internal/models"
	"dr-aditi-backend/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Handler contains all dependencies for HTTP handlers
type Handler struct {
	db       *db.Database
	calendar *services.CalendarService
	email    *services.EmailService
	payment  *services.PaymentService
	config   *config.Config
	location *time.Location
}

// NewHandler creates a new handler with all dependencies
func NewHandler(
	database *db.Database,
	calendar *services.CalendarService,
	email *services.EmailService,
	payment *services.PaymentService,
	cfg *config.Config,
) *Handler {
	loc, _ := time.LoadLocation(cfg.App.TimeZone)
	if loc == nil {
		loc = time.UTC
	}

	return &Handler{
		db:       database,
		calendar: calendar,
		email:    email,
		payment:  payment,
		config:   cfg,
		location: loc,
	}
}

// HealthCheck returns server health status
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
		},
	})
}

// GetConfig returns public configuration for the frontend
func (h *Handler) GetConfig(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"razorpay_key_id": h.payment.GetKeyID(),
			"pricing": map[string]int{
				"video": models.Pricing.Video,
				"voice": models.Pricing.Voice,
			},
			"clinic": map[string]string{
				"name":     h.config.App.ClinicName,
				"doctor":   h.config.App.DoctorName,
				"phone":    h.config.App.ClinicPhone,
				"whatsapp": h.config.App.ClinicWhatsApp,
				"email":    h.config.App.ClinicEmail,
			},
		},
	})
}

// GetAvailableSlots returns available appointment slots
func (h *Handler) GetAvailableSlots(w http.ResponseWriter, r *http.Request) {
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	if startDateStr == "" {
		startDateStr = time.Now().Format("2006-01-02")
	}
	if endDateStr == "" {
		// Default to 14 days from start
		startDate, _ := time.Parse("2006-01-02", startDateStr)
		endDateStr = startDate.AddDate(0, 0, 14).Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid start_date format. Use YYYY-MM-DD")
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid end_date format. Use YYYY-MM-DD")
		return
	}

	slots, err := h.calendar.GetAvailableSlots(startDate, endDate)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to get available slots")
		return
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    slots,
	})
}

// CreateBooking creates a new appointment and returns payment order
func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req models.BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if req.Name == "" || req.Email == "" || req.Phone == "" {
		respondError(w, http.StatusBadRequest, "Name, email, and phone are required")
		return
	}
	if req.ConsultationType != "video" && req.ConsultationType != "voice" {
		respondError(w, http.StatusBadRequest, "Invalid consultation type")
		return
	}
	if req.Date == "" || req.TimeSlot == "" {
		respondError(w, http.StatusBadRequest, "Date and time slot are required")
		return
	}

	// Parse appointment time
	dateTimeStr := fmt.Sprintf("%s %s", req.Date, req.TimeSlot)
	startTime, err := time.ParseInLocation("2006-01-02 15:04", dateTimeStr, h.location)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid date or time format")
		return
	}

	// Check if slot is still available
	available, err := h.calendar.IsSlotAvailable(startTime)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to check slot availability")
		return
	}
	if !available {
		respondError(w, http.StatusConflict, "This time slot is no longer available")
		return
	}

	// Determine price
	amount := models.Pricing.Video
	if req.ConsultationType == "voice" {
		amount = models.Pricing.Voice
	}

	// Create appointment
	apt := &models.Appointment{
		ID:               uuid.New().String(),
		PatientName:      req.Name,
		PatientEmail:     req.Email,
		PatientPhone:     req.Phone,
		ConsultationType: req.ConsultationType,
		StartTime:        startTime,
		EndTime:          startTime.Add(30 * time.Minute),
		Status:           "pending_payment",
		PaymentStatus:    "pending",
		Amount:           amount,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	// Create Razorpay order
	order, err := h.payment.CreateOrder(apt)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create payment order")
		return
	}

	apt.PaymentOrderID = order.OrderID

	// Save appointment to database
	if err := h.db.CreateAppointment(apt); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create appointment")
		return
	}

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"appointment_id":   apt.ID,
			"order_id":         order.OrderID,
			"amount":           order.Amount,
			"currency":         order.Currency,
			"razorpay_key_id":  h.payment.GetKeyID(),
			"patient_name":     apt.PatientName,
			"patient_email":    apt.PatientEmail,
			"patient_phone":    apt.PatientPhone,
			"appointment_time": apt.StartTime.Format(time.RFC3339),
		},
	})
}

// VerifyPayment verifies Razorpay payment and confirms appointment
func (h *Handler) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	var verification models.PaymentVerification
	if err := json.NewDecoder(r.Body).Decode(&verification); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Get appointment
	apt, err := h.db.GetAppointmentByOrderID(verification.RazorpayOrderID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if apt == nil {
		respondError(w, http.StatusNotFound, "Appointment not found")
		return
	}

	// Verify payment signature
	verified, err := h.payment.VerifyPayment(&verification)
	if err != nil || !verified {
		respondError(w, http.StatusBadRequest, "Payment verification failed")
		return
	}

	// Create calendar event and get meet link
	calendarEventID, meetLink, err := h.calendar.CreateCalendarEvent(apt)
	if err != nil {
		fmt.Printf("Warning: Failed to create calendar event: %v\n", err)
		// Continue anyway - the appointment is still valid
		meetLink = fmt.Sprintf("https://meet.jit.si/DrAditi-%s", apt.ID[:12])
	}

	// Update appointment in database
	if err := h.db.UpdateAppointmentPayment(
		apt.ID,
		verification.RazorpayPaymentID,
		meetLink,
		calendarEventID,
	); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update appointment")
		return
	}

	// Get updated appointment
	apt, _ = h.db.GetAppointment(apt.ID)

	// Send confirmation email asynchronously
	go func() {
		if err := h.email.SendAppointmentConfirmation(apt); err != nil {
			fmt.Printf("Failed to send confirmation email: %v\n", err)
		}
	}()

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Payment verified and appointment confirmed",
		Data: map[string]interface{}{
			"appointment_id":   apt.ID,
			"status":           "confirmed",
			"meet_link":        apt.MeetLink,
			"appointment_time": apt.StartTime.Format(time.RFC3339),
		},
	})
}

// GetAppointment returns appointment details
func (h *Handler) GetAppointment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		respondError(w, http.StatusBadRequest, "Appointment ID is required")
		return
	}

	apt, err := h.db.GetAppointment(id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if apt == nil {
		respondError(w, http.StatusNotFound, "Appointment not found")
		return
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    apt,
	})
}

// GetTestimonials returns all active testimonials
func (h *Handler) GetTestimonials(w http.ResponseWriter, r *http.Request) {
	testimonials, err := h.db.GetTestimonials()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to get testimonials")
		return
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    testimonials,
	})
}

// Helper functions

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, models.APIResponse{
		Success: false,
		Error:   message,
	})
}
