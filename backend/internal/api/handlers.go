// internal/api/handlers.go
package api

import (
	"Project_PG/Backend/internal/models"
	"Project_PG/Backend/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Handler struct {
	appointmentService *services.AppointmentService
	calendarService    *services.CalendarService
	emailService       *services.EmailService
}

func NewHandler(as *services.AppointmentService, cs *services.CalendarService, es *services.EmailService) *Handler {
	return &Handler{
		appointmentService: as,
		calendarService:    cs,
		emailService:       es,
	}
}

func (h *Handler) BookAppointment(w http.ResponseWriter, r *http.Request) {
	var req models.BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	apt, err := h.appointmentService.CreateAppointment(req)
	if err != nil {
		// Add error logging
		fmt.Printf("Error creating appointment: %v\n", err)
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating appointment: %v", err))
		return
	}

	go h.emailService.SendAppointmentConfirmation(apt)
	respondWithJSON(w, http.StatusCreated, apt)
}

func (h *Handler) GetAvailableSlots(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid start date format")
		return
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid end date format")
		return
	}

	slots, err := h.calendarService.GetAvailableSlots(start, end)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error fetching available slots")
		return
	}

	respondWithJSON(w, http.StatusOK, slots)
}

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
