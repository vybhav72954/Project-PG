// internal/services/email.go
package services

import (
	"Project_PG/Backend/internal/models"
	"fmt"
	"net/smtp"
)

// EmailService handles all email communications
type EmailService struct {
	senderEmail string
	senderName  string
	smtpHost    string
	smtpPort    string
	password    string
}

func NewEmailService(senderEmail, senderName, smtpHost, smtpPort, password string) *EmailService {
	return &EmailService{
		senderEmail: senderEmail,
		senderName:  senderName,
		smtpHost:    smtpHost,
		smtpPort:    smtpPort,
		password:    password,
	}
}

func (s *EmailService) SendAppointmentConfirmation(apt *models.Appointment) error {
	subject := "Appointment Confirmation"
	body := fmt.Sprintf(`
        <html>
        <body>
            <p>Dear <strong>%s</strong>,</p>

            <p>Greetings from <strong>Dr. Aditi's Homeopathy Clinic</strong>!</p>

            <p>Your appointment has been confirmed for:</p>
            <p><strong>%s</strong></p>

            <p>Join your consultation using this link:</p>
            <p><a href="%s">%s</a></p>

            <p>Please join <strong>5 minutes before</strong> the scheduled time.</p>

            <p>Best regards,</p>
            <p><strong>%s</strong></p>

            <p>If you need to reschedule or face any difficulties, feel free to contact us at: 
            <strong>+91-9877505344</strong></p>
        </body>
        </html>
    `, apt.PatientName,
		apt.StartTime.Format("Monday, January 2, 2006 at 3:04 PM"),
		apt.MeetLink, apt.MeetLink, s.senderName)

	return s.sendEmail(apt.PatientEmail, subject, body)
}

func (s *EmailService) sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.senderEmail, s.password, s.smtpHost)

	msg := fmt.Sprintf("From: %s <%s>\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", s.senderName, s.senderEmail, to, subject, body)

	return smtp.SendMail(
		s.smtpHost+":"+s.smtpPort,
		auth,
		s.senderEmail,
		[]string{to},
		[]byte(msg),
	)
}
