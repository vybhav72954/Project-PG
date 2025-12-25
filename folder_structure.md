```
Project-PG/
├── README.md
│
├── frontend/
│   ├── README.md
│   ├── package.json
│   ├── package-lock.json
│   ├── svelte.config.js
│   ├── vite.config.ts
│   ├── tsconfig.json
│   ├── tailwind.config.js
│   ├── postcss.config.js
│   │
│   ├── src/
│   │   ├── app.html
│   │   ├── app.css
│   │   │
│   │   ├── lib/
│   │   │   ├── index.ts
│   │   │   ├── api.ts                    # Public API calls
│   │   │   ├── adminApi.ts               # Admin API calls
│   │   │   ├── config.ts                 # Site configuration
│   │   │   ├── stores.ts                 # Svelte stores
│   │   │   └── components/
│   │   │       └── Toast.svelte
│   │   │
│   │   └── routes/
│   │       ├── +layout.svelte            # Main layout (header, footer)
│   │       ├── +page.svelte              # Homepage
│   │       │
│   │       ├── about/
│   │       │   └── +page.svelte          # About the doctor
│   │       │
│   │       ├── appointment/
│   │       │   ├── +page.svelte          # Booking flow
│   │       │   ├── confirmed/
│   │       │   │   └── +page.svelte      # Success page
│   │       │   └── failed/
│   │       │       └── +page.svelte      # Payment failed
│   │       │
│   │       ├── testimonials/
│   │       │   └── +page.svelte          # Patient reviews
│   │       │
│   │       ├── blog/
│   │       │   └── +page.svelte          # Research/articles
│   │       │
│   │       ├── terms/
│   │       │   └── +page.svelte          # Privacy, terms, refund
│   │       │
│   │       └── admin/
│   │           ├── +layout.svelte        # Admin layout (sidebar)
│   │           ├── +page.svelte          # Dashboard
│   │           │
│   │           ├── appointments/
│   │           │   └── +page.svelte
│   │           │
│   │           ├── patients/
│   │           │   └── +page.svelte
│   │           │
│   │           ├── testimonials/
│   │           │   └── +page.svelte
│   │           │
│   │           └── settings/
│   │               └── +page.svelte
│   │
│   ├── static/                           # Static assets (images, favicon)
│   │
│   └── build/                            # Production build output (gitignored)
│
└── backend/
    ├── README.md
    ├── .env                              # Environment variables (gitignored)
    ├── .env.example                      # Example env file
    ├── go.mod
    ├── go.sum
    │
    ├── cmd/
    │   └── server/
    │       └── main.go                   # Entry point
    │
    ├── internal/
    │   ├── api/
    │   │   ├── routes.go                 # Routing, middleware, rate limiting
    │   │   ├── handlers.go               # Public API handlers
    │   │   └── admin_handlers.go         # Admin API handlers
    │   │
    │   ├── config/
    │   │   └── config.go                 # Environment config loader
    │   │
    │   ├── db/
    │   │   └── database.go               # SQLite operations
    │   │
    │   ├── models/
    │   │   └── models.go                 # Data structures
    │   │
    │   └── services/
    │       ├── calendar.go               # Google Calendar integration
    │       ├── email.go                  # SMTP email service
    │       ├── payment.go                # Razorpay integration
    │       └── scheduler.go              # Email reminder scheduler
    │
    └── data/
        └── clinic.db                     # SQLite database (gitignored)
```