<script lang="ts">
  import { fly } from 'svelte/transition';
  import { toasts } from '$lib';
</script>

<div class="fixed top-4 right-4 z-50 space-y-2">
  {#each $toasts as toast (toast.id)}
    <div
      class="flex items-center space-x-3 px-4 py-3 rounded-lg shadow-lg max-w-sm
        {toast.type === 'success' ? 'bg-green-500' : ''}
        {toast.type === 'error' ? 'bg-red-500' : ''}
        {toast.type === 'info' ? 'bg-blue-500' : ''}
        text-white"
      in:fly={{ x: 100, duration: 300 }}
      out:fly={{ x: 100, duration: 200 }}
    >
      <i class="fas 
        {toast.type === 'success' ? 'fa-check-circle' : ''}
        {toast.type === 'error' ? 'fa-exclamation-circle' : ''}
        {toast.type === 'info' ? 'fa-info-circle' : ''}
      "></i>
      <span class="flex-1 text-sm">{toast.message}</span>
      <button 
        class="hover:opacity-75" 
        on:click={() => toasts.remove(toast.id)}
        aria-label="Dismiss"
      >
        <i class="fas fa-times"></i>
      </button>
    </div>
  {/each}
</div>
