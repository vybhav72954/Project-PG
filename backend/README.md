# Friends2Health APIs

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8)](https://golang.org/)
[![SQLite](https://img.shields.io/badge/SQLite-3-003B57)](https://www.sqlite.org/)

## Commands

```bash
go mod download                 # install dependencies
go run cmd/server/main.go       # start server at localhost:8080
go build -o server cmd/server/main.go   # build binary
```

## Environment Variables

Create `.env` in this directory:

```env
# ============================================
# REQUIRED - Basic Configuration
# ============================================
PORT=8080
ENVIRONMENT=development
FRONTEND_URL=http://localhost:5173
DATABASE_PATH=./data/clinic.db

# Clinic Information
DOCTOR_NAME=Dr. Aditi Singh
CLINIC_NAME=Friends2health Homoeo Clinic
CLINIC_PHONE=+91-9877505344
CLINIC_WHATSAPP=919877505344
CLINIC_EMAIL=contact@friends2health.com
TIMEZONE=Asia/Kolkata

# ============================================
# OPTIONAL - Payment Gateway (Razorpay)
# ============================================
# Without this: Payments are simulated (for testing)
# With this: Real payments work

# Get keys from: https://dashboard.razorpay.com/app/keys
RAZORPAY_KEY_ID=rzp_test_xxxxxxxxxxxx
RAZORPAY_KEY_SECRET=your_secret_key_here

# ============================================
# OPTIONAL - Email Notifications
# ============================================
# Without this: No emails sent (bookings still work)
# With this: Confirmation emails sent to patients

# For Gmail, use App Password (not your regular password)
# Generate at: https://myaccount.google.com/apppasswords
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_EMAIL=your-email@gmail.com
SMTP_PASSWORD=your-app-password
SMTP_SENDER_NAME=Friends2health Homoeo Clinic

# ============================================
# OPTIONAL - Google Calendar Integration
# ============================================
# Without this: Appointments stored in database only
# With this: Appointments also added to Google Calendar

# Setup instructions below
GOOGLE_CREDENTIALS_FILE=./google-credentials.json
GOOGLE_CALENDAR_ID=your-calendar-id@group.calendar.google.com
DOCTOR_EMAIL=doctor@gmail.com
```

## API Endpoints

### Public

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/health | Health check |
| GET | /api/config | Public site config |
| GET | /api/slots | Available appointment slots |
| GET | /api/testimonials | Active testimonials |
| POST | /api/appointments/book | Create booking |
| POST | /api/appointments/verify-payment | Verify Razorpay payment |

### Admin (requires X-Admin-Token header)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/admin/login | Login |
| GET | /api/admin/dashboard | Statistics |
| GET | /api/admin/appointments | List appointments |
| PATCH | /api/admin/appointments/:id/status | Update status |
| GET | /api/admin/patients | List patients |
| GET | /api/admin/patients/history | Patient history |
| GET | /api/admin/testimonials | List testimonials |
| POST | /api/admin/testimonials | Create testimonial |
| PUT | /api/admin/testimonials/:id | Update testimonial |
| DELETE | /api/admin/testimonials/:id | Delete testimonial |
| GET | /api/admin/blocked-dates | List blocked dates |
| POST | /api/admin/blocked-dates | Block a date |
| DELETE | /api/admin/blocked-dates/:id | Unblock a date |

## Database

SQLite file at `data/clinic.db`. Created automatically on first run.

```bash
# Reset (delete and restart server)
del data\clinic.db

# Backup
copy data\clinic.db data\backup.db

# Restore
copy data\backup.db data\clinic.db
```

Use [DB Browser for SQLite](https://sqlitebrowser.org/) to view or edit.

## Security

- Rate limiting: 100 req/min general, 5/min login, 10/min booking
- Admin routes protected by token
- Razorpay signature verification
- Parameterized SQL queries

## Email Reminders

The scheduler runs every 10 minutes and sends:
- 24 hour reminder
- 1 hour reminder

Reminders are tracked in database to prevent duplicates.

## Deployment

Build and run the binary:

```bash
go build -o server cmd/server/main.go
./server
```
