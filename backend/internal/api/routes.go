package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// RateLimiter implements a simple in-memory rate limiter
type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
	limit    int           // max requests
	window   time.Duration // time window
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
	// Cleanup old entries every minute
	go func() {
		for {
			time.Sleep(time.Minute)
			rl.cleanup()
		}
	}()
	return rl
}

func (rl *RateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	for ip, times := range rl.requests {
		var valid []time.Time
		for _, t := range times {
			if now.Sub(t) < rl.window {
				valid = append(valid, t)
			}
		}
		if len(valid) == 0 {
			delete(rl.requests, ip)
		} else {
			rl.requests[ip] = valid
		}
	}
}

func (rl *RateLimiter) isAllowed(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	times := rl.requests[ip]

	// Filter to only recent requests
	var recent []time.Time
	for _, t := range times {
		if now.Sub(t) < rl.window {
			recent = append(recent, t)
		}
	}

	if len(recent) >= rl.limit {
		return false
	}

	rl.requests[ip] = append(recent, now)
	return true
}

// RateLimitMiddleware creates a rate limiting middleware
func RateLimitMiddleware(rl *RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
				ip = realIP
			}

			if !rl.isAllowed(ip) {
				http.Error(w, `{"success":false,"error":"Too many requests. Please try again later."}`, http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// SetupRoutes configures all API routes
func SetupRoutes(h *Handler, adminH *AdminHandlers, frontendURL string) http.Handler {
	r := chi.NewRouter()

	// Rate limiters
	generalLimiter := NewRateLimiter(100, time.Minute) // 100 requests/minute for general
	loginLimiter := NewRateLimiter(5, time.Minute)     // 5 attempts/minute for login
	bookingLimiter := NewRateLimiter(10, time.Minute)  // 10 bookings/minute

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(5))
	r.Use(RateLimitMiddleware(generalLimiter)) // General rate limiting

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

		// Appointments - with booking rate limit
		r.Route("/appointments", func(r chi.Router) {
			r.With(RateLimitMiddleware(bookingLimiter)).Post("/book", h.CreateBooking)
			r.Post("/verify-payment", h.VerifyPayment)
			r.Get("/", h.GetAppointment)
		})

		// Testimonials (public)
		r.Get("/testimonials", h.GetTestimonials)

		// Admin routes
		r.Route("/admin", func(r chi.Router) {
			// Login with strict rate limit
			r.With(RateLimitMiddleware(loginLimiter)).Post("/login", adminH.Login)

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
