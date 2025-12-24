<script>
  import { onMount } from 'svelte'
  import { fly } from 'svelte/transition'

  /**
   * @type {string[]}
   */
  export let images = []
  export let slideHold = 6000
  export let slideAnimation = 400
  let currentImageIndex = 0

  // Custom transition: slide in from left
  /**
   * @param {HTMLImageElement} node
   */
  function slideFromLeft(node, { duration = slideAnimation } = {}) {
    return {
      duration,
      css: (/** @type {number} */ t) => `
        transform: translateX(${(1 - t) * -100}%);
        opacity: ${t};
      `
    }
  }

  // Automatically change the image after a delay
  onMount(() => {
    const interval = setInterval(() => {
      currentImageIndex = (currentImageIndex + 1) % images.length
    }, slideHold)
    return () => clearInterval(interval)
  })

  // Change image when indicator is clicked
  /**
   * @param {number} index
   */
  function changeImage(index) {
    currentImageIndex = index
  }
</script>

<div class="relative h-full w-full overflow-hidden bg-[#d5c455]">
  {#each images as image, index (index)}
    {#if index === currentImageIndex}
      <img
        src={image}
        alt="Carousel slide {index + 1}"
        class="absolute top-0 left-0 h-full w-full object-cover"
        in:slideFromLeft
        out:fly={{ x: 100, duration: slideAnimation }}
      />
    {/if}
  {/each}
  <!-- Indicators -->
  <div
    class="absolute bottom-6 left-1/2 z-10 flex -translate-x-1/2 transform"
  >
    {#each images as _, index (index)}
      <button
        type="button"
        class="mx-1 h-3 w-3 rounded-full {index === currentImageIndex
          ? 'bg-blue-500'
          : 'bg-opacity-75 bg-white'} hover:bg-blue-300"
        on:click={() => changeImage(index)}
        aria-label="Go to Slide {index + 1}"
      ></button>
    {/each}
  </div>
</div>
