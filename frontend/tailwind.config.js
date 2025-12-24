/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        // Eucalyptus Green palette
        primary: {
          50: '#f0f7f4',
          100: '#DAF2E8',  // Lightest
          200: '#c5e8d9',
          300: '#A3D9B0',  // Medium
          400: '#7fc49a',
          500: '#5F8575',  // Main brand color
          600: '#4d6e60',
          700: '#3f5a4e',
          800: '#354a41',
          900: '#2d3e37',
          950: '#17211d'
        },
        // Accent colors for buttons and highlights
        accent: {
          DEFAULT: '#5F8575',
          light: '#A3D9B0',
          lighter: '#DAF2E8',
          dark: '#4d6e60'
        },
        // Warm accent for CTAs (complementary)
        cta: {
          DEFAULT: '#E8A838',  // Warm gold/amber
          hover: '#d49730',
          light: '#FFF3D6'
        }
      },
      fontFamily: {
        sans: ['Inter', 'system-ui', 'sans-serif'],
        serif: ['Georgia', 'serif']
      }
    }
  },
  plugins: []
};
