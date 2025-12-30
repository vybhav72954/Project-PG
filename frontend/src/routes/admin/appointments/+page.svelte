<script lang="ts">
  import { onMount } from 'svelte';
  import { adminApi, type Appointment } from '$lib/adminApi';

  let appointments: Appointment[] = [];
  let loading = true;
  let error = '';
  let statusFilter = 'all';
  let dateFilter = '';

  let showNotesModal = false;
  let selectedAppointment: Appointment | null = null;
  let notesInput = '';
  let savingNotes = false;

  onMount(async () => {
    await loadAppointments();
  });

  async function loadAppointments() {
    loading = true;
    error = '';
    const response = await adminApi.getAppointments(
      statusFilter !== 'all' ? statusFilter : undefined,
      dateFilter || undefined
    );
    if (response.success && response.data) {
      appointments = response.data;
    } else {
      error = response.error || 'Failed to load appointments';
    }
    loading = false;
  }

  async function updateStatus(id: string, newStatus: string) {
    const response = await adminApi.updateAppointmentStatus(id, newStatus);
    if (response.success) {
      await loadAppointments();
    } else {
      alert(response.error || 'Failed to update status');
    }
  }

  function openNotesModal(apt: Appointment) {
    selectedAppointment = apt;
    notesInput = apt.notes || '';
    showNotesModal = true;
  }

  async function saveNotes() {
    if (!selectedAppointment) return;
    savingNotes = true;
    const response = await adminApi.updateAppointmentNotes(selectedAppointment.id, notesInput);
    if (response.success) {
      showNotesModal = false;
      await loadAppointments();
    } else {
      alert(response.error || 'Failed to save notes');
    }
    savingNotes = false;
  }

  function formatDate(dateStr: string): string {
    return new Date(dateStr).toLocaleDateString('en-IN', { weekday: 'short', year: 'numeric', month: 'short', day: 'numeric' });
  }

  function formatTime(dateStr: string): string {
    return new Date(dateStr).toLocaleTimeString('en-IN', { hour: '2-digit', minute: '2-digit', hour12: true });
  }

  function getStatusColor(status: string): string {
    switch (status) {
      case 'confirmed': return 'bg-green-100 text-green-800';
      case 'completed': return 'bg-blue-100 text-blue-800';
      case 'cancelled': return 'bg-red-100 text-red-800';
      case 'pending_payment': return 'bg-yellow-100 text-yellow-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  }

  function formatStatus(status: string): string {
    return status.replace('_', ' ').replace(/\b\w/g, l => l.toUpperCase());
  }

  $: if (statusFilter || dateFilter !== undefined) loadAppointments();
</script>

<div>
  <div class="flex flex-col md:flex-row md:items-center md:justify-between mb-6">
    <div>
      <h1 class="text-2xl font-bold text-gray-900">Appointments</h1>
      <p class="text-gray-600">Manage all patient appointments</p>
    </div>
  </div>

  <div class="bg-white rounded-xl p-4 shadow-sm border border-gray-100 mb-6">
    <div class="flex flex-col md:flex-row gap-4">
      <div class="flex-1">
        <label class="block text-sm font-medium text-gray-700 mb-1">Status</label>
        <select bind:value={statusFilter} class="input">
          <option value="all">All Statuses</option>
          <option value="confirmed">Confirmed</option>
          <option value="completed">Completed</option>
          <option value="cancelled">Cancelled</option>
          <option value="pending_payment">Pending Payment</option>
        </select>
      </div>
      <div class="flex-1">
        <label class="block text-sm font-medium text-gray-700 mb-1">Date</label>
        <input type="date" bind:value={dateFilter} class="input" />
      </div>
      <div class="flex items-end">
        <button on:click={() => { statusFilter = 'all'; dateFilter = ''; }} class="btn-secondary">
          <i class="fas fa-times mr-2"></i>Clear Filters
        </button>
      </div>
    </div>
  </div>

  {#if loading}
    <div class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-10 w-10 border-4 border-primary-500 border-t-transparent"></div>
    </div>
  {:else if error}
    <div class="bg-red-50 text-red-600 px-6 py-4 rounded-xl">{error}</div>
  {:else if appointments.length === 0}
    <div class="bg-white rounded-xl p-12 text-center shadow-sm border border-gray-100">
      <i class="fas fa-calendar-times text-4xl text-gray-300 mb-4"></i>
      <p class="text-gray-500">No appointments found</p>
    </div>
  {:else}
    <div class="hidden md:block bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Patient</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Date & Time</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Type</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Amount</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          {#each appointments as apt}
            <tr class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <p class="font-medium text-gray-900">{apt.patient_name}</p>
                <p class="text-sm text-gray-500">{apt.patient_email}</p>
                <p class="text-sm text-gray-500">{apt.patient_phone}</p>
              </td>
              <td class="px-6 py-4">
                <p class="text-gray-900">{formatDate(apt.start_time)}</p>
                <p class="text-sm text-gray-500">{formatTime(apt.start_time)}</p>
              </td>
              <td class="px-6 py-4 capitalize">{apt.consultation_type}</td>
              <td class="px-6 py-4 font-medium">₹{apt.amount / 100}</td>
              <td class="px-6 py-4">
                <span class="px-2 py-1 text-xs font-medium rounded-full {getStatusColor(apt.status)}">{formatStatus(apt.status)}</span>
              </td>
              <td class="px-6 py-4">
                <div class="flex space-x-2">
                  {#if apt.status === 'confirmed'}
                    <button on:click={() => updateStatus(apt.id, 'completed')} class="text-green-600 hover:text-green-800" title="Complete"><i class="fas fa-check-circle"></i></button>
                    <button on:click={() => updateStatus(apt.id, 'cancelled')} class="text-red-600 hover:text-red-800" title="Cancel"><i class="fas fa-times-circle"></i></button>
                    {#if apt.meet_link && apt.consultation_type === 'video'}
                      <a href={apt.meet_link} target="_blank" rel="noopener" class="text-blue-600 hover:text-blue-800" title="Join"><i class="fas fa-video"></i></a>
                    {/if}
                  {/if}
                  <button on:click={() => openNotesModal(apt)} class="text-purple-600 hover:text-purple-800" title="Notes"><i class="fas fa-sticky-note"></i></button>
                </div>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>

    <div class="md:hidden space-y-4">
      {#each appointments as apt}
        <div class="bg-white rounded-xl p-4 shadow-sm border border-gray-100">
          <div class="flex justify-between items-start mb-3">
            <div>
              <p class="font-medium text-gray-900">{apt.patient_name}</p>
              <p class="text-sm text-gray-500">{apt.patient_phone}</p>
            </div>
            <span class="px-2 py-1 text-xs font-medium rounded-full {getStatusColor(apt.status)}">{formatStatus(apt.status)}</span>
          </div>
          <div class="grid grid-cols-2 gap-2 text-sm">
            <div><p class="text-gray-500">Date</p><p class="font-medium">{formatDate(apt.start_time)}</p></div>
            <div><p class="text-gray-500">Time</p><p class="font-medium">{formatTime(apt.start_time)}</p></div>
            <div><p class="text-gray-500">Type</p><p class="font-medium capitalize">{apt.consultation_type}</p></div>
            <div><p class="text-gray-500">Amount</p><p class="font-medium">₹{apt.amount / 100}</p></div>
          </div>
          <div class="flex space-x-2 mt-4 pt-4 border-t">
            {#if apt.status === 'confirmed'}
              <button on:click={() => updateStatus(apt.id, 'completed')} class="flex-1 bg-green-100 text-green-700 py-2 rounded-lg hover:bg-green-200"><i class="fas fa-check mr-1"></i> Complete</button>
              <button on:click={() => updateStatus(apt.id, 'cancelled')} class="flex-1 bg-red-100 text-red-700 py-2 rounded-lg hover:bg-red-200"><i class="fas fa-times mr-1"></i> Cancel</button>
            {/if}
            <button on:click={() => openNotesModal(apt)} class="flex-1 bg-purple-100 text-purple-700 py-2 rounded-lg hover:bg-purple-200"><i class="fas fa-sticky-note mr-1"></i> Notes</button>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

{#if showNotesModal && selectedAppointment}
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
    <div class="bg-white rounded-xl max-w-lg w-full p-6">
      <div class="flex justify-between items-start mb-4">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">Appointment Notes</h3>
          <p class="text-sm text-gray-500">{selectedAppointment.patient_name} - {formatDate(selectedAppointment.start_time)}</p>
        </div>
        <button on:click={() => showNotesModal = false} class="text-gray-400 hover:text-gray-600"><i class="fas fa-times text-xl"></i></button>
      </div>
      <textarea bind:value={notesInput} placeholder="Add notes (symptoms, diagnosis, prescription, follow-up...)" rows="6" class="input w-full"></textarea>
      <div class="flex justify-end space-x-3 mt-4">
        <button on:click={() => showNotesModal = false} class="btn-secondary">Cancel</button>
        <button on:click={saveNotes} disabled={savingNotes} class="btn-primary">
          {#if savingNotes}<i class="fas fa-spinner fa-spin mr-2"></i>{/if}Save Notes
        </button>
      </div>
    </div>
  </div>
{/if}
