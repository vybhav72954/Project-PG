// Admin API client for backend communication

const API_BASE = '/api/admin';

// Get stored admin token
function getToken(): string | null {
  if (typeof window !== 'undefined') {
    return sessionStorage.getItem('admin_token');
  }
  return null;
}

// Set admin token
export function setToken(token: string): void {
  if (typeof window !== 'undefined') {
    sessionStorage.setItem('admin_token', token);
  }
}

// Clear admin token
export function clearToken(): void {
  if (typeof window !== 'undefined') {
    sessionStorage.removeItem('admin_token');
  }
}

// Check if logged in
export function isLoggedIn(): boolean {
  return getToken() !== null;
}

export interface DashboardStats {
  today_appointments: number;
  week_appointments: number;
  upcoming_appointments: number;
  pending_appointments: number;
  total_patients: number;
  month_revenue: number;
}

export interface Appointment {
  id: string;
  patient_id: string;
  patient_name: string;
  patient_email: string;
  patient_phone: string;
  consultation_type: string;
  start_time: string;
  end_time: string;
  status: string;
  payment_status: string;
  payment_id: string;
  amount: number;
  meet_link: string;
  notes: string;
  created_at: string;
}

export interface PatientSummary {
  email: string;
  name: string;
  phone: string;
  appointment_count: number;
  last_appointment: string;
  first_visit: string;
}

export interface Testimonial {
  id: string;
  name: string;
  review: string;
  rating: number;
  condition: string;
  is_active: boolean;
  created_at: string;
}

export interface BlockedDate {
  id: string;
  date: string;
  reason: string;
  created_at: string;
}

export interface APIResponse<T> {
  success: boolean;
  data?: T;
  error?: string;
  message?: string;
}

class AdminApiClient {
  private async request<T>(endpoint: string, options?: RequestInit): Promise<APIResponse<T>> {
    const token = getToken();

    try {
      const response = await fetch(`${API_BASE}${endpoint}`, {
        ...options,
        headers: {
          'Content-Type': 'application/json',
          'X-Admin-Token': token || '',
          ...options?.headers
        }
      });

      const data = await response.json();

      if (response.status === 401) {
        clearToken();
        if (typeof window !== 'undefined') {
          window.location.href = '/admin';
        }
      }

      return data;
    } catch (error) {
      console.error('Admin API Error:', error);
      return {
        success: false,
        error: 'Network error. Please try again.'
      };
    }
  }

  async login(password: string): Promise<APIResponse<{ token: string }>> {
    const response = await fetch(`${API_BASE}/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ password })
    });

    const data = await response.json();

    if (data.success && data.data?.token) {
      setToken(data.data.token);
    }

    return data;
  }

  logout(): void {
    clearToken();
  }

  async getDashboard(): Promise<APIResponse<DashboardStats>> {
    return this.request<DashboardStats>('/dashboard');
  }

  async getAppointments(status?: string, startDate?: string, endDate?: string): Promise<APIResponse<Appointment[]>> {
    const params = new URLSearchParams();
    if (status) params.append('status', status);
    if (startDate) params.append('start_date', startDate);
    if (endDate) params.append('end_date', endDate);

    const query = params.toString() ? `?${params.toString()}` : '';
    return this.request<Appointment[]>(`/appointments${query}`);
  }

  async updateAppointmentStatus(id: string, status: string): Promise<APIResponse<void>> {
    return this.request<void>(`/appointments/${id}/status`, {
      method: 'PATCH',
      body: JSON.stringify({ status })
    });
  }

  async updateAppointmentNotes(id: string, notes: string): Promise<APIResponse<void>> {
    return this.request<void>(`/appointments/${id}/notes`, {
      method: 'PATCH',
      body: JSON.stringify({ notes })
    });
  }

  async getPatients(): Promise<APIResponse<PatientSummary[]>> {
    return this.request<PatientSummary[]>('/patients');
  }

  async getPatientHistory(email: string): Promise<APIResponse<Appointment[]>> {
    return this.request<Appointment[]>(`/patients/history?email=${encodeURIComponent(email)}`);
  }

  async getTestimonials(): Promise<APIResponse<Testimonial[]>> {
    return this.request<Testimonial[]>('/testimonials');
  }

  async createTestimonial(testimonial: Omit<Testimonial, 'id' | 'created_at'>): Promise<APIResponse<Testimonial>> {
    return this.request<Testimonial>('/testimonials', {
      method: 'POST',
      body: JSON.stringify(testimonial)
    });
  }

  async updateTestimonial(id: string, testimonial: Partial<Testimonial>): Promise<APIResponse<void>> {
    return this.request<void>(`/testimonials/${id}`, {
      method: 'PUT',
      body: JSON.stringify(testimonial)
    });
  }

  async deleteTestimonial(id: string): Promise<APIResponse<void>> {
    return this.request<void>(`/testimonials/${id}`, {
      method: 'DELETE'
    });
  }

  async toggleTestimonial(id: string): Promise<APIResponse<void>> {
    return this.request<void>(`/testimonials/${id}/toggle`, {
      method: 'PATCH'
    });
  }

  async getBlockedDates(): Promise<APIResponse<BlockedDate[]>> {
    return this.request<BlockedDate[]>('/blocked-dates');
  }

  async addBlockedDate(date: string, reason?: string): Promise<APIResponse<BlockedDate>> {
    return this.request<BlockedDate>('/blocked-dates', {
      method: 'POST',
      body: JSON.stringify({ date, reason })
    });
  }

  async removeBlockedDate(id: string): Promise<APIResponse<void>> {
    return this.request<void>(`/blocked-dates/${id}`, {
      method: 'DELETE'
    });
  }
}

export const adminApi = new AdminApiClient();
