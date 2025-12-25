# Friends2Health Frontend

[![SvelteKit](https://img.shields.io/badge/SvelteKit-2.0-FF3E00)](https://kit.svelte.dev/)
[![Tailwind](https://img.shields.io/badge/Tailwind-3.4-38B2AC)](https://tailwindcss.com/)

## Commands

```bash
npm install         # install dependencies
npm run dev         # development server at localhost:5173
npm run build       # production build
npm run preview     # preview production build
```

## Configuration

Site-wide settings are in `src/lib/config.ts`:

```typescript
export const siteConfig = {
  clinicName: 'Friends2health Homoeo Clinic',
  doctor: {
    name: 'Dr. Aditi Singh',
    title: 'BHMS, MD (Hom)',
  },
  contact: {
    phone: '+91-9988776655',
    whatsapp: '919988776655',
    email: 'contact@example.com',
  },
  // ...
};
```

## Admin Panel

Located at `/admin`. Requires password set in backend `ADMIN_PASSWORD` env var.

Auto-logout after 15 minutes of inactivity.

```bash
npm run build
# deploy the 'build' directory
```

For static hosting, switch to `@sveltejs/adapter-static` in `svelte.config.js`.
