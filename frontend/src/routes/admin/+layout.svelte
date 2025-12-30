<script lang="ts">
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { onMount, onDestroy } from 'svelte';
  import { adminApi, isLoggedIn, clearToken } from '$lib/adminApi';
  import { browser } from '$app/environment';

  let isAuthenticated = false;
  let loading = true;
  let sidebarOpen = false;
  let loginPassword = '';
  let loginError = '';

  const IDLE_TIMEOUT = 15 * 60 * 1000;
  let idleTimer: number | undefined;

  const navItems = [
    { href: '/admin', label: 'Dashboard', icon: 'fa-chart-line', exact: true },
    { href: '/admin/appointments', label: 'Appointments', icon: 'fa-calendar-alt', exact: false },
    { href: '/admin/patients', label: 'Patients', icon: 'fa-users', exact: false },
    { href: '/admin/testimonials', label: 'Testimonials', icon: 'fa-star', exact: false },
    { href: '/admin/settings', label: 'Settings', icon: 'fa-cog', exact: false }
  ];

  function resetIdleTimer() {
    if (browser && isAuthenticated) {
      if (idleTimer) clearTimeout(idleTimer);
      idleTimer = window.setTimeout(() => {
        handleLogout();
        alert('You have been logged out due to inactivity.');
      }, IDLE_TIMEOUT);
    }
  }

  function setupIdleDetection() {
    if (browser) {
      const events = ['mousedown', 'mousemove', 'keydown', 'scroll', 'touchstart', 'click'];
      events.forEach(event => {
        document.addEventListener(event, resetIdleTimer, { passive: true });
      });
      resetIdleTimer();
    }
  }

  function cleanupIdleDetection() {
    if (browser) {
      if (idleTimer) clearTimeout(idleTimer);
      const events = ['mousedown', 'mousemove', 'keydown', 'scroll', 'touchstart', 'click'];
      events.forEach(event => {
        document.removeEventListener(event, resetIdleTimer);
      });
    }
  }

  onMount(async () => {
    // Check if token exists in localStorage
    if (!isLoggedIn()) {
      loading = false;
      return;
    }

    // Validate token by making an API call to backend
    try {
      const response = await adminApi.getDashboard();
      if (response.success) {
        isAuthenticated = true;
        setupIdleDetection();
      } else {
        // Token invalid or expired, clear it
        clearToken();
        isAuthenticated = false;
      }
    } catch {
      clearToken();
      isAuthenticated = false;
    }

    loading = false;
  });

  onDestroy(() => {
    cleanupIdleDetection();
  });

  function handleLogout() {
    cleanupIdleDetection();
    adminApi.logout();
    isAuthenticated = false;
    goto('/admin');
  }

  async function handleLogin() {
    loginError = '';
    if (!loginPassword) {
      loginError = 'Please enter password';
      return;
    }
    const response = await adminApi.login(loginPassword);
    if (response.success) {
      isAuthenticated = true;
      loginPassword = '';
      setupIdleDetection();
    } else {
      loginError = response.error || 'Invalid password';
    }
  }

  function getNavClass(href: string, exact: boolean): string {
    const isActive = exact ? $page.url.pathname === href : $page.url.pathname.startsWith(href);
    const base = 'flex items-center px-4 py-3 text-sm rounded-lg transition-colors';
    return isActive
      ? base + ' bg-primary-500 text-white'
      : base + ' text-gray-300 hover:bg-gray-800 hover:text-white';
  }
</script>

<svelte:head>
  <title>Admin Panel - Friends2health Homoeo Clinic</title>
</svelte:head>

{#if loading}
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="animate-spin rounded-full h-12 w-12 border-4 border-primary-500 border-t-transparent"></div>
  </div>
{:else if !isAuthenticated}
  <div class="min-h-screen flex items-center justify-center bg-gray-100 p-4">
    <div class="max-w-md w-full">
      <div class="bg-white rounded-2xl shadow-lg p-8">
        <div class="text-center mb-8">
          <div class="w-16 h-16 bg-primary-500 rounded-xl flex items-center justify-center mx-auto mb-4">
            <span class="text-white font-bold text-xl">F2H</span>
          </div>
          <h1 class="text-2xl font-bold text-gray-900">Admin Login</h1>
          <p class="text-gray-600 mt-2">Friends2health Homoeo Clinic</p>
        </div>

        {#if loginError}
          <div class="mb-4 p-3 bg-red-50 border border-red-200 text-red-700 rounded-lg text-sm">
            {loginError}
          </div>
        {/if}

        <form on:submit|preventDefault={handleLogin}>
          <div class="mb-6">
            <label for="password" class="block text-sm font-medium text-gray-700 mb-2">
              Password
            </label>
            <input
              type="password"
              id="password"
              bind:value={loginPassword}
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 outline-none transition-colors"
              placeholder="Enter admin password"
              required
            />
          </div>
          <button
            type="submit"
            class="w-full bg-primary-500 text-white py-3 px-4 rounded-lg font-semibold hover:bg-primary-600 transition-colors"
          >
            Login
          </button>
        </form>

        <div class="mt-6 text-center">
          <a href="/" class="text-sm text-gray-500 hover:text-primary-600 transition-colors">← Back to website</a>
        </div>
      </div>
    </div>
  </div>
{:else}
  <div class="min-h-screen bg-gray-100 flex">
    <aside class="fixed inset-y-0 left-0 z-50 w-64 bg-gray-900 transform transition-transform duration-200 ease-in-out lg:translate-x-0 {sidebarOpen ? 'translate-x-0' : '-translate-x-full'}">
      <div class="flex flex-col h-full">
        <div class="flex items-center justify-between h-16 px-4 bg-gray-800">
          <a href="/admin" class="flex items-center space-x-2">
            <div class="w-8 h-8 bg-primary-500 rounded-lg flex items-center justify-center">
              <span class="text-white font-bold text-xs">F2H</span>
            </div>
            <span class="text-white font-semibold">Admin Panel</span>
          </a>
          <button on:click={() => sidebarOpen = false} class="lg:hidden text-gray-400 hover:text-white">
            <i class="fas fa-times"></i>
          </button>
        </div>

        <nav class="flex-1 px-2 py-4 space-y-1 overflow-y-auto">
          {#each navItems as item}
            <a href={item.href} class={getNavClass(item.href, item.exact)}>
              <i class="fas {item.icon} w-5"></i>
              <span class="ml-3">{item.label}</span>
            </a>
          {/each}
        </nav>

        <div class="p-4 border-t border-gray-800">
          <a href="/" class="flex items-center px-4 py-2 text-sm text-gray-400 hover:text-white rounded-lg hover:bg-gray-800">
            <i class="fas fa-external-link-alt w-5"></i>
            <span class="ml-3">View Website</span>
          </a>
          <button
            on:click={handleLogout}
            class="w-full flex items-center px-4 py-2 mt-2 text-sm text-red-400 hover:text-red-300 rounded-lg hover:bg-gray-800"
          >
            <i class="fas fa-sign-out-alt w-5"></i>
            <span class="ml-3">Logout</span>
          </button>
        </div>
      </div>
    </aside>

    {#if sidebarOpen}
      <button
        class="fixed inset-0 z-40 bg-black bg-opacity-50 lg:hidden"
        on:click={() => sidebarOpen = false}
        aria-label="Close sidebar"
      ></button>
    {/if}

    <div class="flex-1 lg:ml-64 min-h-screen">
      <header class="sticky top-0 z-30 bg-white shadow-sm">
        <div class="flex items-center justify-between h-16 px-4">
          <button on:click={() => sidebarOpen = true} class="lg:hidden text-gray-600 hover:text-gray-900">
            <i class="fas fa-bars text-xl"></i>
          </button>
          <div class="hidden lg:block"></div>
          <div class="flex items-center space-x-4">
            <span class="text-sm text-gray-600">Admin</span>
            <div class="w-8 h-8 bg-primary-500 rounded-full flex items-center justify-center">
              <i class="fas fa-user text-white text-sm"></i>
            </div>
          </div>
        </div>
      </header>

      <main class="p-4 lg:p-8">
        <slot />
      </main>
    </div>
  </div>
{/if}
