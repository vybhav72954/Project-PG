package services

import (
	"dr-aditi-backend/internal/db"
	"log"
	"time"
)

// ReminderScheduler handles scheduling and sending appointment reminders
type ReminderScheduler struct {
	database     *db.Database
	emailService *EmailService
	stopChan     chan struct{}
}

// NewReminderScheduler creates a new reminder scheduler
func NewReminderScheduler(database *db.Database, emailService *EmailService) *ReminderScheduler {
	return &ReminderScheduler{
		database:     database,
		emailService: emailService,
		stopChan:     make(chan struct{}),
	}
}

// Start begins the reminder scheduler
// It checks for appointments needing reminders every 10 minutes
func (rs *ReminderScheduler) Start() {
	// Run immediately on start
	rs.checkAndSendReminders()

	// Then run every 10 minutes
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				rs.checkAndSendReminders()
			case <-rs.stopChan:
				ticker.Stop()
				log.Println("Reminder scheduler stopped")
				return
			}
		}
	}()
}

// Stop stops the reminder scheduler
func (rs *ReminderScheduler) Stop() {
	close(rs.stopChan)
}

// checkAndSendReminders checks for appointments needing reminders and sends them
func (rs *ReminderScheduler) checkAndSendReminders() {
	rs.send24hReminders()
	rs.send1hReminders()
}

// send24hReminders sends 24-hour reminders
func (rs *ReminderScheduler) send24hReminders() {
	appointments, err := rs.database.GetAppointmentsNeedingReminder24h()
	if err != nil {
		log.Printf("Error getting appointments for 24h reminder: %v", err)
		return
	}

	for _, apt := range appointments {
		log.Printf("Sending 24h reminder for appointment %s to %s", apt.ID, apt.PatientEmail)

		err := rs.emailService.SendAppointmentReminder24h(&apt)
		if err != nil {
			log.Printf("Error sending 24h reminder to %s: %v", apt.PatientEmail, err)
			continue
		}

		// Mark as sent
		err = rs.database.MarkReminder24hSent(apt.ID)
		if err != nil {
			log.Printf("Error marking 24h reminder as sent for %s: %v", apt.ID, err)
		} else {
			log.Printf("24h reminder sent to %s", apt.PatientEmail)
		}
	}

	if len(appointments) > 0 {
		log.Printf("Processed %d 24-hour reminders", len(appointments))
	}
}

// send1hReminders sends 1-hour reminders
func (rs *ReminderScheduler) send1hReminders() {
	appointments, err := rs.database.GetAppointmentsNeedingReminder1h()
	if err != nil {
		log.Printf("Error getting appointments for 1h reminder: %v", err)
		return
	}

	for _, apt := range appointments {
		log.Printf("Sending 1h reminder for appointment %s to %s", apt.ID, apt.PatientEmail)

		err := rs.emailService.SendAppointmentReminder(&apt)
		if err != nil {
			log.Printf("Error sending 1h reminder to %s: %v", apt.PatientEmail, err)
			continue
		}

		// Mark as sent
		err = rs.database.MarkReminder1hSent(apt.ID)
		if err != nil {
			log.Printf("Error marking 1h reminder as sent for %s: %v", apt.ID, err)
		} else {
			log.Printf("1h reminder sent to %s", apt.PatientEmail)
		}
	}

	if len(appointments) > 0 {
		log.Printf("Processed %d 1-hour reminders", len(appointments))
	}
}
