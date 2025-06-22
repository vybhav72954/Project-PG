<script>
  import { onMount } from 'svelte'
  import Carousel from '$lib/Carousel.svelte'

  /**
   * @type {string[]}
   */
  let images = [] // Default images

  // Function to dynamically load images from the carousel folder
  async function loadCarouselImages() {
    try {
      // This approach requires a server endpoint to fetch directory contents
      const response = await fetch('/api/carousel/images')
      if (response.ok) {
        const data = await response.json()
        if (Array.isArray(data) && data.length > 0) {
          images = data
        }
      }
    } catch (error) {
      console.error('Failed to load carousel images:', error)
      // Fallback to default images
    }
  }

  onMount(() => {
    loadCarouselImages()
  })
</script>

<div class="h-[70vh] w-full lg:h-full lg:min-h-screen">
  <Carousel {images} slideHold={6000} slideAnimation={400} />
</div>
