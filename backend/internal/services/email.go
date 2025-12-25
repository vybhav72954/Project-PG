package services

import (
	"dr-aditi-backend/internal/config"
	"dr-aditi-backend/internal/models"
	"fmt"
	"net/smtp"
)

// EmailService handles sending emails
type EmailService struct {
	config    *config.EmailConfig
	appConfig *config.AppConfig
}

// NewEmailService creates a new email service
func NewEmailService(emailCfg *config.EmailConfig, appCfg *config.AppConfig) *EmailService {
	return &EmailService{
		config:    emailCfg,
		appConfig: appCfg,
	}
}

// SendAppointmentConfirmation sends confirmation email after successful booking
func (es *EmailService) SendAppointmentConfirmation(apt *models.Appointment) error {
	if es.config.SenderEmail == "" || es.config.Password == "" {
		fmt.Println("Email service not configured, skipping email")
		return nil
	}

	consultationType := "Video Call"
	if apt.ConsultationType == "voice" {
		consultationType = "Voice Call"
	}

	subject := "Appointment Confirmed - " + es.appConfig.ClinicName

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background: linear-gradient(135deg, #5F8575 0%%, #4d6e60 100%%); padding: 30px; text-align: center; border-radius: 10px 10px 0 0;">
        <h1 style="color: white; margin: 0;">Appointment Confirmed!</h1>
    </div>
    
    <div style="background: #f9f9f9; padding: 30px; border-radius: 0 0 10px 10px;">
        <p style="font-size: 16px;">Dear <strong>%s</strong>,</p>
        
        <p>Thank you for booking an appointment with <strong>%s</strong>. Your consultation has been confirmed.</p>
        
        <div style="background: white; padding: 20px; border-radius: 8px; margin: 20px 0; border-left: 4px solid #5F8575;">
            <h3 style="margin-top: 0; color: #5F8575;">Appointment Details</h3>
            <table style="width: 100%%; border-collapse: collapse;">
                <tr>
                    <td style="padding: 8px 0; color: #666;">Date & Time:</td>
                    <td style="padding: 8px 0; font-weight: bold;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px 0; color: #666;">Consultation Type:</td>
                    <td style="padding: 8px 0; font-weight: bold;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px 0; color: #666;">Amount Paid:</td>
                    <td style="padding: 8px 0; font-weight: bold;">₹%d</td>
                </tr>
            </table>
        </div>
        
        <div style="background: #e8f5e9; padding: 20px; border-radius: 8px; margin: 20px 0;">
            <h3 style="margin-top: 0; color: #2e7d32;">Join Your Consultation</h3>
            <p>Click the button below to join your video/voice consultation:</p>
            <a href="%s" style="display: inline-block; background: #5F8575; color: white; padding: 12px 30px; text-decoration: none; border-radius: 5px; font-weight: bold;">Join Meeting</a>
            <p style="font-size: 12px; color: #666; margin-top: 15px;">
                Or copy this link: <a href="%s" style="color: #5F8575;">%s</a>
            </p>
        </div>
        
        <div style="background: #fff3e0; padding: 15px; border-radius: 8px; margin: 20px 0;">
            <p style="margin: 0;"><strong>📌 Important:</strong> Please join 5 minutes before your scheduled time.</p>
        </div>
        
        <h3>Before Your Appointment</h3>
        <ul style="color: #666;">
            <li>Ensure you have a stable internet connection</li>
            <li>Find a quiet, well-lit place for the consultation</li>
            <li>Keep your medical history and current medications list ready</li>
            <li>Note down any questions you want to ask</li>
        </ul>
        
        <hr style="border: none; border-top: 1px solid #eee; margin: 30px 0;">
        
        <p style="font-size: 14px; color: #666;">
            Need to reschedule? Contact us at:<br>
            📞 <a href="tel:%s" style="color: #5F8575;">%s</a><br>
            💬 <a href="https://wa.me/%s" style="color: #25D366;">WhatsApp</a><br>
            ✉️ <a href="mailto:%s" style="color: #5F8575;">%s</a>
        </p>
        
        <p style="margin-top: 30px;">
            Warm regards,<br>
            <strong>%s</strong><br>
            %s
        </p>
    </div>
    
    <div style="text-align: center; padding: 20px; color: #999; font-size: 12px;">
        <p>© 2025 %s. All rights reserved.</p>
    </div>
</body>
</html>
`,
		apt.PatientName,
		es.appConfig.DoctorName,
		apt.StartTime.Format("Monday, January 2, 2006 at 3:04 PM"),
		consultationType,
		apt.Amount/100,
		apt.MeetLink, apt.MeetLink, apt.MeetLink,
		es.appConfig.ClinicPhone, es.appConfig.ClinicPhone,
		es.appConfig.ClinicWhatsApp,
		es.appConfig.ClinicEmail, es.appConfig.ClinicEmail,
		es.appConfig.DoctorName,
		es.appConfig.ClinicName,
		es.appConfig.ClinicName,
	)

	return es.sendEmail(apt.PatientEmail, subject, body)
}

// SendAppointmentReminder sends a reminder email before the appointment
func (es *EmailService) SendAppointmentReminder(apt *models.Appointment) error {
	if es.config.SenderEmail == "" || es.config.Password == "" {
		return nil
	}

	subject := "Reminder: Your Appointment in 1 Hour - " + es.appConfig.ClinicName

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background: #5F8575; padding: 20px; text-align: center; border-radius: 10px 10px 0 0;">
        <h2 style="color: white; margin: 0;">⏰ Appointment Reminder</h2>
    </div>
    
    <div style="background: #f9f9f9; padding: 30px; border-radius: 0 0 10px 10px;">
        <p>Dear <strong>%s</strong>,</p>
        
        <p>This is a reminder that your consultation with <strong>%s</strong> is in <strong>1 hour</strong>.</p>
        
        <div style="background: white; padding: 20px; border-radius: 8px; margin: 20px 0; text-align: center;">
            <p style="font-size: 18px; margin: 0;">📅 %s</p>
        </div>
        
        <div style="text-align: center; margin: 20px 0;">
            <a href="%s" style="display: inline-block; background: #5F8575; color: white; padding: 15px 40px; text-decoration: none; border-radius: 5px; font-weight: bold; font-size: 16px;">Join Meeting Now</a>
        </div>
        
        <p style="color: #666; font-size: 14px; text-align: center;">
            Please join 5 minutes before your scheduled time.
        </p>
        
        <hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
        
        <p style="font-size: 14px; color: #666; text-align: center;">
            Having issues? Contact us at %s
        </p>
    </div>
</body>
</html>
`,
		apt.PatientName,
		es.appConfig.DoctorName,
		apt.StartTime.Format("Monday, January 2 at 3:04 PM"),
		apt.MeetLink,
		es.appConfig.ClinicPhone,
	)

	return es.sendEmail(apt.PatientEmail, subject, body)
}

// SendAppointmentReminder24h sends a reminder email 24 hours before appointment
func (es *EmailService) SendAppointmentReminder24h(apt *models.Appointment) error {
	if es.config.SenderEmail == "" || es.config.Password == "" {
		fmt.Println("Email service not configured, skipping 24h reminder")
		return nil
	}

	consultationType := "Video Call"
	if apt.ConsultationType == "voice" {
		consultationType = "Voice Call"
	}

	subject := "Reminder: Your Appointment Tomorrow - " + es.appConfig.ClinicName

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background: linear-gradient(135deg, #5F8575 0%%, #4d6e60 100%%); padding: 30px; text-align: center; border-radius: 10px 10px 0 0;">
        <h1 style="color: white; margin: 0; font-size: 24px;">📅 Appointment Tomorrow</h1>
    </div>
    
    <div style="background: #f9f9f9; padding: 30px; border-radius: 0 0 10px 10px;">
        <p style="font-size: 16px;">Dear <strong>%s</strong>,</p>
        
        <p>This is a friendly reminder that you have an appointment scheduled with <strong>%s</strong> tomorrow.</p>
        
        <div style="background: white; padding: 20px; border-radius: 8px; margin: 20px 0; border-left: 4px solid #5F8575;">
            <h3 style="margin-top: 0; color: #5F8575;">Appointment Details</h3>
            <table style="width: 100%%; border-collapse: collapse;">
                <tr>
                    <td style="padding: 8px 0; color: #666;">📅 Date & Time:</td>
                    <td style="padding: 8px 0; font-weight: bold;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 8px 0; color: #666;">📞 Consultation Type:</td>
                    <td style="padding: 8px 0; font-weight: bold;">%s</td>
                </tr>
            </table>
        </div>
        
        <div style="background: #e8f5e9; padding: 20px; border-radius: 8px; margin: 20px 0; text-align: center;">
            <p style="margin: 0 0 15px 0; font-weight: bold; color: #2e7d32;">Your Meeting Link</p>
            <a href="%s" style="display: inline-block; background: #5F8575; color: white; padding: 12px 30px; text-decoration: none; border-radius: 5px; font-weight: bold;">Join Meeting</a>
            <p style="font-size: 12px; color: #666; margin-top: 15px;">
                Save this link - you'll need it tomorrow!
            </p>
        </div>
        
        <div style="background: #fff3e0; padding: 15px; border-radius: 8px; margin: 20px 0;">
            <p style="margin: 0; font-weight: bold; color: #e65100;">📌 Before Your Appointment:</p>
            <ul style="margin: 10px 0 0 0; padding-left: 20px; color: #666;">
                <li>Ensure stable internet connection</li>
                <li>Find a quiet, well-lit space</li>
                <li>Keep your medical history ready</li>
                <li>Note down questions you want to ask</li>
            </ul>
        </div>
        
        <hr style="border: none; border-top: 1px solid #eee; margin: 30px 0;">
        
        <p style="font-size: 14px; color: #666;">
            Need to reschedule? Contact us at least 12 hours before:<br>
            📞 <a href="tel:%s" style="color: #5F8575;">%s</a><br>
            💬 <a href="https://wa.me/%s" style="color: #25D366;">WhatsApp</a>
        </p>
        
        <p style="margin-top: 30px;">
            Looking forward to seeing you!<br>
            <strong>%s</strong><br>
            %s
        </p>
    </div>
    
    <div style="text-align: center; padding: 20px; color: #999; font-size: 12px;">
        <p>© 2025 %s. All rights reserved.</p>
    </div>
</body>
</html>
`,
		apt.PatientName,
		es.appConfig.DoctorName,
		apt.StartTime.Format("Monday, January 2, 2006 at 3:04 PM"),
		consultationType,
		apt.MeetLink,
		es.appConfig.ClinicPhone, es.appConfig.ClinicPhone,
		es.appConfig.ClinicWhatsApp,
		es.appConfig.DoctorName,
		es.appConfig.ClinicName,
		es.appConfig.ClinicName,
	)

	return es.sendEmail(apt.PatientEmail, subject, body)
}

// sendEmail sends an email via SMTP
func (es *EmailService) sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", es.config.SenderEmail, es.config.Password, es.config.SMTPHost)

	msg := fmt.Sprintf("From: %s <%s>\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n",
		es.config.SenderName, es.config.SenderEmail,
		to, subject, body)

	addr := fmt.Sprintf("%s:%s", es.config.SMTPHost, es.config.SMTPPort)
	err := smtp.SendMail(addr, auth, es.config.SenderEmail, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("sending email: %w", err)
	}

	fmt.Printf("Email sent to %s: %s\n", to, subject)
	return nil
}
