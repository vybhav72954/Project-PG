<script lang="ts">
  import { onMount } from 'svelte';
  import { booking, type BookingState } from '$lib/stores';
  import { get } from 'svelte/store';
  import { goto } from '$app/navigation';
  import { siteConfig } from '$lib/config';

  const { contact } = siteConfig;

  let bookingData: BookingState;
  let loading = true;

  onMount(() => {
    bookingData = get(booking);

    // If no booking data, redirect to appointment page
    if (!bookingData.appointmentId) {
      goto('/appointment');
      return;
    }

    loading = false;
  });

  function formatDate(dateStr: string | null): string {
    if (!dateStr) return '';
    const date = new Date(dateStr);
    return date.toLocaleDateString('en-IN', {
      weekday: 'long',
      day: 'numeric',
      month: 'long',
      year: 'numeric'
    });
  }

  function formatTime(time: string | null): string {
    if (!time) return '';
    const [hours, minutes] = time.split(':');
    const hour = parseInt(hours);
    const ampm = hour >= 12 ? 'PM' : 'AM';
    const hour12 = hour % 12 || 12;
    return `${hour12}:${minutes} ${ampm}`;
  }

  function getAmount(): number {
    return bookingData.consultationType === 'video' ? 999 : 799;
  }

  function handlePrint() {
    window.print();
  }
</script>

<svelte:head>
  <title>Booking Confirmed | Friends2health Homoeo Clinic</title>
</svelte:head>

<section class="py-16">
  <div class="container-custom">
    {#if loading}
      <div class="text-center py-12">
        <i class="fas fa-spinner fa-spin text-3xl text-primary-500"></i>
      </div>
    {:else}
      <div class="max-w-2xl mx-auto">
        <!-- Success Icon -->
        <div class="text-center mb-8">
          <div class="w-24 h-24 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-6">
            <i class="fas fa-check-circle text-5xl text-green-500"></i>
          </div>
          <h1 class="text-3xl font-bold text-gray-900 mb-2">Booking Confirmed!</h1>
          <p class="text-gray-600">Your appointment has been successfully scheduled</p>
        </div>

        <!-- Appointment Details Card -->
        <div class="card p-8 mb-6">
          <h2 class="text-lg font-bold text-gray-900 mb-4 flex items-center gap-2">
            <i class="fas fa-calendar-check text-primary-500"></i>
            Appointment Details
          </h2>

          <div class="space-y-4">
            <div class="flex justify-between py-2 border-b border-gray-100">
              <span class="text-gray-600">Patient Name</span>
              <span class="font-medium text-gray-900">{bookingData.name}</span>
            </div>
            <div class="flex justify-between py-2 border-b border-gray-100">
              <span class="text-gray-600">Consultation Type</span>
              <span class="font-medium text-gray-900">
                {bookingData.consultationType === 'video' ? 'Video Call' : 'Voice Call'}
              </span>
            </div>
            <div class="flex justify-between py-2 border-b border-gray-100">
              <span class="text-gray-600">Date</span>
              <span class="font-medium text-gray-900">{formatDate(bookingData.selectedDate)}</span>
            </div>
            <div class="flex justify-between py-2 border-b border-gray-100">
              <span class="text-gray-600">Time</span>
              <span class="font-medium text-gray-900">{formatTime(bookingData.selectedTime)}</span>
            </div>
            <div class="flex justify-between py-2">
              <span class="text-gray-600">Amount Paid</span>
              <span class="font-bold text-primary-600">₹{getAmount()}</span>
            </div>
          </div>
        </div>

        <!-- Meeting Link (Video) or Call Info (Voice) -->
        {#if bookingData.consultationType === 'video' && bookingData.meetLink}
          <div class="bg-primary-50 border border-primary-200 rounded-xl p-6 mb-6">
            <h3 class="font-bold text-gray-900 mb-2 flex items-center gap-2">
              <i class="fas fa-video text-primary-600"></i>
              Meeting Link
            </h3>
            <p class="text-sm text-gray-600 mb-3">
              Join the consultation at your scheduled time using this link:
            </p>
            <a
              href={bookingData.meetLink}
              target="_blank"
              rel="noopener"
              class="inline-flex items-center gap-2 text-primary-600 hover:text-primary-700 font-medium break-all"
            >
              {bookingData.meetLink}
              <i class="fas fa-external-link-alt text-sm"></i>
            </a>
          </div>
        {:else if bookingData.consultationType === 'voice'}
          <div class="bg-blue-50 border border-blue-200 rounded-xl p-6 mb-6">
            <h3 class="font-bold text-gray-900 mb-2 flex items-center gap-2">
              <i class="fas fa-phone-alt text-blue-600"></i>
              Voice Call Consultation
            </h3>
            <p class="text-gray-600 mb-3">
              Dr. Aditi will call you at your registered phone number:
            </p>
            <p class="text-xl font-bold text-blue-600">
              <i class="fas fa-mobile-alt mr-2"></i>
              {bookingData.phone}
            </p>
            <p class="text-sm text-gray-500 mt-3">
              Please ensure your phone is reachable at the scheduled time.
            </p>
          </div>
        {/if}

        <!-- What's Next -->
        <div class="bg-gray-50 rounded-xl p-6 mb-8">
          <h3 class="font-bold text-gray-900 mb-4">What's Next?</h3>
          <ul class="space-y-3 text-gray-600">
            <li class="flex items-start gap-3">
              <i class="fas fa-envelope text-primary-500 mt-1"></i>
              <span>A confirmation email has been sent to <strong>{bookingData.email}</strong></span>
            </li>
            {#if bookingData.consultationType === 'video'}
              <li class="flex items-start gap-3">
                <i class="fas fa-clock text-primary-500 mt-1"></i>
                <span>Please join the video call 5 minutes before the scheduled time</span>
              </li>
            {:else}
              <li class="flex items-start gap-3">
                <i class="fas fa-phone text-primary-500 mt-1"></i>
                <span>Keep your phone reachable at the scheduled time - Dr. Aditi will call you</span>
              </li>
            {/if}
            <li class="flex items-start gap-3">
              <i class="fas fa-file-medical text-primary-500 mt-1"></i>
              <span>Keep your medical history and current medications handy</span>
            </li>
          </ul>
        </div>

        <!-- Actions -->
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <button on:click={handlePrint} class="btn-secondary">
            <i class="fas fa-print mr-2"></i>
            Print Details
          </button>
          <a href="/" class="btn-primary text-center">
            <i class="fas fa-home mr-2"></i>
            Back to Home
          </a>
        </div>

        <!-- Contact -->
        <div class="text-center mt-8 text-gray-600">
          <p>Need help? Contact us:</p>
          <div class="flex justify-center gap-4 mt-2">
            <a href="tel:{contact.phoneRaw}" class="text-primary-600 hover:text-primary-700">
              <i class="fas fa-phone mr-1"></i> {contact.phone}
            </a>
            <a href="https://wa.me/{contact.whatsapp}" class="text-primary-600 hover:text-primary-700">
              <i class="fab fa-whatsapp mr-1"></i> WhatsApp
            </a>
          </div>
        </div>
      </div>
    {/if}
  </div>
</section>
