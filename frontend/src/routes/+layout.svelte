<script lang="ts">
  import '../app.css'
  import FooterContent from './FooterContent.svelte'
  import { preloadData } from '$app/navigation'

  // let { children } = $props();
  const dev_url = 'https://github.com/yoAeroA00'

  // Prefetch pages when the component mounts
  function prefetchPages() {
    preloadData('/appointment')
    preloadData('/about')
  }
</script>

<svelte:head>
  <link rel="preconnect" href="https://cdnjs.cloudflare.com" />
  <link
    rel="stylesheet"
    href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.7.2/css/all.min.css"
    integrity="sha512-Evv84Mr4kqVGRNSgIGL/F/aIDqQb7xQ2vcrdIwxfjThSH8CSR7PBEakCr51Ck+w+/U6swU2Im1vVX0SVk9ABhg=="
    crossorigin="anonymous"
    referrerpolicy="no-referrer"
  />
</svelte:head>

<!-- Call prefetch on component initialization -->
<svelte:window on:load={prefetchPages} />

<div class="relative flex min-h-screen flex-col">
  <!-- Container for nav and main that fills available space -->
  <div class="flex flex-1 flex-col lg:flex-row">
    <!-- Left Section: Navigation -->
    <nav
      class="relative flex w-full flex-col justify-center bg-gray-100 p-8 lg:w-2/5"
    >
      <div class="flex w-full justify-center lg:justify-center">
        <!-- Inner container preserves left text alignment -->
        <div class="text-left">
          <span class="subheading mb-2 block text-lg font-medium">
            Homeopathy Practitioner
          </span>
          <h1 class="mb-3 text-3xl font-bold sm:text-6xl">I'm</h1>
          <h1 class="mb-8 text-3xl font-bold sm:text-6xl">
            Dr Aditi Singh
          </h1>
          <!-- Navigation Links with prefetch when hovered - using flex-col for mobile -->
          <div
            class="flex flex-col space-y-4 sm:flex-row sm:space-y-0 sm:space-x-4"
          >
            <a
              href="/appointment"
              class="inline-block rounded-lg bg-[#d5c455] px-6 py-3 text-center text-white shadow-md transition-shadow hover:bg-[#c0ae4e]"
              on:mouseenter={() => preloadData('/appointment')}
            >
              Book Appointment
            </a>
            <a
              href="/about"
              class="inline-block rounded-lg border border-[#d5c455] bg-white px-6 py-3 text-center text-[#d5c455] shadow-md transition-shadow hover:bg-[#c0ae4e] hover:text-white"
              on:mouseenter={() => preloadData('/about')}
            >
              About Me
            </a>
          </div>
        </div>
      </div>
      <!-- Footer specifically for desktop - centered at bottom of left section -->
      <footer
        class="hidden w-full p-4 text-center lg:absolute lg:bottom-0 lg:left-0 lg:block"
      >
        <FooterContent {dev_url} />
      </footer>
    </nav>
    <!-- Right Section: Page Content -->
    <main class="w-full flex-1 p-0 lg:min-h-screen lg:w-3/5">
      <slot />
    </main>
  </div>
  <!-- Footer for mobile and tablet only -->
  <footer class="w-full text-center lg:hidden">
    <FooterContent {dev_url} />
  </footer>
</div>

<!-- {@render children()} -->
