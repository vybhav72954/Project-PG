package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application settings
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Calendar CalendarConfig
	Email    EmailConfig
	Razorpay RazorpayConfig
	App      AppConfig
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port        string
	Environment string // development, production
	FrontendURL string // For CORS
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Path string // SQLite database file path
}

// CalendarConfig holds Google Calendar API configuration
type CalendarConfig struct {
	CredentialsFile string
	CalendarID      string
	DoctorEmail     string
	SlotDuration    int // minutes
	WorkStartHour   int // 9 for 9 AM
	WorkEndHour     int // 17 for 5 PM
}

// EmailConfig holds SMTP configuration
type EmailConfig struct {
	SenderEmail string
	SenderName  string
	SMTPHost    string
	SMTPPort    string
	Password    string
}

// RazorpayConfig holds Razorpay payment gateway configuration
type RazorpayConfig struct {
	KeyID     string
	KeySecret string
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	DoctorName     string
	ClinicName     string
	ClinicPhone    string
	ClinicWhatsApp string
	ClinicEmail    string
	TimeZone       string
	AdminPassword  string
}

// getEnv retrieves environment variable or returns default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvRequired retrieves required environment variable
func getEnvRequired(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("required environment variable %s is not set", key)
	}
	return value, nil
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists (ignore error if not found)
	_ = godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Port:        getEnv("PORT", "8080"),
			Environment: getEnv("ENVIRONMENT", "development"),
			FrontendURL: getEnv("FRONTEND_URL", "http://localhost:5173"),
		},
		Database: DatabaseConfig{
			Path: getEnv("DATABASE_PATH", "./data/clinic.db"),
		},
		Calendar: CalendarConfig{
			CredentialsFile: getEnv("GOOGLE_CREDENTIALS_FILE", "credentials.json"),
			CalendarID:      os.Getenv("GOOGLE_CALENDAR_ID"),
			DoctorEmail:     os.Getenv("DOCTOR_EMAIL"),
			SlotDuration:    30, // 30 minutes per slot
			WorkStartHour:   9,  // 9 AM
			WorkEndHour:     19, // 7 PM
		},
		Email: EmailConfig{
			SenderEmail: os.Getenv("SMTP_EMAIL"),
			SenderName:  getEnv("SMTP_SENDER_NAME", "Dr. Aditi's Homeopathy Clinic"),
			SMTPHost:    getEnv("SMTP_HOST", "smtp.gmail.com"),
			SMTPPort:    getEnv("SMTP_PORT", "587"),
			Password:    os.Getenv("SMTP_PASSWORD"),
		},
		Razorpay: RazorpayConfig{
			KeyID:     os.Getenv("RAZORPAY_KEY_ID"),
			KeySecret: os.Getenv("RAZORPAY_KEY_SECRET"),
		},
		App: AppConfig{
			DoctorName:     getEnv("DOCTOR_NAME", "Dr. Aditi Singh"),
			ClinicName:     getEnv("CLINIC_NAME", "Friends2health Homoeo Clinic"),
			ClinicPhone:    getEnv("CLINIC_PHONE", "+91-9877505344"),
			ClinicWhatsApp: getEnv("CLINIC_WHATSAPP", "919877505344"),
			ClinicEmail:    getEnv("CLINIC_EMAIL", "contact@friends2health.com"),
			TimeZone:       getEnv("TIMEZONE", "Asia/Kolkata"),
			AdminPassword:  getEnv("ADMIN_PASSWORD", "admin123"),
		},
	}

	return config, nil
}
