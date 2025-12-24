<script lang="ts">
  import { onMount } from 'svelte';
  import { adminApi, type PatientSummary, type Appointment } from '$lib/adminApi';

  let patients: PatientSummary[] = [];
  let loading = true;
  let error = '';

  let selectedPatient: PatientSummary | null = null;
  let patientHistory: Appointment[] = [];
  let historyLoading = false;

  onMount(async () => {
    await loadPatients();
  });

  async function loadPatients() {
    loading = true;
    error = '';

    const response = await adminApi.getPatients();

    if (response.success && response.data) {
      patients = response.data;
    } else {
      error = response.error || 'Failed to load patients';
    }

    loading = false;
  }

  async function viewHistory(patient: PatientSummary) {
    selectedPatient = patient;
    historyLoading = true;

    const response = await adminApi.getPatientHistory(patient.email);

    if (response.success && response.data) {
      patientHistory = response.data;
    }

    historyLoading = false;
  }

  function closeHistory() {
    selectedPatient = null;
    patientHistory = [];
  }

  function formatDate(dateStr: string): string {
    const date = new Date(dateStr);
    return date.toLocaleDateString('en-IN', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    });
  }

  function formatDateTime(dateStr: string): string {
    const date = new Date(dateStr);
    return date.toLocaleString('en-IN', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  }

  function getStatusColor(status: string): string {
    switch (status) {
      case 'confirmed': return 'bg-green-100 text-green-800';
      case 'completed': return 'bg-blue-100 text-blue-800';
      case 'cancelled': return 'bg-red-100 text-red-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  }
</script>

<div>
  <div class="mb-6">
    <h1 class="text-2xl font-bold text-gray-900">Patients</h1>
    <p class="text-gray-600">View all patients and their appointment history</p>
  </div>

  {#if loading}
    <div class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-10 w-10 border-4 border-primary-500 border-t-transparent"></div>
    </div>
  {:else if error}
    <div class="bg-red-50 text-red-600 px-6 py-4 rounded-xl">
      <i class="fas fa-exclamation-circle mr-2"></i>
      {error}
    </div>
  {:else if patients.length === 0}
    <div class="bg-white rounded-xl p-12 text-center shadow-sm border border-gray-100">
      <i class="fas fa-users text-4xl text-gray-300 mb-4"></i>
      <p class="text-gray-500">No patients yet</p>
    </div>
  {:else}
    <div class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Patient</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider hidden md:table-cell">Contact</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Appointments</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider hidden md:table-cell">Last Visit</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          {#each patients as patient}
            <tr class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div>
                  <p class="font-medium text-gray-900">{patient.name}</p>
                  <p class="text-sm text-gray-500 md:hidden">{patient.phone}</p>
                </div>
              </td>
              <td class="px-6 py-4 hidden md:table-cell">
                <p class="text-gray-900">{patient.email}</p>
                <p class="text-sm text-gray-500">{patient.phone}</p>
              </td>
              <td class="px-6 py-4">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-sm font-medium bg-primary-100 text-primary-800">
                  {patient.appointment_count}
                </span>
              </td>
              <td class="px-6 py-4 hidden md:table-cell">
                <p class="text-gray-900">{formatDate(patient.last_appointment)}</p>
              </td>
              <td class="px-6 py-4">
                <button 
                  on:click={() => viewHistory(patient)}
                  class="text-primary-600 hover:text-primary-800"
                >
                  <i class="fas fa-history mr-1"></i>
                  View History
                </button>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div>

<!-- Patient History Modal -->
{#if selectedPatient}
  <div class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen px-4 pt-4 pb-20">
      <div class="fixed inset-0 bg-black bg-opacity-50" on:click={closeHistory} on:keydown={(e) => e.key === 'Escape' && closeHistory()} role="button" tabindex="0"></div>
      
      <div class="relative bg-white rounded-2xl max-w-2xl w-full max-h-[80vh] overflow-hidden shadow-xl">
        <!-- Header -->
        <div class="px-6 py-4 border-b bg-gray-50 flex items-center justify-between">
          <div>
            <h2 class="text-lg font-semibold text-gray-900">{selectedPatient.name}</h2>
            <p class="text-sm text-gray-500">{selectedPatient.email}</p>
          </div>
          <button on:click={closeHistory} class="text-gray-400 hover:text-gray-600">
            <i class="fas fa-times text-xl"></i>
          </button>
        </div>

        <!-- Content -->
        <div class="p-6 overflow-y-auto max-h-[60vh]">
          {#if historyLoading}
            <div class="flex justify-center py-8">
              <div class="animate-spin rounded-full h-8 w-8 border-4 border-primary-500 border-t-transparent"></div>
            </div>
          {:else if patientHistory.length === 0}
            <p class="text-center text-gray-500 py-8">No appointment history</p>
          {:else}
            <div class="space-y-4">
              {#each patientHistory as apt}
                <div class="border rounded-xl p-4">
                  <div class="flex justify-between items-start mb-2">
                    <div>
                      <p class="font-medium text-gray-900">{formatDateTime(apt.start_time)}</p>
                      <p class="text-sm text-gray-500 capitalize">{apt.consultation_type} Consultation</p>
                    </div>
                    <span class="px-2 py-1 text-xs font-medium rounded-full {getStatusColor(apt.status)}">
                      {apt.status.replace('_', ' ')}
                    </span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-500">Amount: ₹{apt.amount / 100}</span>
                    {#if apt.meet_link}
                      <a href={apt.meet_link} target="_blank" rel="noopener" class="text-blue-600 hover:underline">
                        <i class="fas fa-video mr-1"></i>Meeting Link
                      </a>
                    {/if}
                  </div>
                </div>
              {/each}
            </div>
          {/if}
        </div>

        <!-- Footer -->
        <div class="px-6 py-4 border-t bg-gray-50">
          <div class="flex justify-between text-sm">
            <span class="text-gray-500">First visit: {formatDate(selectedPatient.first_visit)}</span>
            <span class="text-gray-500">Total appointments: {selectedPatient.appointment_count}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
{/if}
