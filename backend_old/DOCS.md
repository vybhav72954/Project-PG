# API Documentation
Base URL: ```http://localhost:8080/api```
## Get Available Slots

Retrieves available appointment slots for a date range.

```bash
GET /api/appointments/available-slots
```

### Query Parameters:

```
start_date: YYYY-MM-DD
end_date: YYYY-MM-DD
```

### Success Response (200):

```json
[
  {
    "start_time": "2024-01-04T09:00:00Z",
    "end_time": "2024-01-04T09:30:00Z",
    "available": true
  }
]
```
### Error Responses:

**500 Internal Server Error**

```json
{
  "success": false,
  "error": "Error fetching events: {specific error message}"
}
```

Common error cases:

- Calendar credentials reading failure
- Calendar access verification failure
- Google Calendar API errors

## Book Appointment

Creates a new appointment.
```bash
POST /api/appointments/book
```

### Request Body:
```json
{
  "patient_name": "Test Patient",
  "patient_email": "patient@example.com",
  "start_time": "2024-01-04T10:00:00Z"
}
```

### Success Response (201):
```json
{
"id": "uuid",
"event_id": "google_calendar_id",
"patient_name": "Test Patient",
"patient_email": "patient@example.com",
"start_time": "2024-01-04T10:00:00Z",
"end_time": "2024-01-04T10:30:00Z",
"status": "confirmed",
"meet_link": "https://meet.jit.si/roomId",
"created_at": "2024-01-04T10:00:00Z",
"updated_at": "2024-01-04T10:00:00Z"
}
```

### Error Responses:

**400 Bad Request**

```json
{
  "success": false,
  "error": "Invalid request body"
}
```

**500 Internal Server Error**
```json
{
"success": false,
"error": "Error creating calendar event: {specific error message}"
}
```

Common error cases:

- Service Initialization:
  - Credentials file not found or invalid
  - JWT configuration parsing failure
  - Google Calendar service initialization failure
  - Calendar access verification failure

- Event Creation:

  - Google Calendar API errors
  - Invalid time slot
  - Calendar permission issues


## Data Models

### TimeSlot
```go
type TimeSlot struct {
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    Available bool      `json:"available"`
}
```

### BookingRequest
```go
type BookingRequest struct {
    PatientName  string    `json:"patient_name" validate:"required"`
    PatientEmail string    `json:"patient_email" validate:"required,email"`
    StartTime    time.Time `json:"start_time" validate:"required"`
}
```

### Appointment
```go
type Appointment struct {
    ID           string    `json:"id"`
    EventID      string    `json:"event_id"`
    PatientName  string    `json:"patient_name"`
    PatientEmail string    `json:"patient_email"`
    StartTime    time.Time `json:"start_time"`
    EndTime      time.Time `json:"end_time"`
    Status       string    `json:"status"`
    MeetLink     string    `json:"meet_link"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
```

## Additional Features

### Email Notifications
- Automatic email sent to patient after successful booking
- Contains:
    - Appointment confirmation
    - Date and time
    - Jitsi Meet link
    - Doctor's details

### Error Handling
Common error scenarios:

1. Service Initialization:
```
- Credentials file not found
- Invalid credentials
- Calendar access failure
```

2. Appointment Booking:
```
- Invalid email format
- Invalid date/time format
- Slot already booked
- Calendar API errors
```

3. Slot Retrieval:
```
- Invalid date range
- Calendar API errors
```

### Time Zones
- All API requests/responses use UTC
- Internal calendar storage uses Asia/Kolkata
- Frontend should handle timezone conversions

## Integration Notes
- Use UTC for all date/time fields in requests
- Handle timezone conversions in frontend
- Email notifications are automatic
- Jitsi Meet links are generated automatically