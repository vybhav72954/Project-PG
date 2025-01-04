package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config holds all application configuration settings
type Config struct {
	Server   ServerConfig   // HTTP server settings
	Calendar CalendarConfig // Google Calendar integration settings
	Email    EmailConfig    // Email service settings
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port     string // Server port to listen on
	TimeZone string // Application timezone
}

// CalendarConfig holds Google Calendar API configuration
type CalendarConfig struct {
	CredentialsFile string // Path to Google Calendar credentials JSON
	CalendarID      string // Google Calendar ID for appointments
	DoctorEmail     string // Doctor's email for calendar sharing
	SlotDuration    int    // Duration of each appointment slot in minutes
}

// EmailConfig holds SMTP server configuration for sending emails
type EmailConfig struct {
	SenderEmail string // Email address used to send notifications
	SenderName  string // Display name for the sender
	SMTPHost    string // SMTP server hostname
	SMTPPort    string // SMTP server port
	Password    string // SMTP authentication password
}

// getEnvOrDefault retrieves environment variable or returns default value if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// LoadConfig loads application configuration from environment variables
// Returns error if required environment variables are missing or invalid
func LoadConfig() (*Config, error) {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("loading .env file: %w", err)
	}

	// Construct configuration with environment variables and defaults
	config := &Config{
		Server: ServerConfig{
			Port:     getEnvOrDefault("PORT", "8080"),
			TimeZone: getEnvOrDefault("TIMEZONE", "Asia/Kolkata"),
		},
		Calendar: CalendarConfig{
			CredentialsFile: getEnvOrDefault("GOOGLE_APPLICATION_CREDENTIALS", "credentials.json"),
			CalendarID:      os.Getenv("CALENDAR_ID"),
			DoctorEmail:     os.Getenv("DOCTOR_EMAIL"),
			SlotDuration:    30, // Fixed 30-minute slots
		},
		Email: EmailConfig{
			SenderEmail: os.Getenv("EMAIL_SENDER"),
			SenderName:  os.Getenv("EMAIL_SENDER_NAME"),
			SMTPHost:    os.Getenv("SMTP_HOST"),
			SMTPPort:    os.Getenv("SMTP_PORT"),
			Password:    os.Getenv("SMTP_PASSWORD"),
		},
	}

	return config, nil
}
