<script lang="ts">
  import { onMount } from 'svelte';

  let currentSlide = 0;
  const heroSlides = [
    {
      tagline: 'Natural | Gentle | Side-Effect Free',
      title: 'HOMOEOPATHY',
      subtitle: 'Trusted by Millions'
    }
  ];

  const stats = [
    { number: '500+', label: 'Happy Patients' },
    { number: '3+', label: 'Years Experience' },
    { number: '50+', label: 'Conditions Treated' },
    { number: '98%', label: 'Success Rate' }
  ];

  const howItWorks = [
    {
      step: 1,
      icon: 'fa-user-plus',
      title: 'Book Appointment',
      description: 'Choose your preferred consultation type and select a convenient time slot.',
      link: '/appointment',
      linkText: 'Book Now'
    },
    {
      step: 2,
      icon: 'fa-comments',
      title: 'Consult with Doctor',
      description: 'Have a detailed consultation with Dr. Aditi via video or voice call.',
      link: '/appointment',
      linkText: 'Start Consultation'
    },
    {
      step: 3,
      icon: 'fa-heart-pulse',
      title: 'Get Treatment',
      description: 'Receive personalized homeopathic treatment plan and follow-up care.',
      link: '/about',
      linkText: 'Learn More'
    }
  ];

  const conditions = [
    { name: 'Skin Disorders', icon: 'fa-hand-dots', examples: 'Eczema, Psoriasis, Acne' },
    { name: 'Digestive Issues', icon: 'fa-virus', examples: 'IBS, Acidity, Constipation' },
    { name: 'Respiratory', icon: 'fa-lungs', examples: 'Asthma, Allergies, Sinusitis' },
    { name: 'Hair Problems', icon: 'fa-head-side', examples: 'Hair Fall, Dandruff, Alopecia' },
    { name: 'Mental Health', icon: 'fa-brain', examples: 'Anxiety, Depression, Stress' },
    { name: "Women's Health", icon: 'fa-venus', examples: 'PCOS, Menstrual Issues' },
    { name: 'Joint Pain', icon: 'fa-bone', examples: 'Arthritis, Back Pain, Gout' },
    { name: "Children's Care", icon: 'fa-baby', examples: 'Growth, Immunity, Behavior' }
  ];

  const testimonials = [
    {
      name: 'Priya M.',
      condition: 'Chronic Migraine',
      text: "After 5 years of suffering, Dr. Aditi's treatment finally gave me relief. Highly recommended!",
      rating: 5
    },
    {
      name: 'Rahul S.',
      condition: 'Eczema',
      text: 'The results were gradual but permanent. My skin has never been better.',
      rating: 5
    },
    {
      name: 'Sunita K.',
      condition: 'Respiratory Issues',
      text: "My daughter's recurring respiratory issues have significantly improved.",
      rating: 5
    }
  ];
</script>

<svelte:head>
  <title>Friends2health Homoeo Clinic | Natural Healing</title>
</svelte:head>

<!-- Hero Section -->
<section class="bg-gradient-hero min-h-[600px] flex items-center">
  <div class="container-custom w-full">
    <div class="grid lg:grid-cols-2 gap-12 items-center">
      <!-- Left: Content -->
      <div class="text-center lg:text-left">
        <p class="text-primary-500 font-medium tracking-wide mb-4">
          Natural <span class="text-gray-400">|</span> Gentle <span class="text-gray-400">|</span> Side-Effect Free
        </p>
        <h1 class="text-5xl md:text-6xl lg:text-7xl font-bold text-primary-500 mb-4">
          HOMOEOPATHY
        </h1>
        <p class="text-xl md:text-2xl text-gray-700 mb-8">
          Trusted by Millions
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center lg:justify-start">
          <a href="/appointment" class="btn-primary text-lg px-8 py-4">
            Book Consultation
          </a>
          <a href="/about" class="btn-secondary text-lg px-8 py-4">
            Know Your Doctor
          </a>
        </div>
      </div>
      
      <!-- Right: Image placeholder -->
      <div class="hidden lg:flex justify-center">
        <div class="relative">
          <!-- Decorative leaves/plants background -->
          <div class="w-96 h-96 bg-primary-100 rounded-full flex items-center justify-center">
            <div class="text-center p-8">
              <i class="fas fa-leaf text-7xl text-primary-500 mb-4"></i>
              <p class="text-primary-600 font-medium">Natural Remedies</p>
              <p class="text-primary-500 text-sm">for Holistic Healing</p>
            </div>
          </div>
          <!-- Floating decorative elements -->
          <div class="absolute -top-4 -right-4 w-16 h-16 bg-primary-300 rounded-full opacity-60"></div>
          <div class="absolute -bottom-6 -left-6 w-24 h-24 bg-primary-200 rounded-full opacity-60"></div>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- Stats Bar -->
<section class="bg-primary-500 py-8">
  <div class="container-custom">
    <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
      {#each stats as stat}
        <div class="text-center">
          <div class="text-3xl md:text-4xl font-bold text-white mb-1">{stat.number}</div>
          <div class="text-primary-100 text-sm md:text-base">{stat.label}</div>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- How It Works -->
<section class="section">
  <div class="container-custom">
    <div class="text-center mb-12">
      <h2 class="section-title">How It Works</h2>
      <p class="section-subtitle">
        Get the best <strong>homeopathic treatment</strong> from the comfort of your home in just 3 simple steps!
      </p>
    </div>
    
    <div class="grid md:grid-cols-3 gap-8">
      {#each howItWorks as item}
        <div class="card-hover text-center p-8 border-2 border-primary-100 rounded-2xl">
          <div class="relative inline-block mb-6">
            <div class="w-16 h-16 bg-cta rounded-lg flex items-center justify-center">
              <i class="fas {item.icon} text-2xl text-gray-900"></i>
            </div>
            <span class="absolute -top-2 -right-2 w-6 h-6 bg-primary-500 text-white text-sm font-bold rounded flex items-center justify-center">
              {item.step}
            </span>
          </div>
          <h3 class="text-xl font-bold text-gray-900 mb-3">{item.title}</h3>
          <p class="text-gray-600 mb-6">{item.description}</p>
          <a href={item.link} class="link-arrow">{item.linkText}</a>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- Clinical Conditions -->
<section class="section-alt">
  <div class="container-custom">
    <div class="grid lg:grid-cols-2 gap-12 items-center">
      <!-- Left: Condition grid -->
      <div class="grid grid-cols-2 gap-4">
        {#each conditions.slice(0, 4) as condition}
          <div class="bg-white p-6 rounded-xl shadow-sm hover:shadow-md transition-shadow">
            <i class="fas {condition.icon} text-2xl text-primary-500 mb-3"></i>
            <h4 class="font-semibold text-gray-900 mb-1">{condition.name}</h4>
            <p class="text-sm text-gray-500">{condition.examples}</p>
          </div>
        {/each}
      </div>
      
      <!-- Right: Content -->
      <div>
        <h2 class="section-title">Clinical Conditions</h2>
        <p class="text-gray-600 mb-6">
          We effectively treat more than <strong>50 common and uncommon</strong> clinical conditions using natural homeopathic remedies.
        </p>
        <p class="text-gray-600 mb-8">
          From chronic skin disorders to respiratory issues, digestive problems to mental health concerns — our holistic approach addresses the root cause, not just the symptoms.
        </p>
        <a href="/appointment" class="btn-primary">
          <i class="fas fa-calendar-check mr-2"></i>
          Book Consultation
        </a>
      </div>
    </div>
    
    <!-- More conditions -->
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mt-8">
      {#each conditions.slice(4) as condition}
        <div class="bg-white p-5 rounded-xl shadow-sm hover:shadow-md transition-shadow">
          <i class="fas {condition.icon} text-xl text-primary-500 mb-2"></i>
          <h4 class="font-semibold text-gray-900 text-sm mb-1">{condition.name}</h4>
          <p class="text-xs text-gray-500">{condition.examples}</p>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- Meet the Doctor -->
<section class="section">
  <div class="container-custom">
    <div class="grid lg:grid-cols-2 gap-12 items-center">
      <!-- Left: Content -->
      <div>
        <h2 class="section-title">Meet Dr. Aditi Singh</h2>
        <p class="text-gray-600 mb-4">
          Dr. Aditi Singh is a dedicated homeopathic physician with <strong>3+ years of clinical experience</strong>. Her patient-centered approach focuses on understanding the complete health picture before prescribing treatment.
        </p>
        <div class="space-y-3 mb-6">
          <div class="flex items-start gap-3">
            <i class="fas fa-graduation-cap text-primary-500 mt-1"></i>
            <div>
              <p class="font-medium text-gray-900">BHMS</p>
              <p class="text-sm text-gray-600">Homoeopathic Medical College & Hospital, Chandigarh</p>
            </div>
          </div>
          <div class="flex items-start gap-3">
            <i class="fas fa-award text-primary-500 mt-1"></i>
            <div>
              <p class="font-medium text-gray-900">MD (Homoeopathy)</p>
              <p class="text-sm text-gray-600">National Institute of Homoeopathy, New Delhi</p>
            </div>
          </div>
        </div>
        <div class="flex gap-4">
          <a href="/about" class="btn-secondary">Learn More</a>
          <a href="/appointment" class="btn-primary">Book Appointment</a>
        </div>
      </div>
      
      <!-- Right: Image placeholder -->
      <div class="flex justify-center">
        <div class="relative">
          <div class="w-80 h-96 bg-gradient-green rounded-2xl flex items-center justify-center">
            <div class="text-center">
              <div class="w-32 h-32 bg-white rounded-full mx-auto mb-4 flex items-center justify-center shadow-lg">
                <i class="fas fa-user-doctor text-5xl text-primary-500"></i>
              </div>
              <p class="text-primary-700 font-semibold">Dr. Aditi Singh</p>
              <p class="text-primary-600 text-sm">BHMS, MD (Hom)</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- Testimonials Preview -->
<section class="section-alt">
  <div class="container-custom">
    <div class="text-center mb-12">
      <h2 class="section-title">What Our Patients Say</h2>
      <p class="section-subtitle">Real stories from real patients</p>
    </div>
    
    <div class="grid md:grid-cols-3 gap-8">
      {#each testimonials as testimonial}
        <div class="bg-white p-6 rounded-2xl shadow-sm">
          <!-- Avatar -->
          <div class="flex items-center gap-4 mb-4">
            <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center">
              <i class="fas fa-user text-gray-400"></i>
            </div>
            <div>
              <p class="font-semibold text-gray-900">{testimonial.name}</p>
              <p class="text-sm text-primary-600">{testimonial.condition}</p>
            </div>
          </div>
          
          <!-- Review -->
          <p class="text-gray-600 mb-4">"{testimonial.text}"</p>
          
          <!-- Rating -->
          <div class="flex gap-1">
            {#each Array(testimonial.rating) as _}
              <i class="fas fa-star text-cta"></i>
            {/each}
          </div>
        </div>
      {/each}
    </div>
    
    <div class="text-center mt-8">
      <a href="/testimonials" class="btn-secondary">
        View All Testimonials
      </a>
    </div>
  </div>
</section>

<!-- CTA Section -->
<section class="py-20 bg-primary-500">
  <div class="container-custom text-center">
    <h2 class="text-3xl md:text-4xl font-bold text-white mb-4">
      Ready to Start Your Healing Journey?
    </h2>
    <p class="text-primary-100 text-lg mb-8 max-w-2xl mx-auto">
      Book your consultation today and experience the gentle power of homeopathy.
    </p>
    <a href="/appointment" class="inline-flex items-center justify-center rounded-md bg-cta px-8 py-4 text-lg font-semibold text-gray-900 shadow-lg hover:bg-cta-hover transition-all">
      <i class="fas fa-calendar-plus mr-2"></i>
      Book Your Appointment
    </a>
  </div>
</section>

<!-- Contact Info Bar -->
<section class="bg-gray-50 py-8 border-t border-gray-100">
  <div class="container-custom">
    <div class="flex flex-col md:flex-row justify-center items-center gap-8 text-center">
      <a href="tel:+919877505344" class="flex items-center gap-2 text-gray-600 hover:text-primary-600 transition-colors">
        <i class="fas fa-phone text-primary-500"></i>
        <span>+91-9877505344</span>
      </a>
      <a href="https://wa.me/919877505344" target="_blank" rel="noopener" class="flex items-center gap-2 text-gray-600 hover:text-primary-600 transition-colors">
        <i class="fab fa-whatsapp text-primary-500"></i>
        <span>WhatsApp Us</span>
      </a>
      <a href="mailto:contact@friends2health.com" class="flex items-center gap-2 text-gray-600 hover:text-primary-600 transition-colors">
        <i class="fas fa-envelope text-primary-500"></i>
        <span>contact@friends2health.com</span>
      </a>
    </div>
  </div>
</section>
