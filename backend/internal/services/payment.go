package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"dr-aditi-backend/internal/config"
	"dr-aditi-backend/internal/models"
	"encoding/hex"
	"fmt"

	razorpay "github.com/razorpay/razorpay-go"
)

// PaymentService handles Razorpay payment operations
type PaymentService struct {
	client *razorpay.Client
	config *config.RazorpayConfig
}

// NewPaymentService creates a new payment service
func NewPaymentService(cfg *config.RazorpayConfig) (*PaymentService, error) {
	if cfg.KeyID == "" || cfg.KeySecret == "" {
		fmt.Println("Warning: Razorpay not configured - payments will be simulated")
		return &PaymentService{config: cfg}, nil
	}

	client := razorpay.NewClient(cfg.KeyID, cfg.KeySecret)

	return &PaymentService{
		client: client,
		config: cfg,
	}, nil
}

// CreateOrder creates a new Razorpay order
func (ps *PaymentService) CreateOrder(apt *models.Appointment) (*models.PaymentOrder, error) {
	if ps.client == nil {
		// Simulate order for development
		return &models.PaymentOrder{
			OrderID:       "order_simulated_" + apt.ID[:8],
			Amount:        apt.Amount,
			Currency:      "INR",
			AppointmentID: apt.ID,
		}, nil
	}

	orderData := map[string]interface{}{
		"amount":   apt.Amount, // in paise
		"currency": "INR",
		"receipt":  apt.ID,
		"notes": map[string]string{
			"appointment_id":    apt.ID,
			"patient_name":      apt.PatientName,
			"patient_email":     apt.PatientEmail,
			"consultation_type": apt.ConsultationType,
		},
	}

	order, err := ps.client.Order.Create(orderData, nil)
	if err != nil {
		return nil, fmt.Errorf("creating Razorpay order: %w", err)
	}

	return &models.PaymentOrder{
		OrderID:       order["id"].(string),
		Amount:        apt.Amount,
		Currency:      "INR",
		AppointmentID: apt.ID,
	}, nil
}

// VerifyPayment verifies the Razorpay payment signature
func (ps *PaymentService) VerifyPayment(verification *models.PaymentVerification) (bool, error) {
	if ps.client == nil {
		// In development mode, accept any payment
		return true, nil
	}

	// Create signature string
	signatureData := verification.RazorpayOrderID + "|" + verification.RazorpayPaymentID

	// Generate HMAC SHA256
	h := hmac.New(sha256.New, []byte(ps.config.KeySecret))
	h.Write([]byte(signatureData))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// Compare signatures
	if expectedSignature != verification.RazorpaySignature {
		return false, fmt.Errorf("signature verification failed")
	}

	return true, nil
}

// GetPayment fetches payment details from Razorpay
func (ps *PaymentService) GetPayment(paymentID string) (map[string]interface{}, error) {
	if ps.client == nil {
		return map[string]interface{}{
			"id":     paymentID,
			"status": "captured",
		}, nil
	}

	payment, err := ps.client.Payment.Fetch(paymentID, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("fetching payment: %w", err)
	}

	return payment, nil
}

// IsConfigured returns whether Razorpay is properly configured
func (ps *PaymentService) IsConfigured() bool {
	return ps.client != nil
}

// GetKeyID returns the Razorpay key ID for frontend
func (ps *PaymentService) GetKeyID() string {
	return ps.config.KeyID
}
