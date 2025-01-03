package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Server   ServerConfig
	Calendar CalendarConfig
	Email    EmailConfig
}

type ServerConfig struct {
	Port     string
	TimeZone string
}

type CalendarConfig struct {
	CredentialsFile string
	CalendarID      string
	DoctorEmail     string
	SlotDuration    int
}

type EmailConfig struct {
	SenderEmail string
	SenderName  string
	SMTPHost    string
	SMTPPort    string
	Password    string
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("loading .env file: %w", err)
	}

	config := &Config{
		Server: ServerConfig{
			Port:     getEnvOrDefault("PORT", "8080"),
			TimeZone: getEnvOrDefault("TIMEZONE", "Asia/Kolkata"),
		},
		Calendar: CalendarConfig{
			CredentialsFile: getEnvOrDefault("GOOGLE_APPLICATION_CREDENTIALS", "credentials.json"),
			CalendarID:      os.Getenv("CALENDAR_ID"),
			DoctorEmail:     os.Getenv("DOCTOR_EMAIL"),
			SlotDuration:    30,
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
