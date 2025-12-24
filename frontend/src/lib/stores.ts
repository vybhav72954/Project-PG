import { writable } from 'svelte/store';
import type { ClinicConfig } from './api';

// Clinic configuration store
export const clinicConfig = writable<ClinicConfig | null>(null);

// Booking flow store
export interface BookingState {
  step: number;
  name: string;
  email: string;
  phone: string;
  consultationType: 'video' | 'voice' | null;
  selectedDate: string | null;
  selectedTime: string | null;
  appointmentId: string | null;
  orderId: string | null;
  meetLink: string | null;
}

const initialBookingState: BookingState = {
  step: 1,
  name: '',
  email: '',
  phone: '',
  consultationType: null,
  selectedDate: null,
  selectedTime: null,
  appointmentId: null,
  orderId: null,
  meetLink: null
};

function createBookingStore() {
  const { subscribe, set, update } = writable<BookingState>(initialBookingState);

  return {
    subscribe,
    setStep: (step: number) => update((s) => ({ ...s, step })),
    setPatientInfo: (name: string, email: string, phone: string) =>
      update((s) => ({ ...s, name, email, phone })),
    setConsultationType: (type: 'video' | 'voice') =>
      update((s) => ({ ...s, consultationType: type })),
    setDateTime: (date: string, time: string) =>
      update((s) => ({ ...s, selectedDate: date, selectedTime: time })),
    setOrderDetails: (appointmentId: string, orderId: string) =>
      update((s) => ({ ...s, appointmentId, orderId })),
    setMeetLink: (meetLink: string) => update((s) => ({ ...s, meetLink })),
    nextStep: () => update((s) => ({ ...s, step: Math.min(s.step + 1, 4) })),
    prevStep: () => update((s) => ({ ...s, step: Math.max(s.step - 1, 1) })),
    reset: () => set(initialBookingState)
  };
}

export const booking = createBookingStore();

// Loading state
export const isLoading = writable(false);

// Toast notifications
export interface Toast {
  id: string;
  type: 'success' | 'error' | 'info';
  message: string;
}

function createToastStore() {
  const { subscribe, update } = writable<Toast[]>([]);

  return {
    subscribe,
    add: (type: Toast['type'], message: string) => {
      const id = Math.random().toString(36).substring(2);
      update((toasts) => [...toasts, { id, type, message }]);
      setTimeout(() => {
        update((toasts) => toasts.filter((t) => t.id !== id));
      }, 5000);
    },
    remove: (id: string) => {
      update((toasts) => toasts.filter((t) => t.id !== id));
    }
  };
}

export const toasts = createToastStore();
