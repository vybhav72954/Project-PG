<script lang="ts">
  import { onMount } from 'svelte';
  import { booking } from '$lib/stores';
  import { get } from 'svelte/store';

  let bookingData = get(booking);
</script>

<svelte:head>
  <title>Payment Failed | Friends2health Homoeo Clinic</title>
</svelte:head>

<section class="py-16">
  <div class="container-custom">
    <div class="max-w-xl mx-auto text-center">
      <!-- Error Icon -->
      <div class="w-24 h-24 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-6">
        <i class="fas fa-times-circle text-5xl text-red-500"></i>
      </div>
      
      <h1 class="text-3xl font-bold text-gray-900 mb-2">Payment Failed</h1>
      <p class="text-gray-600 mb-8">
        Unfortunately, your payment could not be processed. Don't worry, no amount has been deducted from your account.
      </p>

      <!-- What to do -->
      <div class="bg-gray-50 rounded-xl p-6 mb-8 text-left">
        <h3 class="font-bold text-gray-900 mb-4">What can you do?</h3>
        <ul class="space-y-3 text-gray-600">
          <li class="flex items-start gap-3">
            <i class="fas fa-redo text-primary-500 mt-1"></i>
            <span>Try booking again with a different payment method</span>
          </li>
          <li class="flex items-start gap-3">
            <i class="fas fa-wifi text-primary-500 mt-1"></i>
            <span>Check your internet connection and try again</span>
          </li>
          <li class="flex items-start gap-3">
            <i class="fas fa-headset text-primary-500 mt-1"></i>
            <span>Contact us if the problem persists</span>
          </li>
        </ul>
      </div>

      <!-- Actions -->
      <div class="flex flex-col sm:flex-row gap-4 justify-center">
        <a href="/appointment" class="btn-primary">
          <i class="fas fa-redo mr-2"></i>
          Try Again
        </a>
        <a href="https://wa.me/919877505344" target="_blank" rel="noopener" class="btn-secondary">
          <i class="fab fa-whatsapp mr-2"></i>
          Contact Support
        </a>
      </div>

      <!-- Contact Info -->
      <div class="mt-8 text-gray-600">
        <p>Need immediate assistance?</p>
        <a href="tel:+919877505344" class="text-primary-600 hover:text-primary-700 font-medium">
          <i class="fas fa-phone mr-1"></i> +91-9877505344
        </a>
      </div>
    </div>
  </div>
</section>
