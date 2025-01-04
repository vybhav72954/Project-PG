// internal/api/handlers.go
package api

import (
	"Project_PG/Backend/internal/models"
	"Project_PG/Backend/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler contains services needed to process HTTP requests
type Handler struct {
	appointmentService *services.AppointmentService
	calendarService    *services.CalendarService
	emailService       *services.EmailService
}

// NewHandler creates a new handler instance with required services
func NewHandler(as *services.AppointmentService, cs *services.CalendarService, es *services.EmailService) *Handler {
	return &Handler{
		appointmentService: as,
		calendarService:    cs,
		emailService:       es,
	}
}

// BookAppointment handles POST requests to create new appointments
// Expects JSON body with patient details and appointment time
// Creates calendar event and sends confirmation email
func (h *Handler) BookAppointment(w http.ResponseWriter, r *http.Request) {
	// Decode request body into BookingRequest struct
	var req models.BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Create appointment using appointment service
	apt, err := h.appointmentService.CreateAppointment(req)
	if err != nil {
		fmt.Printf("Error creating appointment: %v\n", err)
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating appointment: %v", err))
		return
	}

	// Send confirmation email asynchronously
	go h.emailService.SendAppointmentConfirmation(apt)
	respondWithJSON(w, http.StatusCreated, apt)
}

// GetAvailableSlots handles GET requests to fetch available appointment slots
// Requires start_date and end_date query parameters in YYYY-MM-DD format
func (h *Handler) GetAvailableSlots(w http.ResponseWriter, r *http.Request) {
	// ... rest of the code remains same
}

// Helper functions for consistent response handling
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.APIResponse{
		Success: false,
		Error:   message,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
