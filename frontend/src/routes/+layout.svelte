<script lang="ts">
  import '../app.css';
  import { page } from '$app/stores';
  import { siteConfig } from '$lib/config';

  const { contact } = siteConfig;

  let mobileMenuOpen = false;

  const navLinks = [
    { href: '/', label: 'Home' },
    { href: '/about', label: 'About' },
    { href: '/testimonials', label: 'Testimonials' },
    { href: '/blog', label: 'Research & Blogs' },
    { href: '/appointment', label: 'Book Appointment' }
  ];

  function toggleMobileMenu() {
    mobileMenuOpen = !mobileMenuOpen;
  }

  function closeMobileMenu() {
    mobileMenuOpen = false;
  }

  $: isAdminRoute = $page.url.pathname.startsWith('/admin');
</script>

{#if isAdminRoute}
  <slot />
{:else}
  <header class="bg-white border-b border-gray-100 sticky top-0 z-50">
    <div class="bg-primary-500 text-white py-2 text-sm">
      <div class="container-custom flex justify-between items-center">
        <div class="flex items-center gap-6">
        <a href="tel:{contact.phoneRaw}" class="flex items-center gap-2 hover:text-primary-100 transition-colors">
          <i class="fas fa-phone text-xs"></i>
          <span class="hidden sm:inline">{contact.phone}</span>
        </a>
        <a href="mailto:{contact.email}" class="hidden md:flex items-center gap-2 hover:text-primary-100 transition-colors">
          <i class="fas fa-envelope text-xs"></i>
          <span>{contact.email}</span>
        </a>
        </div>
        <a href="https://wa.me/{contact.whatsapp}" target="_blank" rel="noopener" class="flex items-center gap-2 hover:text-primary-100 transition-colors">
          <i class="fab fa-whatsapp"></i>
          <span>WhatsApp</span>
        </a>
      </div>
    </div>

    <nav class="container-custom py-4">
      <div class="flex items-center justify-between">
        <a href="/" class="flex items-center gap-3">
          <div class="w-10 h-10 bg-primary-500 rounded-lg flex items-center justify-center">
            <span class="text-white font-bold text-sm">F2H</span>
          </div>
          <div class="hidden sm:block">
            <p class="font-bold text-gray-900 leading-tight">Friends2health</p>
            <p class="text-xs text-primary-600">Homoeo Clinic</p>
          </div>
        </a>

        <div class="hidden md:flex items-center gap-8">
          {#each navLinks as link}
            <a
              href={link.href}
              class="text-gray-700 hover:text-primary-600 transition-colors font-medium {$page.url.pathname === link.href ? 'text-primary-600' : ''}"
            >
              {link.label}
            </a>
          {/each}
        </div>

        <div class="hidden md:block">
          <a href="/appointment" class="btn-primary">Book Now</a>
        </div>

        <button
          class="md:hidden p-2 text-gray-700 hover:text-primary-600"
          on:click={toggleMobileMenu}
          aria-label="Toggle menu"
        >
          <i class="fas {mobileMenuOpen ? 'fa-times' : 'fa-bars'} text-xl"></i>
        </button>
      </div>

      {#if mobileMenuOpen}
        <div class="md:hidden mt-4 pb-4 border-t border-gray-100 pt-4 animate-fadeIn">
          <div class="flex flex-col gap-4">
            {#each navLinks as link}
              <a
                href={link.href}
                class="text-gray-700 hover:text-primary-600 transition-colors font-medium py-2 {$page.url.pathname === link.href ? 'text-primary-600' : ''}"
                on:click={closeMobileMenu}
              >
                {link.label}
              </a>
            {/each}
            <a href="/appointment" class="btn-primary text-center mt-2" on:click={closeMobileMenu}>
              Book Appointment
            </a>
          </div>
        </div>
      {/if}
    </nav>
  </header>

  <main class="min-h-screen">
    <slot />
  </main>

  <footer class="bg-gray-900 text-gray-300">
    <div class="container-custom py-12">
      <div class="grid md:grid-cols-4 gap-8">
        <div class="md:col-span-2">
          <div class="flex items-center gap-3 mb-4">
            <div class="w-10 h-10 bg-primary-500 rounded-lg flex items-center justify-center">
              <span class="text-white font-bold text-sm">F2H</span>
            </div>
            <div>
              <p class="font-bold text-white leading-tight">Friends2health</p>
              <p class="text-xs text-primary-300">Homoeo Clinic</p>
            </div>
          </div>
          <p class="text-gray-400 mb-4 max-w-md">
            Natural, gentle, and side-effect free homeopathic treatment by Dr. Aditi Singh.
            Experience holistic healing from the comfort of your home.
          </p>
        </div>

        <div>
          <h4 class="font-semibold text-white mb-4">Quick Links</h4>
          <ul class="space-y-2">
            <li><a href="/" class="hover:text-primary-300 transition-colors">Home</a></li>
            <li><a href="/about" class="hover:text-primary-300 transition-colors">About Dr. Aditi</a></li>
            <li><a href="/testimonials" class="hover:text-primary-300 transition-colors">Testimonials</a></li>
            <li><a href="/blog" class="hover:text-primary-300 transition-colors">Research and Blogs</a></li>
            <li><a href="/terms" class="hover:text-primary-300 transition-colors">Privacy & Terms</a></li>
          </ul>
        </div>

        <div>
          <h4 class="font-semibold text-white mb-4">Contact Us</h4>
          <ul class="space-y-3">
            <li class="flex items-start gap-3">
              <i class="fas fa-phone text-primary-400 mt-1"></i>
              <a href="tel:{contact.phoneRaw}" class="hover:text-primary-300 transition-colors">{contact.phone}</a>
            </li>
            <li class="flex items-start gap-3">
              <i class="fab fa-whatsapp text-primary-400 mt-1"></i>
              <a href="https://wa.me/{contact.whatsapp}" target="_blank" rel="noopener" class="hover:text-primary-300 transition-colors">WhatsApp Us</a>
            </li>
            <li class="flex items-start gap-3">
              <i class="fas fa-envelope text-primary-400 mt-1"></i>
              <a href="mailto:{contact.email}" class="hover:text-primary-300 transition-colors">{contact.email}</a>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <div class="border-t border-gray-800 py-6">
      <div class="container-custom flex flex-col md:flex-row justify-between items-center gap-4 text-sm text-gray-500">
        <p>&copy; {new Date().getFullYear()} Friends2health Homoeo Clinic. All rights reserved.</p>
        <p>Designed with <i class="fas fa-heart text-primary-400"></i> for healing</p>
      </div>
    </div>
  </footer>
{/if}
