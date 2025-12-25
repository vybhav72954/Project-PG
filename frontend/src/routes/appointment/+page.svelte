<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { booking, type BookingState } from '$lib/stores';
  import { api, type DaySlots, type TimeSlot } from '$lib/api';
  import { get } from 'svelte/store';

  // Form state
  let step = 1;
  let loading = false;
  let error = '';

  // Patient info
  let name = '';
  let email = '';
  let phone = '';

  // Consultation type
  let consultationType: 'video' | 'voice' = 'video';
  const pricing = {
    video: 999,
    voice: 799
  };

  // Date/Time selection
  let availableSlots: DaySlots[] = [];
  let selectedDate = '';
  let selectedTime = '';
  let loadingSlots = false;

  // Razorpay
  let razorpayKeyId = '';

  onMount(async () => {
    // Load slots for next 14 days
    await loadSlots();
    
    // Get Razorpay key from config
    try {
      const response = await api.getConfig();
      if (response.success && response.data) {
        razorpayKeyId = response.data.razorpay_key_id;
      }
    } catch (e) {
      console.error('Failed to load config:', e);
    }
  });

  async function loadSlots() {
    loadingSlots = true;
    try {
      const today = new Date();
      const startDate = today.toISOString().split('T')[0];
      const endDate = new Date(today.getTime() + 14 * 24 * 60 * 60 * 1000).toISOString().split('T')[0];
      
      const response = await api.getAvailableSlots(startDate, endDate);
      if (response.success && response.data) {
        availableSlots = response.data;
      }
    } catch (e) {
      console.error('Failed to load slots:', e);
    } finally {
      loadingSlots = false;
    }
  }

  function nextStep() {
    if (step === 1 && (!name || !email || !phone)) {
      error = 'Please fill in all fields';
      return;
    }
    if (step === 1 && !isValidEmail(email)) {
      error = 'Please enter a valid email address';
      return;
    }
    if (step === 1 && !isValidPhone(phone)) {
      error = 'Please enter a valid phone number';
      return;
    }
    if (step === 3 && (!selectedDate || !selectedTime)) {
      error = 'Please select a date and time';
      return;
    }
    error = '';
    step++;
  }

  function prevStep() {
    error = '';
    step--;
  }

  function isValidEmail(email: string): boolean {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
  }

  function isValidPhone(phone: string): boolean {
    return /^[6-9]\d{9}$/.test(phone.replace(/\D/g, ''));
  }

  function formatDate(dateStr: string): string {
    const date = new Date(dateStr);
    return date.toLocaleDateString('en-IN', { weekday: 'short', day: 'numeric', month: 'short' });
  }

  function formatTime(time: string): string {
    const [hours, minutes] = time.split(':');
    const hour = parseInt(hours);
    const ampm = hour >= 12 ? 'PM' : 'AM';
    const hour12 = hour % 12 || 12;
    return `${hour12}:${minutes} ${ampm}`;
  }

  function getAvailableSlotsForDate(date: string): TimeSlot[] {
    const day = availableSlots.find(d => d.date === date);
    return day?.slots?.filter(s => s.available) || [];
  }

  async function handlePayment() {
    loading = true;
    error = '';

    try {
      // Create booking
      const bookingResponse = await api.createBooking({
        name,
        email,
        phone: phone.replace(/\D/g, ''),
        consultation_type: consultationType,
        date: selectedDate,
        time_slot: selectedTime
      });

      if (!bookingResponse.success || !bookingResponse.data) {
        throw new Error(bookingResponse.error || 'Failed to create booking');
      }

      const { order_id, amount, appointment_id } = bookingResponse.data;

      // Store booking info
      booking.setPatientInfo(name, email, phone);
      booking.setConsultationType(consultationType);
      booking.setDateTime(selectedDate, selectedTime);
      booking.setOrderDetails(appointment_id, order_id);

      // Check if we have Razorpay configured
      if (!razorpayKeyId || razorpayKeyId === '') {
        // Development mode - simulate payment
        const verifyResponse = await api.verifyPayment({
          razorpay_order_id: order_id,
          razorpay_payment_id: 'dev_' + Date.now(),
          razorpay_signature: 'dev_signature',
          appointment_id
        });

        if (verifyResponse.success) {
          booking.setMeetLink(verifyResponse.data?.meet_link || '');
          goto('/appointment/confirmed');
        } else {
          throw new Error(verifyResponse.error || 'Payment verification failed');
        }
        return;
      }

      // Open Razorpay checkout
      const options = {
        key: razorpayKeyId,
        amount: amount,
        currency: 'INR',
        name: 'Friends2health Homoeo Clinic',
        description: `${consultationType === 'video' ? 'Video' : 'Voice'} Consultation`,
        order_id: order_id,
        handler: async function(response: any) {
          // Verify payment
          const verifyResponse = await api.verifyPayment({
            razorpay_order_id: response.razorpay_order_id,
            razorpay_payment_id: response.razorpay_payment_id,
            razorpay_signature: response.razorpay_signature,
            appointment_id
          });

          if (verifyResponse.success) {
            booking.setMeetLink(verifyResponse.data?.meet_link || '');
            goto('/appointment/confirmed');
          } else {
            error = verifyResponse.error || 'Payment verification failed';
            goto('/appointment/failed');
          }
        },
        prefill: {
          name,
          email,
          contact: phone
        },
        theme: {
          color: '#5F8575'
        },
        modal: {
          ondismiss: function() {
            loading = false;
          }
        }
      };

      const rzp = new (window as any).Razorpay(options);
      rzp.open();
    } catch (e: any) {
      error = e.message || 'An error occurred';
      loading = false;
    }
  }
</script>

<svelte:head>
  <title>Book Appointment | Friends2health Homoeo Clinic</title>
</svelte:head>

<!-- Hero Section -->
<section class="bg-gradient-hero py-8">
  <div class="container-custom">
    <div class="text-center mb-8">
      <h1 class="text-3xl md:text-4xl font-bold text-gray-900 mb-2">
        Book Your Consultation
      </h1>
      <p class="text-gray-600">Schedule an appointment with Dr. Aditi Singh</p>
    </div>

    <!-- Progress Steps -->
    <div class="flex justify-center mb-8">
      <div class="flex items-center gap-4">
        {#each [1, 2, 3, 4] as s}
          <div class="flex items-center">
            <div class="w-10 h-10 rounded-full flex items-center justify-center font-semibold transition-all
              {step >= s ? 'bg-primary-500 text-white' : 'bg-gray-200 text-gray-500'}">
              {#if step > s}
                <i class="fas fa-check"></i>
              {:else}
                {s}
              {/if}
            </div>
            {#if s < 4}
              <div class="w-12 h-1 mx-2 rounded {step > s ? 'bg-primary-500' : 'bg-gray-200'}"></div>
            {/if}
          </div>
        {/each}
      </div>
    </div>
  </div>
</section>

<!-- Form Section -->
<section class="py-8 pb-16">
  <div class="container-custom">
    <div class="max-w-2xl mx-auto">
      {#if error}
        <div class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg mb-6 flex items-center gap-2">
          <i class="fas fa-exclamation-circle"></i>
          {error}
        </div>
      {/if}

      <!-- Step 1: Patient Info -->
      {#if step === 1}
        <div class="card p-8 animate-fadeIn">
          <h2 class="text-xl font-bold text-gray-900 mb-6">
            <i class="fas fa-user text-primary-500 mr-2"></i>
            Your Information
          </h2>
          
          <div class="space-y-4">
            <div>
              <label for="name" class="block text-sm font-medium text-gray-700 mb-1">Full Name *</label>
              <input 
                id="name"
                type="text" 
                bind:value={name} 
                class="input" 
                placeholder="Enter your full name"
              />
            </div>
            
            <div>
              <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email Address *</label>
              <input 
                id="email"
                type="email" 
                bind:value={email} 
                class="input" 
                placeholder="your@email.com"
              />
            </div>
            
            <div>
              <label for="phone" class="block text-sm font-medium text-gray-700 mb-1">Phone Number *</label>
              <div class="flex">
                <span class="inline-flex items-center px-4 border border-r-0 border-gray-300 bg-gray-50 text-gray-500 rounded-l-md">
                  +91
                </span>
                <input 
                  id="phone"
                  type="tel" 
                  bind:value={phone} 
                  class="input rounded-l-none" 
                  placeholder="9876543210"
                  maxlength="10"
                />
              </div>
            </div>
          </div>

          <div class="mt-8 flex justify-end">
            <button on:click={nextStep} class="btn-primary">
              Continue
              <i class="fas fa-arrow-right ml-2"></i>
            </button>
          </div>
        </div>
      {/if}

      <!-- Step 2: Consultation Type -->
      {#if step === 2}
        <div class="card p-8 animate-fadeIn">
          <h2 class="text-xl font-bold text-gray-900 mb-6">
            <i class="fas fa-video text-primary-500 mr-2"></i>
            Select Consultation Type
          </h2>
          
          <div class="grid md:grid-cols-2 gap-4">
            <button 
              on:click={() => consultationType = 'video'}
              class="p-6 rounded-xl border-2 transition-all text-left
                {consultationType === 'video' 
                  ? 'border-primary-500 bg-primary-50' 
                  : 'border-gray-200 hover:border-primary-300'}"
            >
              <div class="flex items-center justify-between mb-4">
                <div class="w-12 h-12 bg-primary-100 rounded-lg flex items-center justify-center">
                  <i class="fas fa-video text-xl text-primary-600"></i>
                </div>
                <div class="w-6 h-6 rounded-full border-2 flex items-center justify-center
                  {consultationType === 'video' ? 'border-primary-500 bg-primary-500' : 'border-gray-300'}">
                  {#if consultationType === 'video'}
                    <i class="fas fa-check text-white text-xs"></i>
                  {/if}
                </div>
              </div>
              <h3 class="font-bold text-gray-900 mb-1">Video Consultation</h3>
              <p class="text-sm text-gray-600 mb-3">Face-to-face consultation via video call</p>
              <p class="text-2xl font-bold text-primary-600">₹{pricing.video}</p>
            </button>

            <button 
              on:click={() => consultationType = 'voice'}
              class="p-6 rounded-xl border-2 transition-all text-left
                {consultationType === 'voice' 
                  ? 'border-primary-500 bg-primary-50' 
                  : 'border-gray-200 hover:border-primary-300'}"
            >
              <div class="flex items-center justify-between mb-4">
                <div class="w-12 h-12 bg-primary-100 rounded-lg flex items-center justify-center">
                  <i class="fas fa-phone text-xl text-primary-600"></i>
                </div>
                <div class="w-6 h-6 rounded-full border-2 flex items-center justify-center
                  {consultationType === 'voice' ? 'border-primary-500 bg-primary-500' : 'border-gray-300'}">
                  {#if consultationType === 'voice'}
                    <i class="fas fa-check text-white text-xs"></i>
                  {/if}
                </div>
              </div>
              <h3 class="font-bold text-gray-900 mb-1">Voice Consultation</h3>
              <p class="text-sm text-gray-600 mb-3">Audio consultation via phone call</p>
              <p class="text-2xl font-bold text-primary-600">₹{pricing.voice}</p>
            </button>
          </div>

          <div class="mt-8 flex justify-between">
            <button on:click={prevStep} class="btn-secondary">
              <i class="fas fa-arrow-left mr-2"></i>
              Back
            </button>
            <button on:click={nextStep} class="btn-primary">
              Continue
              <i class="fas fa-arrow-right ml-2"></i>
            </button>
          </div>
        </div>
      {/if}

      <!-- Step 3: Date & Time -->
      {#if step === 3}
        <div class="card p-8 animate-fadeIn">
          <h2 class="text-xl font-bold text-gray-900 mb-6">
            <i class="fas fa-calendar text-primary-500 mr-2"></i>
            Select Date & Time
          </h2>
          
          {#if loadingSlots}
            <div class="text-center py-8">
              <i class="fas fa-spinner fa-spin text-2xl text-primary-500 mb-2"></i>
              <p class="text-gray-600">Loading available slots...</p>
            </div>
          {:else}
            <!-- Date Selection -->
            <div class="mb-6">
              <p class="block text-sm font-medium text-gray-700 mb-3">Select Date</p>
              <div class="grid grid-cols-4 sm:grid-cols-5 gap-2">
                {#each availableSlots as day}
                  {@const hasSlots = day.slots?.some(s => s.available) ?? false}
                  <button 
                    on:click={() => { selectedDate = day.date; selectedTime = ''; }}
                    disabled={!hasSlots || day.is_weekend}
                    class="p-3 rounded-lg text-center transition-all
                      {selectedDate === day.date 
                        ? 'bg-primary-500 text-white' 
                        : hasSlots && !day.is_weekend
                          ? 'bg-white border border-gray-200 hover:border-primary-300' 
                          : 'bg-gray-100 text-gray-400 cursor-not-allowed'}"
                  >
                    <div class="text-xs">{formatDate(day.date).split(',')[0]}</div>
                    <div class="font-bold">{new Date(day.date).getDate()}</div>
                  </button>
                {/each}
              </div>
            </div>

            <!-- Time Selection -->
            {#if selectedDate}
              {@const slots = getAvailableSlotsForDate(selectedDate)}
              <div class="mb-6">
                <p class="block text-sm font-medium text-gray-700 mb-3">Select Time</p>
                {#if slots.length === 0}
                  <p class="text-gray-500 text-center py-4">No slots available for this date</p>
                {:else}
                  <div class="grid grid-cols-3 sm:grid-cols-4 gap-2">
                    {#each slots as slot}
                      <button 
                        on:click={() => selectedTime = slot.time}
                        class="p-3 rounded-lg text-center transition-all
                          {selectedTime === slot.time 
                            ? 'bg-primary-500 text-white' 
                            : 'bg-white border border-gray-200 hover:border-primary-300'}"
                      >
                        {formatTime(slot.time)}
                      </button>
                    {/each}
                  </div>
                {/if}
              </div>
            {/if}
          {/if}

          <div class="mt-8 flex justify-between">
            <button on:click={prevStep} class="btn-secondary">
              <i class="fas fa-arrow-left mr-2"></i>
              Back
            </button>
            <button on:click={nextStep} class="btn-primary" disabled={!selectedDate || !selectedTime}>
              Continue
              <i class="fas fa-arrow-right ml-2"></i>
            </button>
          </div>
        </div>
      {/if}

      <!-- Step 4: Summary & Payment -->
      {#if step === 4}
        <div class="card p-8 animate-fadeIn">
          <h2 class="text-xl font-bold text-gray-900 mb-6">
            <i class="fas fa-receipt text-primary-500 mr-2"></i>
            Booking Summary
          </h2>
          
          <div class="bg-gray-50 rounded-xl p-6 mb-6">
            <div class="space-y-4">
              <div class="flex justify-between">
                <span class="text-gray-600">Patient Name</span>
                <span class="font-medium text-gray-900">{name}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">Email</span>
                <span class="font-medium text-gray-900">{email}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">Phone</span>
                <span class="font-medium text-gray-900">+91-{phone}</span>
              </div>
              <hr class="border-gray-200" />
              <div class="flex justify-between">
                <span class="text-gray-600">Consultation Type</span>
                <span class="font-medium text-gray-900">{consultationType === 'video' ? 'Video' : 'Voice'} Call</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">Date</span>
                <span class="font-medium text-gray-900">{formatDate(selectedDate)}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">Time</span>
                <span class="font-medium text-gray-900">{formatTime(selectedTime)}</span>
              </div>
              <hr class="border-gray-200" />
              <div class="flex justify-between text-lg">
                <span class="font-semibold text-gray-900">Total Amount</span>
                <span class="font-bold text-primary-600">₹{pricing[consultationType]}</span>
              </div>
            </div>
          </div>

          <div class="mt-8 flex justify-between">
            <button on:click={prevStep} class="btn-secondary" disabled={loading}>
              <i class="fas fa-arrow-left mr-2"></i>
              Back
            </button>
            <button on:click={handlePayment} class="btn-primary" disabled={loading}>
              {#if loading}
                <i class="fas fa-spinner fa-spin mr-2"></i>
                Processing...
              {:else}
                <i class="fas fa-lock mr-2"></i>
                Pay ₹{pricing[consultationType]}
              {/if}
            </button>
          </div>

          <p class="text-center text-sm text-gray-500 mt-4">
            <i class="fas fa-shield-alt mr-1"></i>
            Secure payment powered by Razorpay
          </p>
        </div>
      {/if}
    </div>
  </div>
</section>
