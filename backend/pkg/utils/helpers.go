package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateUUID creates a unique identifier
func GenerateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return time.Now().Format("20060102150405")
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// ValidateHMAC verifies the Calendly webhook signature
func ValidateHMAC(payload []byte, secret string, signature string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expectedMAC))
}

// GenerateRandomString creates a secure random string
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}

// FormatDateTime formats a time.Time for display
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// ParseDateTime parses a string into time.Time
func ParseDateTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}
