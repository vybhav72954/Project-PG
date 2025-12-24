<script lang="ts">
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  import { adminApi, isLoggedIn, clearToken } from '$lib/adminApi';

  let isAuthenticated = false;
  let loading = true;
  let sidebarOpen = false;

  const navItems = [
    { href: '/admin', label: 'Dashboard', icon: 'fa-chart-line' },
    { href: '/admin/appointments', label: 'Appointments', icon: 'fa-calendar-alt' },
    { href: '/admin/patients', label: 'Patients', icon: 'fa-users' },
    { href: '/admin/testimonials', label: 'Testimonials', icon: 'fa-star' },
    { href: '/admin/settings', label: 'Settings', icon: 'fa-cog' }
  ];

  onMount(() => {
    isAuthenticated = isLoggedIn();
    loading = false;
  });

  function handleLogout() {
    adminApi.logout();
    isAuthenticated = false;
    goto('/admin');
  }

  $: currentPath = $page.url.pathname;
</script>

<svelte:head>
  <title>Admin Panel - Friends2health Homoeo Clinic</title>
</svelte:head>

{#if loading}
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="animate-spin rounded-full h-12 w-12 border-4 border-primary-500 border-t-transparent"></div>
  </div>
{:else if !isAuthenticated}
  <slot />
{:else}
  <div class="min-h-screen bg-gray-100 flex">
    <!-- Sidebar -->
    <aside class="fixed inset-y-0 left-0 z-50 w-64 bg-gray-900 transform transition-transform duration-200 ease-in-out lg:translate-x-0 {sidebarOpen ? 'translate-x-0' : '-translate-x-full'}">
      <div class="flex flex-col h-full">
        <!-- Logo -->
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

        <!-- Navigation -->
        <nav class="flex-1 px-2 py-4 space-y-1 overflow-y-auto">
          {#each navItems as item}
            <a
              href={item.href}
              class="flex items-center px-4 py-3 text-sm rounded-lg transition-colors
                {currentPath === item.href 
                  ? 'bg-primary-500 text-white' 
                  : 'text-gray-300 hover:bg-gray-800 hover:text-white'}"
            >
              <i class="fas {item.icon} w-5"></i>
              <span class="ml-3">{item.label}</span>
            </a>
          {/each}
        </nav>

        <!-- Footer -->
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

    <!-- Overlay -->
    {#if sidebarOpen}
      <div 
        class="fixed inset-0 z-40 bg-black bg-opacity-50 lg:hidden"
        on:click={() => sidebarOpen = false}
        on:keydown={(e) => e.key === 'Escape' && (sidebarOpen = false)}
        role="button"
        tabindex="0"
      ></div>
    {/if}

    <!-- Main content -->
    <div class="flex-1 lg:ml-64">
      <!-- Top bar -->
      <header class="sticky top-0 z-30 bg-white shadow-sm">
        <div class="flex items-center justify-between h-16 px-4">
          <button on:click={() => sidebarOpen = true} class="lg:hidden text-gray-600 hover:text-gray-900">
            <i class="fas fa-bars text-xl"></i>
          </button>
          <div class="flex items-center space-x-4">
            <span class="text-sm text-gray-600">Welcome, Admin</span>
            <div class="w-8 h-8 bg-primary-500 rounded-full flex items-center justify-center">
              <i class="fas fa-user text-white text-sm"></i>
            </div>
          </div>
        </div>
      </header>

      <!-- Page content -->
      <main class="p-4 lg:p-8">
        <slot />
      </main>
    </div>
  </div>
{/if}
