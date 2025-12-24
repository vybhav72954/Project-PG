package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

// SetupRoutes configures all API routes
func SetupRoutes(h *Handler, adminH *AdminHandlers, frontendURL string) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(5))

	// CORS configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{frontendURL, "http://localhost:5173", "http://localhost:4173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Admin-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// API routes
	r.Route("/api", func(r chi.Router) {
		// Health check
		r.Get("/health", h.HealthCheck)

		// Public config
		r.Get("/config", h.GetConfig)

		// Slots
		r.Get("/slots", h.GetAvailableSlots)

		// Appointments
		r.Route("/appointments", func(r chi.Router) {
			r.Post("/book", h.CreateBooking)
			r.Post("/verify-payment", h.VerifyPayment)
			r.Get("/", h.GetAppointment)
		})

		// Testimonials (public)
		r.Get("/testimonials", h.GetTestimonials)

		// Admin routes
		r.Route("/admin", func(r chi.Router) {
			// Login (no auth required)
			r.Post("/login", adminH.Login)

			// Protected admin routes
			r.Group(func(r chi.Router) {
				r.Use(adminH.AdminAuthMiddleware)

				// Dashboard
				r.Get("/dashboard", adminH.GetDashboard)

				// Appointments management
				r.Get("/appointments", adminH.GetAppointments)
				r.Patch("/appointments/{id}/status", adminH.UpdateAppointmentStatus)

				// Patients
				r.Get("/patients", adminH.GetPatients)
				r.Get("/patients/history", adminH.GetPatientHistory)

				// Testimonials management
				r.Get("/testimonials", adminH.GetTestimonials)
				r.Post("/testimonials", adminH.CreateTestimonial)
				r.Put("/testimonials/{id}", adminH.UpdateTestimonial)
				r.Delete("/testimonials/{id}", adminH.DeleteTestimonial)
				r.Patch("/testimonials/{id}/toggle", adminH.ToggleTestimonial)

				// Blocked dates
				r.Get("/blocked-dates", adminH.GetBlockedDates)
				r.Post("/blocked-dates", adminH.AddBlockedDate)
				r.Delete("/blocked-dates/{id}", adminH.RemoveBlockedDate)
			})
		})
	})

	return r
}
