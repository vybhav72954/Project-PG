package db

import (
	"database/sql"
	"dr-aditi-backend/internal/models"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

// Database wraps the SQL database connection
type Database struct {
	conn *sql.DB
}

// New creates a new database connection and initializes tables
func New(dbPath string) (*Database, error) {
	// Ensure directory exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("creating database directory: %w", err)
	}

	// _busy_timeout=10000 waits up to 10 seconds if database is locked
	conn, err := sql.Open("sqlite", dbPath+"?_foreign_keys=on&_busy_timeout=10000")
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	// Limit connections to prevent lock contention
	conn.SetMaxOpenConns(1)

	db := &Database{conn: conn}
	if err := db.initialize(); err != nil {
		return nil, fmt.Errorf("initializing database: %w", err)
	}

	return db, nil
}

// initialize creates all required tables
func (db *Database) initialize() error {
	schema := `
	CREATE TABLE IF NOT EXISTS patients (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		phone TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS appointments (
		id TEXT PRIMARY KEY,
		patient_id TEXT,
		patient_name TEXT NOT NULL,
		patient_email TEXT NOT NULL,
		patient_phone TEXT NOT NULL,
		consultation_type TEXT NOT NULL,
		start_time DATETIME NOT NULL,
		end_time DATETIME NOT NULL,
		status TEXT DEFAULT 'pending_payment',
		payment_status TEXT DEFAULT 'pending',
		payment_id TEXT,
		payment_order_id TEXT,
		amount INTEGER NOT NULL,
		meet_link TEXT,
		calendar_event_id TEXT,
		reminder_24h_sent INTEGER DEFAULT 0,
		reminder_1h_sent INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (patient_id) REFERENCES patients(id)
	);

	CREATE TABLE IF NOT EXISTS testimonials (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		review TEXT NOT NULL,
		rating INTEGER NOT NULL CHECK(rating >= 1 AND rating <= 5),
		condition TEXT,
		is_active INTEGER DEFAULT 1,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_appointments_start_time ON appointments(start_time);
	CREATE INDEX IF NOT EXISTS idx_appointments_status ON appointments(status);
	CREATE INDEX IF NOT EXISTS idx_appointments_payment_order_id ON appointments(payment_order_id);

	CREATE TABLE IF NOT EXISTS blocked_dates (
		id TEXT PRIMARY KEY,
		date DATE NOT NULL UNIQUE,
		reason TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS settings (
		key TEXT PRIMARY KEY,
		value TEXT NOT NULL,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_blocked_dates_date ON blocked_dates(date);
	`

	_, err := db.conn.Exec(schema)
	if err != nil {
		return fmt.Errorf("executing schema: %w", err)
	}

	// Migration: Add reminder columns if they don't exist (for existing databases)
	db.conn.Exec(`ALTER TABLE appointments ADD COLUMN reminder_24h_sent INTEGER DEFAULT 0`)
	db.conn.Exec(`ALTER TABLE appointments ADD COLUMN reminder_1h_sent INTEGER DEFAULT 0`)

	// Migration: Add notes column for doctor notes
	db.conn.Exec(`ALTER TABLE appointments ADD COLUMN notes TEXT DEFAULT ''`)

	// Insert default testimonials if none exist
	var count int
	err = db.conn.QueryRow("SELECT COUNT(*) FROM testimonials").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		if err := db.seedTestimonials(); err != nil {
			return fmt.Errorf("seeding testimonials: %w", err)
		}
	}

	return nil
}

// seedTestimonials adds initial testimonials
func (db *Database) seedTestimonials() error {
	testimonials := []models.Testimonial{
		{
			ID:        "test-1",
			Name:      "Priya",
			Review:    "Dr. Aditi's treatment completely cured my chronic migraine that I had been suffering from for 5 years. Her approach is thorough and she really listens to understand the root cause.",
			Rating:    5,
			Condition: "Chronic Migraine",
		},
		{
			ID:        "test-2",
			Name:      "Rahul",
			Review:    "I was skeptical about homeopathy at first, but Dr. Aditi's treatment for my skin condition worked wonders. The results were gradual but permanent.",
			Rating:    5,
			Condition: "Eczema",
		},
		{
			ID:        "test-3",
			Name:      "Sunita",
			Review:    "My daughter's recurring respiratory issues have significantly improved after consulting with Dr. Aditi. She is very patient with children and explains everything clearly.",
			Rating:    5,
			Condition: "Respiratory Issues",
		},
		{
			ID:        "test-4",
			Name:      "Amit",
			Review:    "The online consultation was very convenient. Dr. Aditi took her time to understand my health history and the treatment plan has been very effective for my digestive problems.",
			Rating:    4,
			Condition: "Digestive Disorders",
		},
		{
			ID:        "test-5",
			Name:      "Meera",
			Review:    "I've been consulting Dr. Aditi for my hormonal imbalance. Her holistic approach and the natural remedies have helped me feel so much better without any side effects.",
			Rating:    5,
			Condition: "Hormonal Imbalance",
		},
		{
			ID:        "test-6",
			Name:      "Vikram",
			Review:    "Dr. Aditi helped me manage my anxiety naturally. Her understanding and compassionate approach made me feel comfortable discussing my mental health concerns.",
			Rating:    5,
			Condition: "Anxiety",
		},
	}

	for _, t := range testimonials {
		_, err := db.conn.Exec(`
			INSERT INTO testimonials (id, name, review, rating, condition, is_active, created_at)
			VALUES (?, ?, ?, ?, ?, 1, ?)
		`, t.ID, t.Name, t.Review, t.Rating, t.Condition, time.Now())
		if err != nil {
			return err
		}
	}

	return nil
}

// Close closes the database connection
func (db *Database) Close() error {
	return db.conn.Close()
}

// CreateAppointment creates a new appointment
func (db *Database) CreateAppointment(apt *models.Appointment) error {
	_, err := db.conn.Exec(`
		INSERT INTO appointments (
			id, patient_id, patient_name, patient_email, patient_phone,
			consultation_type, start_time, end_time, status, payment_status,
			payment_order_id, amount, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, apt.ID, apt.PatientID, apt.PatientName, apt.PatientEmail, apt.PatientPhone,
		apt.ConsultationType, apt.StartTime, apt.EndTime, apt.Status, apt.PaymentStatus,
		apt.PaymentOrderID, apt.Amount, apt.CreatedAt, apt.UpdatedAt)

	return err
}

// GetAppointment retrieves an appointment by ID
func (db *Database) GetAppointment(id string) (*models.Appointment, error) {
	apt := &models.Appointment{}
	err := db.conn.QueryRow(`
		SELECT id, patient_id, patient_name, patient_email, patient_phone,
			consultation_type, start_time, end_time, status, payment_status,
			COALESCE(payment_id, ''), COALESCE(payment_order_id, ''), amount,
			COALESCE(meet_link, ''), COALESCE(calendar_event_id, ''),
			COALESCE(notes, ''), created_at, updated_at
		FROM appointments WHERE id = ?
	`, id).Scan(
		&apt.ID, &apt.PatientID, &apt.PatientName, &apt.PatientEmail, &apt.PatientPhone,
		&apt.ConsultationType, &apt.StartTime, &apt.EndTime, &apt.Status, &apt.PaymentStatus,
		&apt.PaymentID, &apt.PaymentOrderID, &apt.Amount,
		&apt.MeetLink, &apt.CalendarEventID, &apt.Notes,
		&apt.CreatedAt, &apt.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return apt, nil
}

// GetAppointmentByOrderID retrieves an appointment by Razorpay order ID
func (db *Database) GetAppointmentByOrderID(orderID string) (*models.Appointment, error) {
	apt := &models.Appointment{}
	err := db.conn.QueryRow(`
		SELECT id, patient_id, patient_name, patient_email, patient_phone,
			consultation_type, start_time, end_time, status, payment_status,
			COALESCE(payment_id, ''), COALESCE(payment_order_id, ''), amount,
			COALESCE(meet_link, ''), COALESCE(calendar_event_id, ''),
			created_at, updated_at
		FROM appointments WHERE payment_order_id = ?
	`, orderID).Scan(
		&apt.ID, &apt.PatientID, &apt.PatientName, &apt.PatientEmail, &apt.PatientPhone,
		&apt.ConsultationType, &apt.StartTime, &apt.EndTime, &apt.Status, &apt.PaymentStatus,
		&apt.PaymentID, &apt.PaymentOrderID, &apt.Amount,
		&apt.MeetLink, &apt.CalendarEventID,
		&apt.CreatedAt, &apt.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return apt, nil
}

// UpdateAppointmentPayment updates payment details after successful payment
func (db *Database) UpdateAppointmentPayment(id, paymentID, meetLink, calendarEventID string) error {
	_, err := db.conn.Exec(`
		UPDATE appointments
		SET payment_id = ?, payment_status = 'paid', status = 'confirmed',
			meet_link = ?, calendar_event_id = ?, updated_at = ?
		WHERE id = ?
	`, paymentID, meetLink, calendarEventID, time.Now(), id)

	return err
}

// GetBookedSlots returns all confirmed appointments in a date range
func (db *Database) GetBookedSlots(start, end time.Time) ([]time.Time, error) {
	rows, err := db.conn.Query(`
		SELECT start_time FROM appointments
		WHERE start_time >= ? AND start_time < ?
		AND status IN ('confirmed', 'pending_payment')
	`, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slots []time.Time
	for rows.Next() {
		var t time.Time
		if err := rows.Scan(&t); err != nil {
			return nil, err
		}
		slots = append(slots, t)
	}

	return slots, rows.Err()
}

// GetTestimonials returns all active testimonials
func (db *Database) GetTestimonials() ([]models.Testimonial, error) {
	rows, err := db.conn.Query(`
		SELECT id, name, review, rating, COALESCE(condition, ''), created_at
		FROM testimonials
		WHERE is_active = 1
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testimonials []models.Testimonial
	for rows.Next() {
		var t models.Testimonial
		if err := rows.Scan(&t.ID, &t.Name, &t.Review, &t.Rating, &t.Condition, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.IsActive = true
		testimonials = append(testimonials, t)
	}

	return testimonials, rows.Err()
}

// CleanupExpiredPendingAppointments removes pending appointments older than 30 minutes
func (db *Database) CleanupExpiredPendingAppointments() (int64, error) {
	result, err := db.conn.Exec(`
		DELETE FROM appointments
		WHERE status = 'pending_payment'
		AND created_at < datetime('now', '-30 minutes')
	`)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// ============================================================
// ADMIN METHODS
// ============================================================

// GetAllAppointments returns appointments with optional filters
func (db *Database) GetAllAppointments(status string, startDate, endDate time.Time) ([]models.Appointment, error) {
	query := `
		SELECT id, patient_id, patient_name, patient_email, patient_phone,
			consultation_type, start_time, end_time, status, payment_status,
			COALESCE(payment_id, ''), COALESCE(payment_order_id, ''), amount,
			COALESCE(meet_link, ''), COALESCE(calendar_event_id, ''),
			COALESCE(notes, ''), created_at, updated_at
		FROM appointments
		WHERE 1=1
	`
	args := []interface{}{}

	if status != "" && status != "all" {
		query += " AND status = ?"
		args = append(args, status)
	}

	if !startDate.IsZero() {
		query += " AND start_time >= ?"
		args = append(args, startDate)
	}

	if !endDate.IsZero() {
		query += " AND start_time < ?"
		args = append(args, endDate)
	}

	query += " ORDER BY start_time DESC"

	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var apt models.Appointment
		if err := rows.Scan(
			&apt.ID, &apt.PatientID, &apt.PatientName, &apt.PatientEmail, &apt.PatientPhone,
			&apt.ConsultationType, &apt.StartTime, &apt.EndTime, &apt.Status, &apt.PaymentStatus,
			&apt.PaymentID, &apt.PaymentOrderID, &apt.Amount,
			&apt.MeetLink, &apt.CalendarEventID, &apt.Notes,
			&apt.CreatedAt, &apt.UpdatedAt,
		); err != nil {
			return nil, err
		}
		appointments = append(appointments, apt)
	}

	return appointments, rows.Err()
}

// UpdateAppointmentStatus updates the status of an appointment
func (db *Database) UpdateAppointmentStatus(id, status string) error {
	_, err := db.conn.Exec(`
		UPDATE appointments SET status = ?, updated_at = ? WHERE id = ?
	`, status, time.Now(), id)
	return err
}

// UpdateAppointmentNotes updates the notes of an appointment
func (db *Database) UpdateAppointmentNotes(id, notes string) error {
	_, err := db.conn.Exec(`
		UPDATE appointments SET notes = ?, updated_at = ? WHERE id = ?
	`, notes, time.Now(), id)
	return err
}

// GetDashboardStats returns statistics for admin dashboard
func (db *Database) GetDashboardStats() (*models.DashboardStats, error) {
	stats := &models.DashboardStats{}

	// Today's appointments
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	err := db.conn.QueryRow(`
		SELECT COUNT(*) FROM appointments
		WHERE start_time >= ? AND start_time < ? AND status = 'confirmed'
	`, today, tomorrow).Scan(&stats.TodayAppointments)
	if err != nil {
		return nil, err
	}

	// This week's appointments
	weekStart := today.AddDate(0, 0, -int(today.Weekday()))
	weekEnd := weekStart.AddDate(0, 0, 7)

	err = db.conn.QueryRow(`
		SELECT COUNT(*) FROM appointments
		WHERE start_time >= ? AND start_time < ? AND status = 'confirmed'
	`, weekStart, weekEnd).Scan(&stats.WeekAppointments)
	if err != nil {
		return nil, err
	}

	// Total revenue (this month)
	monthStart := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, today.Location())
	nextMonth := monthStart.AddDate(0, 1, 0)

	err = db.conn.QueryRow(`
		SELECT COALESCE(SUM(amount), 0) FROM appointments
		WHERE created_at >= ? AND created_at < ? AND payment_status = 'paid'
	`, monthStart, nextMonth).Scan(&stats.MonthRevenue)
	if err != nil {
		return nil, err
	}

	// Total patients (unique emails)
	err = db.conn.QueryRow(`
		SELECT COUNT(DISTINCT patient_email) FROM appointments WHERE status = 'confirmed'
	`).Scan(&stats.TotalPatients)
	if err != nil {
		return nil, err
	}

	// Pending appointments
	err = db.conn.QueryRow(`
		SELECT COUNT(*) FROM appointments WHERE status = 'pending_payment'
	`).Scan(&stats.PendingAppointments)
	if err != nil {
		return nil, err
	}

	// Upcoming appointments (next 7 days)
	err = db.conn.QueryRow(`
		SELECT COUNT(*) FROM appointments
		WHERE start_time >= ? AND start_time < ? AND status = 'confirmed'
	`, today, today.AddDate(0, 0, 7)).Scan(&stats.UpcomingAppointments)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

// GetAllPatients returns unique patients with their appointment count
func (db *Database) GetAllPatients() ([]models.PatientSummary, error) {
	rows, err := db.conn.Query(`
		SELECT 
			patient_email,
			patient_name,
			patient_phone,
			COUNT(*) as appointment_count,
			MAX(start_time) as last_appointment,
			MIN(created_at) as first_visit
		FROM appointments
		WHERE status IN ('confirmed', 'completed', 'cancelled')
		GROUP BY patient_email
		ORDER BY last_appointment DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	// Common datetime formats used by SQLite/Go
	dateFormats := []string{
		"2006-01-02 15:04:05 -0700 MST",           // Go default: 2025-12-30 10:00:00 +0530 IST
		"2006-01-02 15:04:05.999999999 -0700 MST", // With nanoseconds
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05Z",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05+05:30",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02 15:04:05+05:30",
		"2006-01-02 15:04:05-07:00",
	}

	parseTime := func(s string) time.Time {
		// Strip monotonic clock reading if present (e.g., " m=+692.333939201")
		if idx := strings.Index(s, " m="); idx != -1 {
			s = s[:idx]
		}

		for _, format := range dateFormats {
			if t, err := time.Parse(format, s); err == nil {
				return t
			}
		}
		return time.Time{}
	}

	var patients []models.PatientSummary
	for rows.Next() {
		var p models.PatientSummary
		var lastAppt, firstVisit sql.NullString
		if err := rows.Scan(&p.Email, &p.Name, &p.Phone, &p.AppointmentCount, &lastAppt, &firstVisit); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		if lastAppt.Valid && lastAppt.String != "" {
			p.LastAppointment = parseTime(lastAppt.String)
		}

		if firstVisit.Valid && firstVisit.String != "" {
			p.FirstVisit = parseTime(firstVisit.String)
		}

		patients = append(patients, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return patients, nil
}

// GetPatientAppointments returns all appointments for a patient by email
func (db *Database) GetPatientAppointments(email string) ([]models.Appointment, error) {
	rows, err := db.conn.Query(`
		SELECT id, patient_id, patient_name, patient_email, patient_phone,
			consultation_type, start_time, end_time, status, payment_status,
			COALESCE(payment_id, ''), COALESCE(payment_order_id, ''), amount,
			COALESCE(meet_link, ''), COALESCE(calendar_event_id, ''),
			COALESCE(notes, ''), created_at, updated_at
		FROM appointments
		WHERE patient_email = ?
		ORDER BY start_time DESC
	`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var apt models.Appointment
		if err := rows.Scan(
			&apt.ID, &apt.PatientID, &apt.PatientName, &apt.PatientEmail, &apt.PatientPhone,
			&apt.ConsultationType, &apt.StartTime, &apt.EndTime, &apt.Status, &apt.PaymentStatus,
			&apt.PaymentID, &apt.PaymentOrderID, &apt.Amount,
			&apt.MeetLink, &apt.CalendarEventID, &apt.Notes,
			&apt.CreatedAt, &apt.UpdatedAt,
		); err != nil {
			return nil, err
		}
		appointments = append(appointments, apt)
	}

	return appointments, rows.Err()
}

// GetAllTestimonials returns all testimonials (including inactive)
func (db *Database) GetAllTestimonials() ([]models.Testimonial, error) {
	rows, err := db.conn.Query(`
		SELECT id, name, review, rating, COALESCE(condition, ''), is_active, created_at
		FROM testimonials
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testimonials []models.Testimonial
	for rows.Next() {
		var t models.Testimonial
		var isActive int
		if err := rows.Scan(&t.ID, &t.Name, &t.Review, &t.Rating, &t.Condition, &isActive, &t.CreatedAt); err != nil {
			return nil, err
		}
		t.IsActive = isActive == 1
		testimonials = append(testimonials, t)
	}

	return testimonials, rows.Err()
}

// CreateTestimonial creates a new testimonial
func (db *Database) CreateTestimonial(t *models.Testimonial) error {
	_, err := db.conn.Exec(`
		INSERT INTO testimonials (id, name, review, rating, condition, is_active, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, t.ID, t.Name, t.Review, t.Rating, t.Condition, 1, time.Now())
	return err
}

// UpdateTestimonial updates a testimonial
func (db *Database) UpdateTestimonial(t *models.Testimonial) error {
	isActive := 0
	if t.IsActive {
		isActive = 1
	}
	_, err := db.conn.Exec(`
		UPDATE testimonials SET name = ?, review = ?, rating = ?, condition = ?, is_active = ?
		WHERE id = ?
	`, t.Name, t.Review, t.Rating, t.Condition, isActive, t.ID)
	return err
}

// DeleteTestimonial deletes a testimonial
func (db *Database) DeleteTestimonial(id string) error {
	_, err := db.conn.Exec(`DELETE FROM testimonials WHERE id = ?`, id)
	return err
}

// ToggleTestimonialActive toggles the active status of a testimonial
func (db *Database) ToggleTestimonialActive(id string) error {
	_, err := db.conn.Exec(`
		UPDATE testimonials SET is_active = CASE WHEN is_active = 1 THEN 0 ELSE 1 END
		WHERE id = ?
	`, id)
	return err
}

// GetBlockedDates returns all blocked dates
func (db *Database) GetBlockedDates() ([]models.BlockedDate, error) {
	rows, err := db.conn.Query(`
		SELECT id, date, COALESCE(reason, ''), created_at
		FROM blocked_dates
		ORDER BY date ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dates []models.BlockedDate
	for rows.Next() {
		var d models.BlockedDate
		if err := rows.Scan(&d.ID, &d.Date, &d.Reason, &d.CreatedAt); err != nil {
			return nil, err
		}
		dates = append(dates, d)
	}

	return dates, rows.Err()
}

// AddBlockedDate adds a blocked date
func (db *Database) AddBlockedDate(d *models.BlockedDate) error {
	_, err := db.conn.Exec(`
		INSERT INTO blocked_dates (id, date, reason, created_at)
		VALUES (?, ?, ?, ?)
	`, d.ID, d.Date, d.Reason, time.Now())
	return err
}

// RemoveBlockedDate removes a blocked date
func (db *Database) RemoveBlockedDate(id string) error {
	_, err := db.conn.Exec(`DELETE FROM blocked_dates WHERE id = ?`, id)
	return err
}

// IsDateBlocked checks if a date is blocked
func (db *Database) IsDateBlocked(date time.Time) (bool, error) {
	dateStr := date.Format("2006-01-02")
	var count int
	err := db.conn.QueryRow(`SELECT COUNT(*) FROM blocked_dates WHERE date = ?`, dateStr).Scan(&count)
	return count > 0, err
}

// GetSetting retrieves a setting value
func (db *Database) GetSetting(key string) (string, error) {
	var value string
	err := db.conn.QueryRow(`SELECT value FROM settings WHERE key = ?`, key).Scan(&value)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return value, err
}

// SetSetting sets a setting value
func (db *Database) SetSetting(key, value string) error {
	_, err := db.conn.Exec(`
		INSERT INTO settings (key, value, updated_at) VALUES (?, ?, ?)
		ON CONFLICT(key) DO UPDATE SET value = ?, updated_at = ?
	`, key, value, time.Now(), value, time.Now())
	return err
}

// GetAppointmentsNeedingReminder24h returns confirmed appointments that:
// - Start between 23-25 hours from now (to catch appointments in that window)
// - Haven't had 24h reminder sent yet
func (db *Database) GetAppointmentsNeedingReminder24h() ([]models.Appointment, error) {
	now := time.Now()
	windowStart := now.Add(23 * time.Hour)
	windowEnd := now.Add(25 * time.Hour)

	rows, err := db.conn.Query(`
		SELECT id, patient_id, patient_name, patient_email, patient_phone,
			consultation_type, start_time, end_time, status, payment_status,
			COALESCE(payment_id, ''), COALESCE(payment_order_id, ''), amount,
			COALESCE(meet_link, ''), COALESCE(calendar_event_id, ''),
			created_at, updated_at
		FROM appointments
		WHERE status = 'confirmed'
			AND start_time >= ? AND start_time <= ?
			AND reminder_24h_sent = 0
	`, windowStart, windowEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanAppointments(rows)
}

// GetAppointmentsNeedingReminder1h returns confirmed appointments that:
// - Start between 50 minutes to 70 minutes from now
// - Haven't had 1h reminder sent yet
func (db *Database) GetAppointmentsNeedingReminder1h() ([]models.Appointment, error) {
	now := time.Now()
	windowStart := now.Add(50 * time.Minute)
	windowEnd := now.Add(70 * time.Minute)

	rows, err := db.conn.Query(`
		SELECT id, patient_id, patient_name, patient_email, patient_phone,
			consultation_type, start_time, end_time, status, payment_status,
			COALESCE(payment_id, ''), COALESCE(payment_order_id, ''), amount,
			COALESCE(meet_link, ''), COALESCE(calendar_event_id, ''),
			created_at, updated_at
		FROM appointments
		WHERE status = 'confirmed'
			AND start_time >= ? AND start_time <= ?
			AND reminder_1h_sent = 0
	`, windowStart, windowEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanAppointments(rows)
}

// MarkReminder24hSent marks the 24h reminder as sent for an appointment
func (db *Database) MarkReminder24hSent(appointmentID string) error {
	_, err := db.conn.Exec(`
		UPDATE appointments SET reminder_24h_sent = 1, updated_at = ? WHERE id = ?
	`, time.Now(), appointmentID)
	return err
}

// MarkReminder1hSent marks the 1h reminder as sent for an appointment
func (db *Database) MarkReminder1hSent(appointmentID string) error {
	_, err := db.conn.Exec(`
		UPDATE appointments SET reminder_1h_sent = 1, updated_at = ? WHERE id = ?
	`, time.Now(), appointmentID)
	return err
}

// scanAppointments is a helper to scan appointment rows
func scanAppointments(rows *sql.Rows) ([]models.Appointment, error) {
	var appointments []models.Appointment
	for rows.Next() {
		var apt models.Appointment
		if err := rows.Scan(
			&apt.ID, &apt.PatientID, &apt.PatientName, &apt.PatientEmail, &apt.PatientPhone,
			&apt.ConsultationType, &apt.StartTime, &apt.EndTime, &apt.Status, &apt.PaymentStatus,
			&apt.PaymentID, &apt.PaymentOrderID, &apt.Amount,
			&apt.MeetLink, &apt.CalendarEventID,
			&apt.CreatedAt, &apt.UpdatedAt,
		); err != nil {
			return nil, err
		}
		appointments = append(appointments, apt)
	}
	return appointments, rows.Err()
}
