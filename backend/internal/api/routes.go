// internal/api/routes.go
package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupRoutes(h *Handler) http.Handler {
	r := chi.NewRouter()

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
	})

	// Appointment endpoints
	r.Route("/appointments", func(r chi.Router) {
		r.Get("/available-slots", h.GetAvailableSlots)
		r.Post("/book", h.BookAppointment)
	})

	return r
}
