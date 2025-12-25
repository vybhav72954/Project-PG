<script lang="ts">
  import { siteConfig } from '$lib/config';

  const { doctor } = siteConfig;

  // Content categories
  type ContentType = 'all' | 'research' | 'article' | 'case-study';
  let activeFilter: ContentType = 'all';

  // Research papers and blog posts
  const publications = [
    {
      id: 1,
      type: 'research',
      title: 'Efficacy of Homeopathic Treatment in Chronic Eczema: A Clinical Study',
      abstract: 'This study examines the effectiveness of constitutional homeopathic treatment in patients with chronic eczema over a 6-month period. Results showed significant improvement in 78% of cases.',
      journal: 'Indian Journal of Homeopathic Research',
      year: 2024,
      authors: 'Singh A., et al.',
      link: '#', // Replace with actual link
      tags: ['Dermatology', 'Clinical Study', 'Eczema']
    },
    {
      id: 2,
      type: 'research',
      title: 'Homeopathic Approach to Polycystic Ovarian Syndrome: A Retrospective Analysis',
      abstract: 'A retrospective analysis of 50 PCOS cases treated with individualized homeopathic remedies. The study demonstrates promising outcomes in menstrual regulation and hormonal balance.',
      journal: 'National Journal of Homoeopathy',
      year: 2023,
      authors: 'Singh A., Kumar R.',
      link: '#', // Replace with actual link
      tags: ["Women's Health", 'PCOS', 'Hormonal']
    },
    {
      id: 3,
      type: 'article',
      title: 'Understanding Constitutional Treatment in Homeopathy',
      abstract: 'An introductory guide to constitutional prescribing - how homeopaths select remedies based on the complete picture of a patient rather than just their symptoms.',
      date: '2024-12-15',
      readTime: '5 min read',
      link: '#',
      tags: ['Education', 'Homeopathy Basics']
    },
    {
      id: 4,
      type: 'case-study',
      title: 'Case Study: Chronic Migraine Resolution with Natrum Muriaticum',
      abstract: 'A detailed case study of a 35-year-old female patient suffering from chronic migraines for 8 years. Complete resolution achieved within 4 months of homeopathic treatment.',
      date: '2024-11-20',
      readTime: '8 min read',
      link: '#',
      tags: ['Case Study', 'Migraine', 'Neurology']
    }
  ];

  // Filter publications
  $: filteredPublications = activeFilter === 'all'
    ? publications
    : publications.filter(p => p.type === activeFilter);

  function getTypeLabel(type: string): string {
    switch(type) {
      case 'research': return 'Research Paper';
      case 'article': return 'Article';
      case 'case-study': return 'Case Study';
      default: return type;
    }
  }

  function getTypeColor(type: string): string {
    switch(type) {
      case 'research': return 'bg-blue-100 text-blue-700';
      case 'article': return 'bg-green-100 text-green-700';
      case 'case-study': return 'bg-purple-100 text-purple-700';
      default: return 'bg-gray-100 text-gray-700';
    }
  }
</script>

<svelte:head>
  <title>Research & Articles | {siteConfig.clinicName}</title>
  <meta name="description" content="Research papers, articles, and case studies by {doctor.name}. Explore insights into homeopathic treatment and clinical findings." />
</svelte:head>

<!-- Hero Section -->
<section class="bg-gradient-hero py-16">
  <div class="container-custom">
    <div class="max-w-3xl">
      <span class="badge mb-4">Research & Insights</span>
      <h1 class="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
        Publications & Articles
      </h1>
      <p class="text-xl text-gray-600">
        Explore research papers, clinical case studies, and educational articles by {doctor.name}.
      </p>
    </div>
  </div>
</section>

<!-- Filter Tabs -->
<section class="bg-white border-b border-gray-200 sticky top-[72px] z-40">
  <div class="container-custom">
    <div class="flex gap-2 py-4 overflow-x-auto">
      <button
        on:click={() => activeFilter = 'all'}
        class="px-4 py-2 rounded-full text-sm font-medium transition-colors whitespace-nowrap {activeFilter === 'all' ? 'bg-primary-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'}"
      >
        All
      </button>
      <button
        on:click={() => activeFilter = 'research'}
        class="px-4 py-2 rounded-full text-sm font-medium transition-colors whitespace-nowrap {activeFilter === 'research' ? 'bg-primary-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'}"
      >
        Research Papers
      </button>
      <button
        on:click={() => activeFilter = 'article'}
        class="px-4 py-2 rounded-full text-sm font-medium transition-colors whitespace-nowrap {activeFilter === 'article' ? 'bg-primary-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'}"
      >
        Articles
      </button>
      <button
        on:click={() => activeFilter = 'case-study'}
        class="px-4 py-2 rounded-full text-sm font-medium transition-colors whitespace-nowrap {activeFilter === 'case-study' ? 'bg-primary-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'}"
      >
        Case Studies
      </button>
    </div>
  </div>
</section>

<!-- Publications List -->
<section class="section">
  <div class="container-custom">
    {#if filteredPublications.length === 0}
      <div class="text-center py-12">
        <i class="fas fa-file-alt text-4xl text-gray-300 mb-4"></i>
        <p class="text-gray-500">No publications found in this category.</p>
      </div>
    {:else}
      <div class="space-y-6">
        {#each filteredPublications as pub}
          <article class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden hover:shadow-md transition-shadow">
            <div class="p-6 md:p-8">
              <!-- Header -->
              <div class="flex flex-wrap items-center gap-3 mb-4">
                <span class="px-3 py-1 rounded-full text-xs font-semibold {getTypeColor(pub.type)}">
                  {getTypeLabel(pub.type)}
                </span>
                {#if pub.type === 'research'}
                  <span class="text-sm text-gray-500">
                    <i class="fas fa-book mr-1"></i>{pub.journal} • {pub.year}
                  </span>
                {:else}
                  <span class="text-sm text-gray-500">
                    <i class="fas fa-calendar mr-1"></i>{pub.date}
                    {#if pub.readTime}
                      <span class="mx-2">•</span>
                      <i class="fas fa-clock mr-1"></i>{pub.readTime}
                    {/if}
                  </span>
                {/if}
              </div>

              <!-- Title -->
              <h2 class="text-xl md:text-2xl font-bold text-gray-900 mb-3 hover:text-primary-600 transition-colors">
                <a href={pub.link}>{pub.title}</a>
              </h2>

              <!-- Authors (for research) -->
              {#if pub.type === 'research' && pub.authors}
                <p class="text-sm text-primary-600 mb-3">
                  <i class="fas fa-users mr-1"></i>{pub.authors}
                </p>
              {/if}

              <!-- Abstract -->
              <p class="text-gray-600 mb-4">{pub.abstract}</p>

              <!-- Tags -->
              <div class="flex flex-wrap gap-2 mb-4">
                {#each pub.tags as tag}
                  <span class="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded">
                    {tag}
                  </span>
                {/each}
              </div>

              <!-- Action -->
              <a
                href={pub.link}
                class="inline-flex items-center text-primary-600 font-medium hover:text-primary-700 transition-colors"
              >
                {pub.type === 'research' ? 'View Publication' : 'Read More'}
                <i class="fas fa-arrow-right ml-2 text-sm"></i>
              </a>
            </div>
          </article>
        {/each}
      </div>
    {/if}
  </div>
</section>

<!-- About the Author -->
<section class="section-alt">
  <div class="container-custom">
    <div class="bg-white rounded-2xl shadow-sm p-8 md:p-10">
      <div class="flex flex-col md:flex-row gap-8 items-center">
        <div class="w-32 h-32 bg-primary-100 rounded-full flex items-center justify-center flex-shrink-0">
          <i class="fas fa-user-doctor text-5xl text-primary-500"></i>
        </div>
        <div>
          <h3 class="text-2xl font-bold text-gray-900 mb-2">About the Author</h3>
          <p class="text-lg text-primary-600 mb-3">{doctor.name}, {doctor.title}</p>
          <p class="text-gray-600 mb-4">
            {doctor.name} is a homeopathic physician with {doctor.experience} years of clinical experience.
            She completed her BHMS from {doctor.bhmsCollege} and MD from {doctor.mdCollege}.
            Her research interests include chronic skin disorders, women's health, and constitutional prescribing.
          </p>
          <a href="/about" class="btn-secondary">
            Learn More About {doctor.name.split(' ')[0]}
          </a>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- CTA Section -->
<section class="py-16 bg-primary-500">
  <div class="container-custom text-center">
    <h2 class="text-3xl font-bold text-white mb-4">
      Have Questions About Homeopathy?
    </h2>
    <p class="text-primary-100 mb-8 max-w-xl mx-auto">
      Book a consultation to discuss your health concerns with {doctor.name}.
    </p>
    <a href="/appointment" class="inline-flex items-center justify-center rounded-md bg-cta px-8 py-4 font-semibold text-gray-900 hover:bg-cta-hover transition-all">
      <i class="fas fa-calendar-plus mr-2"></i>
      Book Consultation
    </a>
  </div>
</section>