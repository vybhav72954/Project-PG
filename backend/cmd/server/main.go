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
	adminHandler := api.NewAdminHandlers(database, cfg, emailService)

	// Setup routes
	router := api.SetupRoutes(handler, adminHandler, cfg.Server.FrontendURL)

	// Start reminder scheduler
	reminderScheduler := services.NewReminderScheduler(database, emailService)
	reminderScheduler.Start()
	defer reminderScheduler.Stop()

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

	// Print startup info
	fmt.Println()
	fmt.Printf("Server started on :%s\n", cfg.Server.Port)
	fmt.Printf("  Environment: %s\n", cfg.Server.Environment)
	fmt.Printf("  Frontend:    %s\n", cfg.Server.FrontendURL)
	fmt.Printf("  Database:    %s\n", cfg.Database.Path)

	// Email status
	if cfg.Email.SenderEmail != "" && cfg.Email.Password != "" {
		fmt.Printf("  Email:       configured (%s)\n", cfg.Email.SMTPHost)
	} else {
		fmt.Printf("  Email:       not configured\n")
	}

	// Razorpay status
	if paymentService.IsConfigured() {
		fmt.Printf("  Razorpay:    configured\n")
	} else {
		fmt.Printf("  Razorpay:    not configured (test mode)\n")
	}

	// Calendar status
	if cfg.Calendar.CalendarID != "" {
		fmt.Printf("  Calendar:    configured\n")
	} else {
		fmt.Printf("  Calendar:    not configured (using Jitsi fallback)\n")
	}

	fmt.Printf("  Scheduler:   running\n")
	fmt.Println()

	// Start server
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
