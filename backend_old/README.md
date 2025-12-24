
# Project Setup Guide

## Environment Variables (.env)

Create `backend/.env` with these variables:

```
PORT=8080 
TIMEZONE=Asia/Kolkata 
GOOGLE_APPLICATION_CREDENTIALS=./credentials.json 
CALENDAR_ID=your_calendar_id@group.calendar.google.com 
DOCTOR_EMAIL=your_email@gmail.com (example - vybhavchaturvedi@gmail.com)
EMAIL_SENDER=your_email@gmail.com  (example - drvybhav@clinic.com) - Optional
EMAIL_SENDER_NAME="Your Clinic Name" (example - Dr Vybhav's Clinic)
SMTP_HOST=smtp.gmail.com 
SMTP_PORT=587 
SMTP_PASSWORD=your_app_password (xxxx xxxx xxxx xxxx)`
```
## Google API Setup

1.  Create Google Cloud Project:
    -   Go to console.cloud.google.com
    -   Create new project
    -   Enable Google Calendar API
2.  Create Service Account:
    -   Go to IAM & Admin > Service Accounts
    -   Create service account
    -   Download JSON key as `credentials.json`
    -   Place in `backend/` directory
3.  Calendar Setup:
    -   Create calendar in Google Calendar
    -   Settings > Integrate calendar > Copy Calendar ID
    -   Share calendar with service account email
4.  Update ```.env``` with your Calendar ID
5.  Create App Password for the Google Account. [Refer](https://knowledge.workspace.google.com/kb/how-to-create-app-passwords-000009237) this
6. Paste the 16 letter password in the `.env` file

## Important Notes

-   Service account email format: [project-name@project-id.iam.gserviceaccount.com](mailto:project-name@project-id.iam.gserviceaccount.com)
-   Verify calendar sharing permissions with service account

## Running the Project

```bash
cd Backend 
go mod tidy 
go run cmd/server/main.go
```