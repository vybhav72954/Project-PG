<script lang="ts">
  import { fly } from 'svelte/transition'
  import { superForm } from 'sveltekit-superforms/client'
  import { zodClient } from 'sveltekit-superforms/adapters'
  import type { PageData } from './$types'
  import Calendar from './Calendar.svelte'
  import TimeSlots from './TimeSlots.svelte'
  import { onMount } from 'svelte'
  import { appointmentSchema } from '@/AppointmentSchema'
  // import SuperDebug from 'sveltekit-superforms/client/SuperDebug.svelte'

  export let data: PageData

  // Current stage of the form
  let currentStage = 1
  const totalStages = 3

  // OTP verification state
  let showOtpInput = false
  let otpTimer = 60
  let timerInterval: ReturnType<typeof setInterval> | null = null
  let otpValue = ''

  // Consultation options
  const consultationTypes = [
    {
      id: 'video',
      title: 'Video Conferencing',
      price: '₹999',
      icon: 'fa-solid fa-video'
    },
    {
      id: 'voice',
      title: 'Voice Call',
      price: '₹799',
      icon: 'fa-solid fa-phone',
      note: 'We will Call'
    }
  ]
  let selectedConsultation: string | null = null

  // Calendar and time slot selections
  let selectedDate: Date | null = null
  let selectedTimeSlot: string | null = null

  // Initialize the form with Zod schema from server
  const { form, errors, enhance, submitting } = superForm(data.form, {
    validators: zodClient(appointmentSchema),
    // {
    //   firstName: (firstName) => firstName.length < 3 ? 'Name must be at least 3 characters' : null,
    // }
    validationMethod: 'oninput',
    onSubmit: () => {
      // For demonstration, randomly choose success or failure
      const success = Math.random() > 0.3 // 70% chance of success

      if (success) {
        // Simulate successful submission
        setTimeout(() => {
          window.location.href = '/appointment/confirmed'
        }, 1000)
      } else {
        // Simulate failed submission
        setTimeout(() => {
          window.location.href = '/appointment/failed'
        }, 1000)
      }

      // Prevent actual form submission for this demo
      return { cancel: true }
    }
  })

  // Function to start OTP timer
  function startOtpTimer() {
    otpTimer = 60
    if (timerInterval) clearInterval(timerInterval)

    timerInterval = setInterval(() => {
      otpTimer--
      if (otpTimer <= 0 && timerInterval) {
        clearInterval(timerInterval)
      }
    }, 1000)
  }

  // Function to handle OTP verification
  function handleVerifyClick() {
    if (!showOtpInput) {
      showOtpInput = true
      startOtpTimer()
    } else {
      // Verify OTP
      if (otpValue.length === 6) {
        // For demo purposes, any 6-digit code works
        advanceStage()
      }
    }
  }

  // Function to resend OTP
  function resendOtp() {
    otpValue = ''
    startOtpTimer()
    // In a real app, would call API to resend OTP
  }

  // Go to next stage
  function advanceStage() {
    if (currentStage < totalStages) {
      currentStage++
    }
  }

  // Go to previous stage
  function goBack() {
    if (currentStage > 1) {
      currentStage--
    }
  }

  // Select consultation type
  function selectConsultation(type: string) {
    selectedConsultation = type
    advanceStage()
  }

  // Handle final form submission
  function handleSubmit() {
    // Form validation happens via Superforms
  }

  // Cleanup on component unmount
  onMount(() => {
    return () => {
      if (timerInterval) clearInterval(timerInterval)
    }
  })
</script>

<div class="flex h-full w-full flex-col items-center justify-center p-6">
  <!-- <SuperDebug data={form} /> -->

  <!-- Form progress indicator -->
  <div class="mb-8 flex w-full max-w-lg items-center justify-between">
    {#each Array(totalStages) as _, i}
      <div class="flex flex-col items-center">
        <div
          class="flex h-10 w-10 items-center justify-center rounded-full {i +
            1 <=
          currentStage
            ? 'bg-[#d5c455] text-white'
            : 'bg-gray-200 text-gray-600'}"
        >
          {i + 1}
        </div>
        <span class="mt-2 text-xs text-gray-500">
          Stage {i + 1}
        </span>
      </div>

      {#if i < totalStages - 1}
        <div class="h-1 w-16 bg-gray-200">
          <div
            class="h-full bg-[#d5c455]"
            style="width: {i + 1 < currentStage ? '100%' : '0%'}"
          ></div>
        </div>
      {/if}
    {/each}
  </div>

  <!-- Form container -->
  <div
    class="w-full max-w-lg rounded-lg bg-white p-8 shadow-lg"
    in:fly={{ y: 20, duration: 300 }}
  >
    <form method="POST" use:enhance>
      <!-- Stage 1: Patient Information -->
      {#if currentStage === 1}
        <div class="mb-6 text-center">
          <h2 class="text-2xl font-bold text-gray-900">
            Patient Information
          </h2>
          <p class="text-gray-600">Let's start with your basic details</p>
        </div>

        <div class="mb-4">
          <label
            for="name"
            class="mb-2 block text-sm font-medium text-gray-700"
            >Full Name</label
          >
          <input
            type="text"
            id="name"
            bind:value={$form.name}
            class="w-full rounded-lg border border-gray-300 p-3 shadow-sm focus:border-[#d5c455] focus:outline-none"
            placeholder="Enter your full name"
            required
          />
          {#if $errors.name}
            <p class="mt-1 text-sm text-red-600">{$errors.name}</p>
          {/if}
        </div>

        <div class="mb-4">
          <label
            for="contact"
            class="mb-2 block text-sm font-medium text-gray-700"
            >Email or Mobile Number</label
          >
          <input
            type="text"
            id="contact"
            bind:value={$form.contact}
            class="w-full rounded-lg border border-gray-300 p-3 shadow-sm focus:border-[#d5c455] focus:outline-none"
            placeholder="Enter email or mobile number"
            required
          />
          {#if $errors.contact}
            <p class="mt-1 text-sm text-red-600">{$errors.contact}</p>
          {/if}
        </div>

        <!-- OTP verification section -->
        {#if !showOtpInput}
          <button
            type="button"
            on:click={handleVerifyClick}
            class="mt-4 w-full rounded-lg bg-[#d5c455] px-6 py-3 text-center font-medium text-white shadow-md transition-shadow hover:bg-[#c0ae4e]"
          >
            Verify Through OTP
          </button>
        {:else}
          <div class="mt-4 mb-4">
            <label
              for="otp"
              class="mb-2 block text-sm font-medium text-gray-700"
              >Enter OTP</label
            >
            <div class="flex gap-3">
              <input
                type="text"
                id="otp"
                bind:value={otpValue}
                class="w-full rounded-lg border border-gray-300 p-3 shadow-sm focus:border-[#d5c455] focus:outline-none"
                placeholder="6-digit OTP"
                maxlength="6"
                required
              />
              <button
                type="button"
                on:click={handleVerifyClick}
                class="rounded-lg bg-[#d5c455] px-4 py-3 text-white shadow-md transition-shadow hover:bg-[#c0ae4e]"
                disabled={otpValue.length !== 6}
              >
                Verify
              </button>
            </div>
            <div class="mt-2 flex justify-between text-sm">
              <span
                >OTP expires in: {Math.floor(otpTimer / 60)}:{otpTimer %
                  60 <
                10
                  ? '0'
                  : ''}{otpTimer % 60}</span
              >
              {#if otpTimer <= 0}
                <button
                  type="button"
                  on:click={resendOtp}
                  class="text-[#d5c455] hover:underline"
                >
                  Resend OTP
                </button>
              {/if}
            </div>
          </div>
        {/if}
      {/if}

      <!-- Stage 2: Consultation Type -->
      {#if currentStage === 2}
        <div class="mb-6 text-center">
          <h2 class="text-2xl font-bold text-gray-900">
            Consultation Type
          </h2>
          <p class="text-gray-600">
            Choose your preferred consultation method
          </p>
        </div>

        <div class="flex flex-col gap-4">
          {#each consultationTypes as type}
            <button
              type="button"
              on:click={() => selectConsultation(type.id)}
              class="flex items-center justify-between rounded-lg border border-gray-300 p-4 text-left shadow-sm transition-all hover:border-[#d5c455] hover:shadow-md"
            >
              <div class="flex items-center">
                <div
                  class="mr-4 flex h-12 w-12 items-center justify-center rounded-full bg-[#d5c455] text-white"
                >
                  <i class={type.icon}></i>
                </div>
                <div>
                  <h3 class="font-medium">{type.title}</h3>
                  {#if type.note}
                    <p class="text-sm text-gray-500">{type.note}</p>
                  {/if}
                </div>
              </div>
              <span class="font-medium">{type.price}</span>
            </button>
          {/each}
        </div>

        <p class="mt-6 text-center text-sm text-gray-500">
          Consultation charges include follow-up within 7 days
        </p>

        <button
          type="button"
          on:click={goBack}
          class="mt-6 flex items-center text-[#d5c455] hover:underline"
        >
          <i class="fa-solid fa-arrow-left mr-2"></i> Back
        </button>
      {/if}

      <!-- Stage 3: Date and Time Selection -->
      {#if currentStage === 3}
        <div class="mb-6 text-center">
          <h2 class="text-2xl font-bold text-gray-900">
            Schedule Appointment
          </h2>
          <p class="text-gray-600">Select your preferred date and time</p>
        </div>

        <!-- Calendar component - Fixed the label association issue -->
        <div class="mb-6">
          <span
            id="calendar-label"
            class="mb-2 block text-sm font-medium text-gray-700"
            >Select Date</span
          >
          <div aria-labelledby="calendar-label">
            <Calendar bind:selectedDate />
          </div>
          <input
            type="hidden"
            name="appointmentDate"
            bind:value={$form.appointmentDate}
          />
          {#if $errors.appointmentDate}
            <p class="mt-1 text-sm text-red-600">
              {$errors.appointmentDate}
            </p>
          {/if}
        </div>

        <!-- Time slots component - Fixed the label association issue -->
        <div class="mb-6">
          <span
            id="timeslots-label"
            class="mb-2 block text-sm font-medium text-gray-700"
            >Select Time</span
          >
          <div aria-labelledby="timeslots-label">
            <TimeSlots bind:selectedTimeSlot {selectedDate} />
          </div>
          <input
            type="hidden"
            name="appointmentTime"
            bind:value={$form.appointmentTime}
          />
          {#if $errors.appointmentTime}
            <p class="mt-1 text-sm text-red-600">
              {$errors.appointmentTime}
            </p>
          {/if}
        </div>

        <div class="mt-8 flex justify-between">
          <button
            type="button"
            on:click={goBack}
            class="flex items-center text-[#d5c455] hover:underline"
          >
            <i class="fa-solid fa-arrow-left mr-2"></i> Back
          </button>

          <button
            type="submit"
            class="rounded-lg bg-[#d5c455] px-6 py-3 text-center font-medium text-white shadow-md transition-all hover:bg-[#c0ae4e] disabled:opacity-50"
            disabled={!selectedDate || !selectedTimeSlot || $submitting}
          >
            {$submitting ? 'Processing...' : 'Book My Appointment'}
          </button>
        </div>
      {/if}
    </form>
  </div>
</div>
