<script>
  import { onMount } from 'svelte'
  import { fly } from 'svelte/transition'

  let visible = false

  onMount(() => {
    // Add a small delay for animation
    setTimeout(() => {
      visible = true
    }, 100)
  })
</script>

<div class="flex h-full w-full flex-col items-center justify-center p-6">
  {#if visible}
    <div
      class="w-full max-w-lg rounded-lg bg-white p-8 text-center shadow-lg"
      in:fly={{ y: 20, duration: 300 }}
    >
      <div class="mb-8 text-green-500">
        <svg
          class="mx-auto h-24 w-24"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M5 13l4 4L19 7"
          />
        </svg>
      </div>

      <h1 class="mb-4 text-4xl font-bold text-gray-900">
        Appointment Confirmed!
      </h1>
      <p class="mb-8 text-lg text-gray-600">
        Your consultation has been successfully booked. We've sent a
        confirmation to your email/phone.
      </p>

      <div class="mt-8 rounded-lg bg-gray-50 p-4">
        <h3 class="mb-2 text-lg font-semibold text-gray-800">
          What's Next?
        </h3>
        <ul class="text-left text-gray-600">
          <li class="mb-2 flex items-start">
            <i class="fa-solid fa-check mt-1 mr-2 text-green-500"></i>
            <span
              >You'll receive detailed instructions before your appointment</span
            >
          </li>
          <li class="mb-2 flex items-start">
            <i class="fa-solid fa-check mt-1 mr-2 text-green-500"></i>
            <span
              >Please be available 5 minutes before your scheduled time</span
            >
          </li>
          <li class="flex items-start">
            <i class="fa-solid fa-check mt-1 mr-2 text-green-500"></i>
            <span>Have your patient details and medical history ready</span
            >
          </li>
        </ul>
      </div>

      <a
        href="/"
        class="mt-8 inline-block rounded-lg bg-[#d5c455] px-6 py-3 text-white transition-colors hover:bg-[#c0ae4e]"
      >
        Return to Home
      </a>
    </div>
  {/if}
</div>
