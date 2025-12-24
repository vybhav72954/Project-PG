package api

import (
	"dr-aditi-backend/internal/config"
	"dr-aditi-backend/internal/db"
	"dr-aditi-backend/internal/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// AdminHandlers handles all admin-related API requests
type AdminHandlers struct {
	db     *db.Database
	config *config.Config
}

// NewAdminHandlers creates a new AdminHandlers instance
func NewAdminHandlers(database *db.Database, cfg *config.Config) *AdminHandlers {
	return &AdminHandlers{
		db:     database,
		config: cfg,
	}
}

// AdminAuthMiddleware checks for admin authentication
func (h *AdminHandlers) AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for admin token in header
		token := r.Header.Get("X-Admin-Token")
		
		// Simple token-based auth - token should match admin password hash or be the password itself
		// In production, use proper JWT tokens
		expectedToken := h.config.App.AdminPassword
		
		if expectedToken == "" {
			// If no admin password set, use default for development
			expectedToken = "admin123"
		}
		
		if token != expectedToken {
			respondJSON(w, http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Error:   "Unauthorized - invalid or missing admin token",
			})
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// Login handles admin login
func (h *AdminHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var req models.AdminLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}
	
	expectedPassword := h.config.App.AdminPassword
	if expectedPassword == "" {
		expectedPassword = "admin123" // Default for development
	}
	
	if req.Password != expectedPassword {
		respondJSON(w, http.StatusUnauthorized, models.APIResponse{
			Success: false,
			Error:   "Invalid password",
		})
		return
	}
	
	// Return the password as token (simple auth - use JWT in production)
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data: map[string]string{
			"token": expectedPassword,
		},
		Message: "Login successful",
	})
}

// GetDashboard returns dashboard statistics
func (h *AdminHandlers) GetDashboard(w http.ResponseWriter, r *http.Request) {
	stats, err := h.db.GetDashboardStats()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to fetch dashboard stats",
		})
		return
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    stats,
	})
}

// GetAppointments returns all appointments with optional filters
func (h *AdminHandlers) GetAppointments(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")
	
	var startDate, endDate time.Time
	var err error
	
	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, models.APIResponse{
				Success: false,
				Error:   "Invalid start_date format",
			})
			return
		}
	}
	
	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, models.APIResponse{
				Success: false,
				Error:   "Invalid end_date format",
			})
			return
		}
		// Include the full end date
		endDate = endDate.Add(24 * time.Hour)
	}
	
	appointments, err := h.db.GetAllAppointments(status, startDate, endDate)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to fetch appointments",
		})
		return
	}
	
	if appointments == nil {
		appointments = []models.Appointment{}
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    appointments,
	})
}

// UpdateAppointmentStatus updates the status of an appointment
func (h *AdminHandlers) UpdateAppointmentStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	var req models.UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}
	
	// Validate status
	validStatuses := map[string]bool{
		"confirmed": true,
		"cancelled": true,
		"completed": true,
	}
	
	if !validStatuses[req.Status] {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid status. Must be: confirmed, cancelled, or completed",
		})
		return
	}
	
	if err := h.db.UpdateAppointmentStatus(id, req.Status); err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to update appointment status",
		})
		return
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Appointment status updated",
	})
}

// GetPatients returns all unique patients
func (h *AdminHandlers) GetPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := h.db.GetAllPatients()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to fetch patients",
		})
		return
	}
	
	if patients == nil {
		patients = []models.PatientSummary{}
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    patients,
	})
}

// GetPatientHistory returns appointment history for a patient
func (h *AdminHandlers) GetPatientHistory(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Email is required",
		})
		return
	}
	
	appointments, err := h.db.GetPatientAppointments(email)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to fetch patient history",
		})
		return
	}
	
	if appointments == nil {
		appointments = []models.Appointment{}
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    appointments,
	})
}

// GetTestimonials returns all testimonials (including inactive)
func (h *AdminHandlers) GetTestimonials(w http.ResponseWriter, r *http.Request) {
	testimonials, err := h.db.GetAllTestimonials()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to fetch testimonials",
		})
		return
	}
	
	if testimonials == nil {
		testimonials = []models.Testimonial{}
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    testimonials,
	})
}

// CreateTestimonial creates a new testimonial
func (h *AdminHandlers) CreateTestimonial(w http.ResponseWriter, r *http.Request) {
	var req models.TestimonialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}
	
	if req.Name == "" || req.Review == "" || req.Rating < 1 || req.Rating > 5 {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Name, review, and rating (1-5) are required",
		})
		return
	}
	
	testimonial := &models.Testimonial{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Review:    req.Review,
		Rating:    req.Rating,
		Condition: req.Condition,
		IsActive:  true,
	}
	
	if err := h.db.CreateTestimonial(testimonial); err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to create testimonial",
		})
		return
	}
	
	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Data:    testimonial,
		Message: "Testimonial created",
	})
}

// UpdateTestimonial updates a testimonial
func (h *AdminHandlers) UpdateTestimonial(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	var req models.TestimonialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}
	
	testimonial := &models.Testimonial{
		ID:        id,
		Name:      req.Name,
		Review:    req.Review,
		Rating:    req.Rating,
		Condition: req.Condition,
		IsActive:  req.IsActive,
	}
	
	if err := h.db.UpdateTestimonial(testimonial); err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to update testimonial",
		})
		return
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Testimonial updated",
	})
}

// DeleteTestimonial deletes a testimonial
func (h *AdminHandlers) DeleteTestimonial(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	if err := h.db.DeleteTestimonial(id); err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to delete testimonial",
		})
		return
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Testimonial deleted",
	})
}

// ToggleTestimonial toggles the active status of a testimonial
func (h *AdminHandlers) ToggleTestimonial(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	if err := h.db.ToggleTestimonialActive(id); err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to toggle testimonial",
		})
		return
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Testimonial toggled",
	})
}

// GetBlockedDates returns all blocked dates
func (h *AdminHandlers) GetBlockedDates(w http.ResponseWriter, r *http.Request) {
	dates, err := h.db.GetBlockedDates()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to fetch blocked dates",
		})
		return
	}
	
	if dates == nil {
		dates = []models.BlockedDate{}
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    dates,
	})
}

// AddBlockedDate adds a new blocked date
func (h *AdminHandlers) AddBlockedDate(w http.ResponseWriter, r *http.Request) {
	var req models.BlockDateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}
	
	if req.Date == "" {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Date is required",
		})
		return
	}
	
	// Validate date format
	_, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid date format. Use YYYY-MM-DD",
		})
		return
	}
	
	blockedDate := &models.BlockedDate{
		ID:     uuid.New().String(),
		Date:   req.Date,
		Reason: req.Reason,
	}
	
	if err := h.db.AddBlockedDate(blockedDate); err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to add blocked date (may already be blocked)",
		})
		return
	}
	
	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Data:    blockedDate,
		Message: "Date blocked",
	})
}

// RemoveBlockedDate removes a blocked date
func (h *AdminHandlers) RemoveBlockedDate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	if err := h.db.RemoveBlockedDate(id); err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to remove blocked date",
		})
		return
	}
	
	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Blocked date removed",
	})
}
