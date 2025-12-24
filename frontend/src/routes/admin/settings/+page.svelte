<script lang="ts">
  import { onMount } from 'svelte';
  import { adminApi, type BlockedDate } from '$lib/adminApi';

  let blockedDates: BlockedDate[] = [];
  let loading = true;
  let error = '';

  // Form state
  let newDate = '';
  let newReason = '';
  let addLoading = false;

  onMount(async () => {
    await loadBlockedDates();
  });

  async function loadBlockedDates() {
    loading = true;
    error = '';

    const response = await adminApi.getBlockedDates();

    if (response.success && response.data) {
      blockedDates = response.data;
    } else {
      error = response.error || 'Failed to load blocked dates';
    }

    loading = false;
  }

  async function addBlockedDate() {
    if (!newDate) {
      alert('Please select a date');
      return;
    }

    addLoading = true;

    const response = await adminApi.addBlockedDate(newDate, newReason);

    if (response.success) {
      newDate = '';
      newReason = '';
      await loadBlockedDates();
    } else {
      alert(response.error || 'Failed to block date');
    }

    addLoading = false;
  }

  async function removeBlockedDate(id: string) {
    if (!confirm('Are you sure you want to unblock this date?')) return;

    const response = await adminApi.removeBlockedDate(id);

    if (response.success) {
      await loadBlockedDates();
    } else {
      alert(response.error || 'Failed to unblock date');
    }
  }

  function formatDate(dateStr: string): string {
    const date = new Date(dateStr);
    return date.toLocaleDateString('en-IN', {
      weekday: 'long',
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    });
  }

  function isPastDate(dateStr: string): boolean {
    const today = new Date();
    today.setHours(0, 0, 0, 0);
    return new Date(dateStr) < today;
  }

  // Get today's date for min attribute
  $: todayStr = new Date().toISOString().split('T')[0];
</script>

<div>
  <div class="mb-6">
    <h1 class="text-2xl font-bold text-gray-900">Settings</h1>
    <p class="text-gray-600">Manage clinic settings and blocked dates</p>
  </div>

  <!-- Blocked Dates Section -->
  <div class="bg-white rounded-xl shadow-sm border border-gray-100 mb-6">
    <div class="px-6 py-4 border-b">
      <h2 class="text-lg font-semibold text-gray-900">Blocked Dates</h2>
      <p class="text-sm text-gray-500">Block dates when the clinic is closed (holidays, vacations, etc.)</p>
    </div>

    <div class="p-6">
      <!-- Add New Blocked Date -->
      <div class="bg-gray-50 rounded-xl p-4 mb-6">
        <h3 class="font-medium text-gray-900 mb-4">Block a Date</h3>
        <div class="grid md:grid-cols-3 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Date</label>
            <input 
              type="date" 
              bind:value={newDate} 
              min={todayStr}
              class="input" 
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Reason (optional)</label>
            <input 
              type="text" 
              bind:value={newReason} 
              class="input" 
              placeholder="e.g., Diwali Holiday"
            />
          </div>
          <div class="flex items-end">
            <button on:click={addBlockedDate} class="btn-primary w-full" disabled={addLoading}>
              {#if addLoading}
                <i class="fas fa-spinner fa-spin mr-2"></i>
              {:else}
                <i class="fas fa-plus mr-2"></i>
              {/if}
              Block Date
            </button>
          </div>
        </div>
      </div>

      <!-- Blocked Dates List -->
      {#if loading}
        <div class="flex justify-center py-8">
          <div class="animate-spin rounded-full h-8 w-8 border-4 border-primary-500 border-t-transparent"></div>
        </div>
      {:else if error}
        <div class="bg-red-50 text-red-600 px-4 py-3 rounded-lg text-sm">
          {error}
        </div>
      {:else if blockedDates.length === 0}
        <div class="text-center py-8 text-gray-500">
          <i class="fas fa-calendar-check text-3xl mb-2"></i>
          <p>No dates are currently blocked</p>
        </div>
      {:else}
        <div class="space-y-3">
          {#each blockedDates as blocked}
            <div class="flex items-center justify-between p-4 border rounded-xl {isPastDate(blocked.date) ? 'bg-gray-50 opacity-60' : ''}">
              <div>
                <p class="font-medium text-gray-900">{formatDate(blocked.date)}</p>
                {#if blocked.reason}
                  <p class="text-sm text-gray-500">{blocked.reason}</p>
                {/if}
                {#if isPastDate(blocked.date)}
                  <span class="text-xs text-gray-400">Past date</span>
                {/if}
              </div>
              <button 
                on:click={() => removeBlockedDate(blocked.id)}
                class="text-red-600 hover:text-red-800 p-2"
                title="Unblock"
              >
                <i class="fas fa-times"></i>
              </button>
            </div>
          {/each}
        </div>
      {/if}
    </div>
  </div>

  <!-- Pricing Info (Read-only for now) -->
  <div class="bg-white rounded-xl shadow-sm border border-gray-100 mb-6">
    <div class="px-6 py-4 border-b">
      <h2 class="text-lg font-semibold text-gray-900">Pricing</h2>
      <p class="text-sm text-gray-500">Current consultation prices (change in .env file)</p>
    </div>

    <div class="p-6">
      <div class="grid md:grid-cols-2 gap-4">
        <div class="p-4 bg-gray-50 rounded-xl">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
              <i class="fas fa-video text-blue-600"></i>
            </div>
            <div>
              <p class="font-medium text-gray-900">Video Consultation</p>
              <p class="text-2xl font-bold text-primary-600">₹999</p>
            </div>
          </div>
        </div>
        <div class="p-4 bg-gray-50 rounded-xl">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
              <i class="fas fa-phone text-green-600"></i>
            </div>
            <div>
              <p class="font-medium text-gray-900">Voice Consultation</p>
              <p class="text-2xl font-bold text-primary-600">₹799</p>
            </div>
          </div>
        </div>
      </div>
      <p class="text-sm text-gray-500 mt-4">
        <i class="fas fa-info-circle mr-1"></i>
        To change pricing, update the values in the backend .env file and restart the server.
      </p>
    </div>
  </div>

  <!-- Working Hours Info -->
  <div class="bg-white rounded-xl shadow-sm border border-gray-100">
    <div class="px-6 py-4 border-b">
      <h2 class="text-lg font-semibold text-gray-900">Working Hours</h2>
      <p class="text-sm text-gray-500">Clinic appointment availability</p>
    </div>

    <div class="p-6">
      <div class="grid md:grid-cols-2 gap-6">
        <div>
          <p class="text-sm text-gray-500 mb-1">Working Days</p>
          <p class="font-medium text-gray-900">Monday - Saturday</p>
        </div>
        <div>
          <p class="text-sm text-gray-500 mb-1">Closed</p>
          <p class="font-medium text-gray-900">Sunday</p>
        </div>
        <div>
          <p class="text-sm text-gray-500 mb-1">Appointment Hours</p>
          <p class="font-medium text-gray-900">9:00 AM - 7:00 PM</p>
        </div>
        <div>
          <p class="text-sm text-gray-500 mb-1">Slot Duration</p>
          <p class="font-medium text-gray-900">30 minutes</p>
        </div>
      </div>
      <p class="text-sm text-gray-500 mt-4">
        <i class="fas fa-info-circle mr-1"></i>
        Working hours are configured in the backend. Contact your developer to change these settings.
      </p>
    </div>
  </div>
</div>
