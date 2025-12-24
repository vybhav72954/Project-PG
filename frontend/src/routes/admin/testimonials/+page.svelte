<script lang="ts">
  import { onMount } from 'svelte';
  import { adminApi, type Testimonial } from '$lib/adminApi';

  let testimonials: Testimonial[] = [];
  let loading = true;
  let error = '';

  // Form state
  let showForm = false;
  let editingId: string | null = null;
  let formData = {
    name: '',
    review: '',
    rating: 5,
    condition: '',
    is_active: true
  };
  let formLoading = false;
  let formError = '';

  onMount(async () => {
    await loadTestimonials();
  });

  async function loadTestimonials() {
    loading = true;
    error = '';

    const response = await adminApi.getTestimonials();

    if (response.success && response.data) {
      testimonials = response.data;
    } else {
      error = response.error || 'Failed to load testimonials';
    }

    loading = false;
  }

  function openAddForm() {
    editingId = null;
    formData = { name: '', review: '', rating: 5, condition: '', is_active: true };
    formError = '';
    showForm = true;
  }

  function openEditForm(testimonial: Testimonial) {
    editingId = testimonial.id;
    formData = {
      name: testimonial.name,
      review: testimonial.review,
      rating: testimonial.rating,
      condition: testimonial.condition,
      is_active: testimonial.is_active
    };
    formError = '';
    showForm = true;
  }

  function closeForm() {
    showForm = false;
    editingId = null;
    formError = '';
  }

  async function handleSubmit() {
    if (!formData.name.trim() || !formData.review.trim()) {
      formError = 'Name and review are required';
      return;
    }

    formLoading = true;
    formError = '';

    let response;
    if (editingId) {
      response = await adminApi.updateTestimonial(editingId, formData);
    } else {
      response = await adminApi.createTestimonial(formData);
    }

    if (response.success) {
      closeForm();
      await loadTestimonials();
    } else {
      formError = response.error || 'Failed to save testimonial';
    }

    formLoading = false;
  }

  async function toggleActive(id: string) {
    const response = await adminApi.toggleTestimonial(id);
    if (response.success) {
      await loadTestimonials();
    } else {
      alert(response.error || 'Failed to toggle testimonial');
    }
  }

  async function deleteTestimonial(id: string) {
    if (!confirm('Are you sure you want to delete this testimonial?')) return;

    const response = await adminApi.deleteTestimonial(id);
    if (response.success) {
      await loadTestimonials();
    } else {
      alert(response.error || 'Failed to delete testimonial');
    }
  }

  function renderStars(rating: number): string {
    return '★'.repeat(rating) + '☆'.repeat(5 - rating);
  }
</script>

<div>
  <div class="flex flex-col md:flex-row md:items-center md:justify-between mb-6">
    <div>
      <h1 class="text-2xl font-bold text-gray-900">Testimonials</h1>
      <p class="text-gray-600">Manage patient reviews displayed on the website</p>
    </div>
    <button on:click={openAddForm} class="btn-primary mt-4 md:mt-0">
      <i class="fas fa-plus mr-2"></i>Add Testimonial
    </button>
  </div>

  {#if loading}
    <div class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-10 w-10 border-4 border-primary-500 border-t-transparent"></div>
    </div>
  {:else if error}
    <div class="bg-red-50 text-red-600 px-6 py-4 rounded-xl">
      <i class="fas fa-exclamation-circle mr-2"></i>
      {error}
    </div>
  {:else if testimonials.length === 0}
    <div class="bg-white rounded-xl p-12 text-center shadow-sm border border-gray-100">
      <i class="fas fa-star text-4xl text-gray-300 mb-4"></i>
      <p class="text-gray-500">No testimonials yet</p>
      <button on:click={openAddForm} class="btn-primary mt-4">
        Add First Testimonial
      </button>
    </div>
  {:else}
    <div class="grid gap-4">
      {#each testimonials as testimonial}
        <div class="bg-white rounded-xl p-6 shadow-sm border border-gray-100 {!testimonial.is_active ? 'opacity-60' : ''}">
          <div class="flex justify-between items-start">
            <div class="flex-1">
              <div class="flex items-center gap-3 mb-2">
                <span class="font-medium text-gray-900">{testimonial.name}</span>
                <span class="text-primary-500">{renderStars(testimonial.rating)}</span>
                {#if !testimonial.is_active}
                  <span class="px-2 py-0.5 text-xs bg-gray-100 text-gray-600 rounded-full">Hidden</span>
                {/if}
              </div>
              <p class="text-gray-600 mb-2">"{testimonial.review}"</p>
              {#if testimonial.condition}
                <p class="text-sm text-gray-500">Treated for: {testimonial.condition}</p>
              {/if}
            </div>
            <div class="flex items-center gap-2 ml-4">
              <button 
                on:click={() => toggleActive(testimonial.id)}
                class="p-2 rounded-lg hover:bg-gray-100 {testimonial.is_active ? 'text-green-600' : 'text-gray-400'}"
                title={testimonial.is_active ? 'Hide' : 'Show'}
              >
                <i class="fas {testimonial.is_active ? 'fa-eye' : 'fa-eye-slash'}"></i>
              </button>
              <button 
                on:click={() => openEditForm(testimonial)}
                class="p-2 rounded-lg hover:bg-gray-100 text-blue-600"
                title="Edit"
              >
                <i class="fas fa-edit"></i>
              </button>
              <button 
                on:click={() => deleteTestimonial(testimonial.id)}
                class="p-2 rounded-lg hover:bg-gray-100 text-red-600"
                title="Delete"
              >
                <i class="fas fa-trash"></i>
              </button>
            </div>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<!-- Add/Edit Modal -->
{#if showForm}
  <div class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen px-4 pt-4 pb-20">
      <div class="fixed inset-0 bg-black bg-opacity-50" on:click={closeForm} on:keydown={(e) => e.key === 'Escape' && closeForm()} role="button" tabindex="0"></div>
      
      <div class="relative bg-white rounded-2xl max-w-lg w-full shadow-xl">
        <!-- Header -->
        <div class="px-6 py-4 border-b flex items-center justify-between">
          <h2 class="text-lg font-semibold text-gray-900">
            {editingId ? 'Edit Testimonial' : 'Add Testimonial'}
          </h2>
          <button on:click={closeForm} class="text-gray-400 hover:text-gray-600">
            <i class="fas fa-times text-xl"></i>
          </button>
        </div>

        <!-- Form -->
        <form on:submit|preventDefault={handleSubmit} class="p-6 space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Patient Name *</label>
            <input type="text" bind:value={formData.name} class="input" placeholder="First name only" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Review *</label>
            <textarea bind:value={formData.review} class="input" rows="4" placeholder="Patient's review..."></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Rating</label>
            <div class="flex gap-2">
              {#each [1, 2, 3, 4, 5] as star}
                <button
                  type="button"
                  on:click={() => formData.rating = star}
                  class="text-2xl {formData.rating >= star ? 'text-primary-500' : 'text-gray-300'} hover:text-primary-400"
                >
                  ★
                </button>
              {/each}
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Condition Treated</label>
            <input type="text" bind:value={formData.condition} class="input" placeholder="e.g., Migraine, Eczema, etc." />
          </div>

          <div class="flex items-center">
            <input type="checkbox" id="is_active" bind:checked={formData.is_active} class="w-4 h-4 text-primary-600 border-gray-300 rounded focus:ring-primary-500" />
            <label for="is_active" class="ml-2 text-sm text-gray-700">Show on website</label>
          </div>

          {#if formError}
            <div class="bg-red-50 text-red-600 px-4 py-3 rounded-lg text-sm">
              {formError}
            </div>
          {/if}

          <div class="flex gap-3 pt-4">
            <button type="button" on:click={closeForm} class="btn-secondary flex-1">
              Cancel
            </button>
            <button type="submit" class="btn-primary flex-1" disabled={formLoading}>
              {#if formLoading}
                <i class="fas fa-spinner fa-spin mr-2"></i>
              {/if}
              {editingId ? 'Save Changes' : 'Add Testimonial'}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
{/if}
