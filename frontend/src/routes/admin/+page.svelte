<script lang="ts">
  import { onMount } from 'svelte';
  import { adminApi, isLoggedIn, type DashboardStats } from '$lib/adminApi';

  let isAuthenticated = false;
  let password = '';
  let loginError = '';
  let loginLoading = false;

  let stats: DashboardStats | null = null;
  let loading = true;

  onMount(async () => {
    isAuthenticated = isLoggedIn();
    if (isAuthenticated) {
      await loadDashboard();
    }
    loading = false;
  });

  async function handleLogin() {
    if (!password.trim()) {
      loginError = 'Please enter password';
      return;
    }

    loginLoading = true;
    loginError = '';

    const response = await adminApi.login(password);
    
    if (response.success) {
      isAuthenticated = true;
      await loadDashboard();
    } else {
      loginError = response.error || 'Invalid password';
    }

    loginLoading = false;
  }

  async function loadDashboard() {
    loading = true;
    const response = await adminApi.getDashboard();
    if (response.success && response.data) {
      stats = response.data;
    }
    loading = false;
  }

  function formatCurrency(paise: number): string {
    return '₹' + (paise / 100).toLocaleString('en-IN');
  }
</script>

{#if !isAuthenticated}
  <!-- Login Page -->
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-900 to-gray-800 px-4">
    <div class="max-w-md w-full">
      <div class="bg-white rounded-2xl shadow-xl p-8">
        <div class="text-center mb-8">
          <div class="w-16 h-16 bg-primary-500 rounded-full flex items-center justify-center mx-auto mb-4">
            <i class="fas fa-user-shield text-2xl text-white"></i>
          </div>
          <h1 class="text-2xl font-bold text-gray-900">Admin Login</h1>
          <p class="text-gray-600 mt-2">Dr. Aditi's Homeopathy Clinic</p>
        </div>

        <form on:submit|preventDefault={handleLogin} class="space-y-6">
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700 mb-2">
              Password
            </label>
            <input
              type="password"
              id="password"
              bind:value={password}
              class="input"
              placeholder="Enter admin password"
              disabled={loginLoading}
            />
          </div>

          {#if loginError}
            <div class="bg-red-50 text-red-600 px-4 py-3 rounded-lg text-sm">
              <i class="fas fa-exclamation-circle mr-2"></i>
              {loginError}
            </div>
          {/if}

          <button type="submit" class="btn-primary w-full" disabled={loginLoading}>
            {#if loginLoading}
              <i class="fas fa-spinner fa-spin mr-2"></i>
              Logging in...
            {:else}
              <i class="fas fa-sign-in-alt mr-2"></i>
              Login
            {/if}
          </button>
        </form>

        <div class="mt-6 text-center">
          <a href="/" class="text-sm text-gray-500 hover:text-primary-600">
            <i class="fas fa-arrow-left mr-1"></i>
            Back to Website
          </a>
        </div>
      </div>

      <p class="text-center text-gray-400 text-sm mt-4">
        Default password: admin123
      </p>
    </div>
  </div>
{:else}
  <!-- Dashboard -->
  <div>
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-gray-900">Dashboard</h1>
      <p class="text-gray-600">Welcome back! Here's what's happening today.</p>
    </div>

    {#if loading}
      <div class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-10 w-10 border-4 border-primary-500 border-t-transparent"></div>
      </div>
    {:else if stats}
      <!-- Stats Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
        <!-- Today's Appointments -->
        <div class="bg-white rounded-xl p-6 shadow-sm border border-gray-100">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">Today's Appointments</p>
              <p class="text-3xl font-bold text-gray-900 mt-1">{stats.today_appointments}</p>
            </div>
            <div class="w-12 h-12 bg-blue-100 rounded-xl flex items-center justify-center">
              <i class="fas fa-calendar-day text-blue-600 text-xl"></i>
            </div>
          </div>
        </div>

        <!-- This Week -->
        <div class="bg-white rounded-xl p-6 shadow-sm border border-gray-100">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">This Week</p>
              <p class="text-3xl font-bold text-gray-900 mt-1">{stats.week_appointments}</p>
            </div>
            <div class="w-12 h-12 bg-green-100 rounded-xl flex items-center justify-center">
              <i class="fas fa-calendar-week text-green-600 text-xl"></i>
            </div>
          </div>
        </div>

        <!-- Upcoming (7 days) -->
        <div class="bg-white rounded-xl p-6 shadow-sm border border-gray-100">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">Upcoming (7 days)</p>
              <p class="text-3xl font-bold text-gray-900 mt-1">{stats.upcoming_appointments}</p>
            </div>
            <div class="w-12 h-12 bg-purple-100 rounded-xl flex items-center justify-center">
              <i class="fas fa-clock text-purple-600 text-xl"></i>
            </div>
          </div>
        </div>

        <!-- Month Revenue -->
        <div class="bg-white rounded-xl p-6 shadow-sm border border-gray-100">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">This Month's Revenue</p>
              <p class="text-3xl font-bold text-green-600 mt-1">{formatCurrency(stats.month_revenue)}</p>
            </div>
            <div class="w-12 h-12 bg-green-100 rounded-xl flex items-center justify-center">
              <i class="fas fa-rupee-sign text-green-600 text-xl"></i>
            </div>
          </div>
        </div>

        <!-- Total Patients -->
        <div class="bg-white rounded-xl p-6 shadow-sm border border-gray-100">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">Total Patients</p>
              <p class="text-3xl font-bold text-gray-900 mt-1">{stats.total_patients}</p>
            </div>
            <div class="w-12 h-12 bg-indigo-100 rounded-xl flex items-center justify-center">
              <i class="fas fa-users text-indigo-600 text-xl"></i>
            </div>
          </div>
        </div>

        <!-- Pending Payments -->
        <div class="bg-white rounded-xl p-6 shadow-sm border border-gray-100">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500">Pending Payments</p>
              <p class="text-3xl font-bold text-amber-600 mt-1">{stats.pending_appointments}</p>
            </div>
            <div class="w-12 h-12 bg-amber-100 rounded-xl flex items-center justify-center">
              <i class="fas fa-hourglass-half text-amber-600 text-xl"></i>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="bg-white rounded-xl p-6 shadow-sm border border-gray-100">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">Quick Actions</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <a href="/admin/appointments" class="flex flex-col items-center p-4 bg-gray-50 rounded-xl hover:bg-gray-100 transition-colors">
            <i class="fas fa-calendar-alt text-2xl text-primary-600 mb-2"></i>
            <span class="text-sm text-gray-700">View Appointments</span>
          </a>
          <a href="/admin/patients" class="flex flex-col items-center p-4 bg-gray-50 rounded-xl hover:bg-gray-100 transition-colors">
            <i class="fas fa-user-plus text-2xl text-primary-600 mb-2"></i>
            <span class="text-sm text-gray-700">View Patients</span>
          </a>
          <a href="/admin/testimonials" class="flex flex-col items-center p-4 bg-gray-50 rounded-xl hover:bg-gray-100 transition-colors">
            <i class="fas fa-star text-2xl text-primary-600 mb-2"></i>
            <span class="text-sm text-gray-700">Manage Reviews</span>
          </a>
          <a href="/admin/settings" class="flex flex-col items-center p-4 bg-gray-50 rounded-xl hover:bg-gray-100 transition-colors">
            <i class="fas fa-calendar-times text-2xl text-primary-600 mb-2"></i>
            <span class="text-sm text-gray-700">Block Dates</span>
          </a>
        </div>
      </div>
    {/if}
  </div>
{/if}
