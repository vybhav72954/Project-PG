package main

import (
	"dr-aditi-backend/internal/api"
	"dr-aditi-backend/internal/config"
	"dr-aditi-backend/internal/db"
	"dr-aditi-backend/internal/services"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	database, err := db.New(cfg.Database.Path)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize services
	calendarService, err := services.NewCalendarService(&cfg.Calendar, database, cfg.App.TimeZone)
	if err != nil {
		log.Fatalf("Failed to initialize calendar service: %v", err)
	}

	emailService := services.NewEmailService(&cfg.Email, &cfg.App)

	paymentService, err := services.NewPaymentService(&cfg.Razorpay)
	if err != nil {
		log.Fatalf("Failed to initialize payment service: %v", err)
	}

	// Create handler with all dependencies
	handler := api.NewHandler(database, calendarService, emailService, paymentService, cfg)

	// Create admin handlers
	adminHandler := api.NewAdminHandlers(database, cfg)

	// Setup routes
	router := api.SetupRoutes(handler, adminHandler, cfg.Server.FrontendURL)

	// Start cleanup goroutine for expired pending appointments
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			deleted, err := database.CleanupExpiredPendingAppointments()
			if err != nil {
				log.Printf("Error cleaning up expired appointments: %v", err)
			} else if deleted > 0 {
				log.Printf("Cleaned up %d expired pending appointments", deleted)
			}
		}
	}()

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("🏥 Dr. Aditi's Clinic Backend starting on %s", serverAddr)
	log.Printf("   Environment: %s", cfg.Server.Environment)
	log.Printf("   Frontend URL: %s", cfg.Server.FrontendURL)

	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
