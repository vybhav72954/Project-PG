# Friends2Health Homoeo Clinic Website

Complete code for Dr Aditi Singh's Homoeopathy Clinic **Friends2Health**
## ✨ Features

### Patient Features
- Online appointment booking
- Razorpay payment integration
- Email confirmations
- Video/Voice consultation options

### Admin Panel
- Dashboard with stats
- Appointment management (view, complete, cancel)
- Patient list with history
- Testimonials management
- Block dates (holidays/vacations)

## Prerequisites

Before you start, make sure you have:

1. **Go 1.21+** - [Download Go](https://go.dev/dl/)
   ```bash
   go version  # Should show go1.21 or higher
   ```

2. **Node.js 18+** - [Download Node.js](https://nodejs.org/)
   ```bash
   node --version  # Should show v18 or higher
   npm --version
   ```

## Quick Start

### Step 1: Setup Backend

```bash
cd backend

# Copy environment template
cp .env.example .env

# Download Go dependencies
go mod tidy

# Run the server
go run ./cmd/server
```

The backend will start at `http://localhost:8080`

### Step 2: Setup Frontend

Open a **new terminal**:

```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

The frontend will start at `http://localhost:5173`

### Step 3: Open in Browser

Visit `http://localhost:5173`

### Step 4: Access Admin Panel

Visit `http://localhost:5173/admin`

---

## ⚙️ Configuration Guide

### Backend Environment Variables (.env)

Create `backend/.env` with the following:

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

---

## 🔧 Setting Up External Services

### 1. Razorpay (Payment Gateway)

1. **Create Account**: Go to [Razorpay Dashboard](https://dashboard.razorpay.com)
2. **Get Test Keys**:
   - Dashboard → Settings → API Keys → Generate Test Key
   - Copy `Key ID` and `Key Secret`
3. **Add to .env**:
   ```env
   RAZORPAY_KEY_ID=rzp_test_xxxxxxxxxxxx
   RAZORPAY_KEY_SECRET=your_secret_key
   ```

**For Production**: Generate Live Keys instead of Test Keys

### 2. Gmail SMTP (Email Notifications)

1. **Enable 2-Factor Authentication** on your Gmail account
2. **Generate App Password**:
   - Go to [Google App Passwords](https://myaccount.google.com/apppasswords)
   - Select "Mail" and your device
   - Copy the 16-character password
3. **Add to .env**:
   ```env
   SMTP_HOST=smtp.gmail.com
   SMTP_PORT=587
   SMTP_EMAIL=your-email@gmail.com
   SMTP_PASSWORD=xxxx-xxxx-xxxx-xxxx
   SMTP_SENDER_NAME=Friends2health Homoeo Clinic
   ```

### 3. Google Calendar (Optional)

This is the most complex setup. Skip if you just want database-only bookings.

1. **Create Google Cloud Project**:
   - Go to [Google Cloud Console](https://console.cloud.google.com)
   - Create new project

2. **Enable Calendar API**:
   - APIs & Services → Library → Search "Google Calendar API" → Enable

3. **Create Service Account**:
   - APIs & Services → Credentials → Create Credentials → Service Account
   - Download JSON key file
   - Save as `backend/google-credentials.json`

4. **Share Calendar with Service Account**:
   - Open Google Calendar
   - Create a new calendar for appointments
   - Settings → Share with specific people
   - Add the service account email (from JSON file)
   - Give "Make changes to events" permission

5. **Get Calendar ID**:
   - Calendar Settings → Integrate calendar → Calendar ID
   - Looks like: `abc123@group.calendar.google.com`

6. **Add to .env**:
   ```env
   GOOGLE_CREDENTIALS_FILE=./google-credentials.json
   GOOGLE_CALENDAR_ID=abc123@group.calendar.google.com
   DOCTOR_EMAIL=doctor@gmail.com
   ```

---

## 🖼️ Adding Images

### Favicon

Add a favicon image:
```
frontend/static/favicon.png
```
Recommended: 32x32 or 64x64 PNG

### Doctor Photo

The about page has a placeholder. To add a real photo:
1. Add image to `frontend/static/doctor.jpg`
2. Update `frontend/src/routes/about/+page.svelte` to use the image

---

## 🏭 Production Deployment

### Build Frontend

```bash
cd frontend
npm run build
```

This creates a `build/` folder with static files.

### Build Backend

```bash
cd backend
go build -o server ./cmd/server
```

This creates a `server` executable.

### Deployment Options

#### Option A: Single Server (VPS)

1. Upload both `backend/server` and `frontend/build/`
2. Use Nginx to:
   - Serve frontend static files
   - Proxy `/api/*` to backend

Example Nginx config:
```nginx
server {
    listen 80;
    server_name yourdomain.com;

    # Frontend
    location / {
        root /var/www/frontend/build;
        try_files $uri $uri/ /index.html;
    }

    # Backend API
    location /api/ {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

#### Option B: Separate Hosting

- **Frontend**: Vercel, Netlify, Cloudflare Pages
- **Backend**: Railway, Render, DigitalOcean App Platform

Update `FRONTEND_URL` in backend `.env` to your frontend domain.

---

## 📁 Database

The app uses SQLite. The database file is created automatically at the path specified in `DATABASE_PATH`.

Default location: `backend/data/clinic.db`

### Tables Created:
- `patients` - Patient information
- `appointments` - Booking records
- `testimonials` - Patient reviews (seeded with sample data)

### Backup
Simply copy the `.db` file to backup all data.

---

## 🧪 Testing the Flow

1. **Homepage**: Check all sections render
2. **About Page**: Verify doctor info
3. **Testimonials**: Should show 6 sample reviews
4. **Book Appointment**:
   - Fill patient info
   - Select consultation type
   - Pick date/time
   - Complete payment (simulated without Razorpay)
   - See confirmation page

---

## 🔍 Troubleshooting

### "go: command not found"
Install Go from https://go.dev/dl/

### "npm: command not found"  
Install Node.js from https://nodejs.org/

### Backend won't start
- Check if port 8080 is free: `lsof -i :8080`
- Verify .env file exists and has correct format
- Check Go dependencies: `go mod tidy`

### Frontend won't start
- Delete `node_modules` and reinstall: `rm -rf node_modules && npm install`
- Check if port 5173 is free

### Payments not working
- Verify Razorpay keys in .env
- Check browser console for errors
- Ensure Razorpay script loads (check network tab)

### Emails not sending
- Verify SMTP settings
- For Gmail, must use App Password (not regular password)
- Check spam folder

---

## 📞 Support

For issues with this codebase, check:
1. Browser developer console (F12)
2. Backend terminal logs
3. Network tab for API errors

---

## 📄 License

Private project for Friends2Health Homoeo Clinic.
