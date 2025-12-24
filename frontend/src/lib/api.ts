// API client for backend communication

const API_BASE = '/api';

export interface TimeSlot {
  time: string;
  available: boolean;
}

export interface DaySlots {
  date: string;
  slots: TimeSlot[];
  is_weekend: boolean;
}

export interface BookingRequest {
  name: string;
  email: string;
  phone: string;
  consultation_type: 'video' | 'voice';
  date: string;
  time_slot: string;
}

export interface BookingResponse {
  appointment_id: string;
  order_id: string;
  amount: number;
  currency: string;
  razorpay_key_id: string;
  patient_name: string;
  patient_email: string;
  patient_phone: string;
  appointment_time: string;
}

export interface PaymentVerification {
  razorpay_order_id: string;
  razorpay_payment_id: string;
  razorpay_signature: string;
  appointment_id: string;
}

export interface Testimonial {
  id: string;
  name: string;
  review: string;
  rating: number;
  condition: string;
}

export interface ClinicConfig {
  razorpay_key_id: string;
  pricing: {
    video: number;
    voice: number;
  };
  clinic: {
    name: string;
    doctor: string;
    phone: string;
    whatsapp: string;
    email: string;
  };
}

export interface APIResponse<T> {
  success: boolean;
  data?: T;
  error?: string;
  message?: string;
}

class ApiClient {
  private async request<T>(endpoint: string, options?: RequestInit): Promise<APIResponse<T>> {
    try {
      const response = await fetch(`${API_BASE}${endpoint}`, {
        ...options,
        headers: {
          'Content-Type': 'application/json',
          ...options?.headers
        }
      });

      const data = await response.json();
      return data;
    } catch (error) {
      console.error('API Error:', error);
      return {
        success: false,
        error: 'Network error. Please try again.'
      };
    }
  }

  async getConfig(): Promise<APIResponse<ClinicConfig>> {
    return this.request<ClinicConfig>('/config');
  }

  async getAvailableSlots(startDate?: string, endDate?: string): Promise<APIResponse<DaySlots[]>> {
    let url = '/slots';
    const params = new URLSearchParams();
    if (startDate) params.append('start_date', startDate);
    if (endDate) params.append('end_date', endDate);
    if (params.toString()) url += `?${params.toString()}`;
    
    return this.request<DaySlots[]>(url);
  }

  async createBooking(booking: BookingRequest): Promise<APIResponse<BookingResponse>> {
    return this.request<BookingResponse>('/appointments/book', {
      method: 'POST',
      body: JSON.stringify(booking)
    });
  }

  async verifyPayment(verification: PaymentVerification): Promise<APIResponse<{ appointment_id: string; status: string; meet_link: string }>> {
    return this.request('/appointments/verify-payment', {
      method: 'POST',
      body: JSON.stringify(verification)
    });
  }

  async getTestimonials(): Promise<APIResponse<Testimonial[]>> {
    return this.request<Testimonial[]>('/testimonials');
  }
}

export const api = new ApiClient();
