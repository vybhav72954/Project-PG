<script lang="ts">
  import { onMount } from 'svelte';
  import { api, type Testimonial } from '$lib/api';

  let testimonials: Testimonial[] = [];
  let loading = true;
  let error = '';

  onMount(async () => {
    try {
      const response = await api.getTestimonials();
      if (response.success && response.data) {
        testimonials = response.data;
      } else {
        error = response.error || 'Failed to load testimonials';
      }
    } catch (e) {
      error = 'Failed to load testimonials';
    } finally {
      loading = false;
    }
  });

  function getInitials(name: string): string {
    return name.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2);
  }
</script>

<svelte:head>
  <title>Patient Testimonials | Friends2health Homoeo Clinic</title>
</svelte:head>

<!-- Hero Section -->
<section class="bg-gradient-hero py-16">
  <div class="container-custom text-center">
    <span class="badge mb-4">Testimonials</span>
    <h1 class="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
      What Our Patients Say
    </h1>
    <p class="text-xl text-gray-600 max-w-2xl mx-auto">
      Real stories from real patients who have experienced the healing power of homeopathy
    </p>
  </div>
</section>

<!-- Testimonials Grid -->
<section class="section">
  <div class="container-custom">
    {#if loading}
      <div class="flex justify-center py-12">
        <div class="flex items-center gap-3 text-gray-500">
          <i class="fas fa-spinner fa-spin text-2xl text-primary-500"></i>
          <span>Loading testimonials...</span>
        </div>
      </div>
    {:else if error}
      <div class="text-center py-12">
        <i class="fas fa-exclamation-circle text-4xl text-red-400 mb-4"></i>
        <p class="text-gray-600">{error}</p>
      </div>
    {:else if testimonials.length === 0}
      <div class="text-center py-12">
        <i class="fas fa-comments text-4xl text-gray-300 mb-4"></i>
        <p class="text-gray-600">No testimonials yet. Be the first to share your experience!</p>
      </div>
    {:else}
      <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {#each testimonials as testimonial}
          <div class="card-hover p-6">
            <!-- Header -->
            <div class="flex items-center gap-4 mb-4">
              <div class="w-14 h-14 bg-primary-100 rounded-full flex items-center justify-center">
                <span class="text-primary-600 font-bold">{getInitials(testimonial.name)}</span>
              </div>
              <div>
                <p class="font-semibold text-gray-900">{testimonial.name}</p>
                {#if testimonial.condition}
                  <p class="text-sm text-primary-600">{testimonial.condition}</p>
                {/if}
              </div>
            </div>
            
            <!-- Rating -->
            <div class="flex gap-1 mb-4">
              {#each Array(5) as _, i}
                <i class="fas fa-star {i < testimonial.rating ? 'text-cta' : 'text-gray-200'}"></i>
              {/each}
            </div>
            
            <!-- Review -->
            <p class="text-gray-600 leading-relaxed">"{testimonial.review}"</p>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</section>

<!-- CTA Section -->
<section class="py-16 bg-primary-50">
  <div class="container-custom text-center">
    <h2 class="text-2xl md:text-3xl font-bold text-gray-900 mb-4">
      Ready to Experience Natural Healing?
    </h2>
    <p class="text-gray-600 mb-8 max-w-xl mx-auto">
      Join hundreds of satisfied patients who have found relief through homeopathy.
    </p>
    <a href="/appointment" class="btn-primary">
      <i class="fas fa-calendar-plus mr-2"></i>
      Book Your Consultation
    </a>
  </div>
</section>
