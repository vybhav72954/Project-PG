# Friends2Health Homoeo Clinic Website

[![Frontend](https://img.shields.io/badge/frontend-SvelteKit-FF3E00)](./frontend)
[![Backend](https://img.shields.io/badge/backend-Go-00ADD8)](./backend)
[![Database](https://img.shields.io/badge/database-SQLite-003B57)](./backend/data)

Complete code for Dr Aditi Singh's Homoeopathy Clinic **Friends2Health**

## Requirements

- Node.js 18+
- Go 1.21+
- Gmail account for SMTP
- Razorpay account for payments

## Quick Start

```bash
# Terminal 1: Backend
cd backend
cp .env.example .env    # configure your credentials
go run cmd/server/main.go

# Terminal 2: Frontend
cd frontend
npm install
npm run dev
```

## Access

| URL | Description |
|-----|-------------|
| http://localhost:5173 | Frontend |
| http://localhost:5173/admin | Admin Panel |
| http://localhost:8080/api | Backend API |

## Documentation

- [Frontend README](./frontend/SV_README)
- [Backend README](./backend/README.md)

---

## 📄 License

Private project for Friends2Health Homoeo Clinic.
