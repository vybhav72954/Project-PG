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
      <div class="mb-8 text-red-500">
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
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
      </div>
      <h1 class="mb-4 text-4xl font-bold text-gray-900">Booking Failed</h1>
      <p class="mb-8 text-lg text-gray-600">
        We encountered an issue processing your request. Please try again
        later.
      </p>
      <div class="mt-8 rounded-lg bg-gray-50 p-4">
        <h3 class="mb-2 text-lg font-semibold text-gray-800">
          What to Do Next?
        </h3>
        <ul class="text-left text-gray-600">
          <li class="mb-2 flex items-start">
            <i class="fa-solid fa-exclamation mt-1 mr-2 text-red-500"></i>
            <span>Check your internet connection and refresh the page</span
            >
          </li>
          <li class="mb-2 flex items-start">
            <i class="fa-solid fa-exclamation mt-1 mr-2 text-red-500"></i>
            <span>Try booking at a different time slot</span>
          </li>
          <li class="flex items-start">
            <i class="fa-solid fa-exclamation mt-1 mr-2 text-red-500"></i>
            <span>Contact our support team if the issue persists</span>
          </li>
        </ul>
      </div>
      <a
        href="/appointment"
        class="mt-8 inline-block rounded-lg bg-[#d5c455] px-6 py-3 text-white transition-colors hover:bg-[#c0ae4e]"
      >
        Try Again
      </a>
    </div>
  {/if}
</div>
