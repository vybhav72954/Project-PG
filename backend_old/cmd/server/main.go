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
	// Load application configuration from environment variables and .env file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Google Calendar service with credentials and configuration
	calendarService, err := services.NewCalendarService(&cfg.Calendar)
	if err != nil {
		log.Fatalf("Failed to initialize calendar service: %v", err)
	}

	// Initialize email service for sending appointment confirmations
	// Uses SMTP configuration from environment variables
	emailService := services.NewEmailService(
		cfg.Email.SenderEmail,
		cfg.Email.SenderName,
		cfg.Email.SMTPHost,
		cfg.Email.SMTPPort,
		cfg.Email.Password,
	)

	// Initialize appointment service which coordinates between calendar and email services
	appointmentService := services.NewAppointmentService(calendarService)

	// Create HTTP handler with initialized services
	handler := api.NewHandler(appointmentService, calendarService, emailService)

	// Setup Chi router with middleware for logging, recovery, and request tracking
	router := chi.NewRouter()
	router.Use(middleware.Logger)    // Log HTTP requests
	router.Use(middleware.Recoverer) // Recover from panics
	router.Use(middleware.RequestID) // Add unique ID to each request
	router.Use(middleware.RealIP)    // Get real IP behind proxy

	// Mount all API routes under /api path
	router.Mount("/api", api.SetupRoutes(handler))

	// Start HTTP server on configured port
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
