// internal/api/routes.go
package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupRoutes(h *Handler) http.Handler {
	r := chi.NewRouter()

	// Health check endpoint to verify API is running
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
	})

	// Group appointment-related endpoints under /appointments
	r.Route("/appointments", func(r chi.Router) {
		// GET /appointments/available-slots - Returns available time slots
		r.Get("/available-slots", h.GetAvailableSlots)

		// POST /appointments/book - Creates new appointment
		r.Post("/book", h.BookAppointment)
	})

	return r
}
