// cmd/server/main.go
package main

import (
	"Project_PG/Backend/internal/api"
	"Project_PG/Backend/internal/config"
	"Project_PG/Backend/internal/services"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize services
	calendarService, err := services.NewCalendarService(&cfg.Calendar)
	if err != nil {
		log.Fatalf("Failed to initialize calendar service: %v", err)
	}

	emailService := services.NewEmailService(
		cfg.Email.SenderEmail,
		cfg.Email.SenderName,
		cfg.Email.SMTPHost,
		cfg.Email.SMTPPort,
		cfg.Email.Password,
	)

	appointmentService := services.NewAppointmentService(calendarService)

	// Initialize handler
	handler := api.NewHandler(appointmentService, calendarService, emailService)

	// Create router with global middleware
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	// Mount API routes
	router.Mount("/api", api.SetupRoutes(handler))

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
