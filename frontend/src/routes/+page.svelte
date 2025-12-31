<script lang="ts">
  import { siteConfig } from '$lib/config';

  const { doctor, contact } = siteConfig;

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
      description: 'Choose your preferred consultation type and select a convenient time slot.'
    },
    {
      step: 2,
      icon: 'fa-comments',
      title: 'Consult with Doctor',
      description: 'Have a detailed consultation with Dr. Aditi via video or voice call.'
    },
    {
      step: 3,
      icon: 'fa-heart-pulse',
      title: 'Get Treatment',
      description: 'Receive personalized homeopathic treatment plan and follow-up care.'
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

  // Success stories - before/after cases
  const successStories = [
    {
      condition: 'Eczema Treatment',
      duration: '3 months',
      beforeImg: 'https://placehold.co/300x200/fee2e2/991b1b?text=Before',
      afterImg: 'https://placehold.co/300x200/dcfce7/166534?text=After',
      description: 'Chronic eczema cleared with constitutional treatment'
    },
    {
      condition: 'Hair Restoration',
      duration: '6 months',
      beforeImg: 'https://placehold.co/300x200/fee2e2/991b1b?text=Before',
      afterImg: 'https://placehold.co/300x200/dcfce7/166534?text=After',
      description: 'Significant hair regrowth with homeopathic remedies'
    },
    {
      condition: 'Psoriasis Recovery',
      duration: '4 months',
      beforeImg: 'https://placehold.co/300x200/fee2e2/991b1b?text=Before',
      afterImg: 'https://placehold.co/300x200/dcfce7/166534?text=After',
      description: 'Psoriasis patches healed naturally'
    },
    {
      condition: 'Acne Cure',
      duration: '2 months',
      beforeImg: 'https://placehold.co/300x200/fee2e2/991b1b?text=Before',
      afterImg: 'https://placehold.co/300x200/dcfce7/166534?text=After',
      description: 'Clear skin achieved without harsh chemicals'
    }
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
  <title>Friends2health Homoeo Clinic | Dr. Aditi Singh</title>
</svelte:head>

<!-- Hero Section - Meet Dr. Aditi (FIRST) -->
<section class="bg-gradient-hero min-h-[600px] flex items-center">
  <div class="container-custom w-full">
    <div class="grid lg:grid-cols-2 gap-12 items-center">
      <!-- Left: Content -->
      <div class="text-center lg:text-left">
        <p class="text-primary-600 font-medium tracking-wide mb-2">
          Welcome to Friends2health Homoeo Clinic
        </p>
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold text-gray-900 mb-4">
          Meet <span class="text-primary-500">{doctor.name}</span>
        </h1>
        <p class="text-lg text-gray-600 mb-2">
          {doctor.title} • {doctor.experience} Years Experience
        </p>
        <p class="text-gray-600 mb-6 max-w-lg">
          Dedicated homeopathic physician providing natural, gentle, and side-effect free treatment.
          Experience holistic healing from the comfort of your home.
        </p>

        <div class="space-y-2 mb-8">
          <div class="flex items-center gap-3 justify-center lg:justify-start">
            <i class="fas fa-graduation-cap text-primary-500 w-5"></i>
            <span class="text-gray-700 text-sm">BHMS - {doctor.bhmsCollege}</span>
          </div>
          <div class="flex items-center gap-3 justify-center lg:justify-start">
            <i class="fas fa-award text-primary-500 w-5"></i>
            <span class="text-gray-700 text-sm">MD - {doctor.mdCollege}</span>
          </div>
        </div>

        <div class="flex flex-col sm:flex-row gap-4 justify-center lg:justify-start">
          <a href="/appointment" class="btn-primary text-lg px-8 py-4">
            Book Consultation
          </a>
          <a href="/about" class="btn-secondary text-lg px-8 py-4">
            Learn More
          </a>
        </div>
      </div>

      <!-- Right: Doctor Image -->
      <div class="hidden lg:flex justify-center">
        <div class="relative">
          <div class="w-96 h-[480px] bg-gradient-green rounded-2xl overflow-hidden shadow-xl">
            <img
              src="/images/home-page_img.jpg"
              alt="Dr. Aditi Singh - Homeopathic Physician"
              class="w-full h-full object-cover"
            />
            <!-- Overlay with name -->
            <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6">
              <p class="text-white font-bold text-xl">{doctor.name}</p>
              <p class="text-white/80 text-sm">{doctor.title}</p>
            </div>
          </div>
          <!-- Decorative elements -->
          <div class="absolute -top-4 -right-4 w-20 h-20 bg-cta/30 rounded-full"></div>
          <div class="absolute -bottom-6 -left-6 w-28 h-28 bg-primary-200 rounded-full opacity-60"></div>
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
          <p class="text-gray-600">{item.description}</p>
        </div>
      {/each}
    </div>

    <div class="text-center mt-10">
      <a href="/appointment" class="btn-primary">
        <i class="fas fa-calendar-check mr-2"></i>
        Get Started Now
      </a>
    </div>
  </div>
</section>

<!-- Clinical Conditions -->
<section class="section-alt">
  <div class="container-custom">
    <div class="text-center mb-12">
      <h2 class="section-title">Conditions We Treat</h2>
      <p class="section-subtitle">
        We effectively treat more than <strong>50 common and uncommon</strong> clinical conditions using natural homeopathic remedies.
      </p>
    </div>

    <!-- Conditions Grid -->
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
      {#each conditions as condition}
        <div class="bg-white p-5 rounded-xl shadow-sm hover:shadow-md transition-shadow">
          <i class="fas {condition.icon} text-2xl text-primary-500 mb-3"></i>
          <h4 class="font-semibold text-gray-900 mb-1">{condition.name}</h4>
          <p class="text-sm text-gray-500">{condition.examples}</p>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- Why Homeopathy -->
<section class="section">
  <div class="container-custom">
    <div class="grid lg:grid-cols-2 gap-12 items-center">
      <!-- Left: Image Grid -->
      <div class="grid grid-cols-2 gap-4">
        <div class="space-y-4">
          <img
            src="https://images.unsplash.com/photo-1512069772995-ec65ed45afd6?w=300&h=250&fit=crop"
            alt="Natural remedies"
            class="rounded-xl shadow-md w-full h-48 object-cover"
          />
          <img
            src="https://images.unsplash.com/photo-1584308666744-24d5c474f2ae?w=300&h=200&fit=crop"
            alt="Homeopathic medicine"
            class="rounded-xl shadow-md w-full h-36 object-cover"
          />
        </div>
        <div class="space-y-4 pt-8">
          <img
            src="https://images.unsplash.com/photo-1587854692152-cbe660dbde88?w=300&h=200&fit=crop"
            alt="Natural healing"
            class="rounded-xl shadow-md w-full h-36 object-cover"
          />
          <img
            src="https://images.unsplash.com/photo-1471864190281-a93a3070b6de?w=300&h=250&fit=crop"
            alt="Holistic treatment"
            class="rounded-xl shadow-md w-full h-48 object-cover"
          />
        </div>
      </div>

      <!-- Right: Content -->
      <div>
        <h2 class="section-title">Why Choose Homeopathy?</h2>
        <p class="text-gray-600 mb-6">
          Homeopathy treats the person as a whole, addressing the root cause of illness rather than just suppressing symptoms.
        </p>

        <div class="space-y-4">
          <div class="flex items-start gap-4">
            <div class="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center flex-shrink-0">
              <i class="fas fa-leaf text-primary-600"></i>
            </div>
            <div>
              <h4 class="font-semibold text-gray-900">100% Natural</h4>
              <p class="text-sm text-gray-600">Made from natural substances with no harmful chemicals</p>
            </div>
          </div>

          <div class="flex items-start gap-4">
            <div class="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center flex-shrink-0">
              <i class="fas fa-shield-heart text-primary-600"></i>
            </div>
            <div>
              <h4 class="font-semibold text-gray-900">No Side Effects</h4>
              <p class="text-sm text-gray-600">Safe for all ages including children and elderly</p>
            </div>
          </div>

          <div class="flex items-start gap-4">
            <div class="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center flex-shrink-0">
              <i class="fas fa-bullseye text-primary-600"></i>
            </div>
            <div>
              <h4 class="font-semibold text-gray-900">Treats Root Cause</h4>
              <p class="text-sm text-gray-600">Addresses underlying issues, not just symptoms</p>
            </div>
          </div>

          <div class="flex items-start gap-4">
            <div class="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center flex-shrink-0">
              <i class="fas fa-hand-holding-heart text-primary-600"></i>
            </div>
            <div>
              <h4 class="font-semibold text-gray-900">Personalized Care</h4>
              <p class="text-sm text-gray-600">Treatment tailored to your unique constitution</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- Success Stories - Before/After -->
<section class="section-alt">
  <div class="container-custom">
    <div class="text-center mb-12">
      <h2 class="section-title">Success Stories Across the Horizon</h2>
      <p class="section-subtitle">
        Real results from real patients — witness the transformative power of homeopathy
      </p>
    </div>

    <div class="grid md:grid-cols-2 gap-8">
      {#each successStories as story}
        <div class="bg-white rounded-2xl shadow-sm overflow-hidden">
          <!-- Before/After Images -->
          <div class="grid grid-cols-2">
            <div class="relative">
              <img
                src={story.beforeImg}
                alt="Before treatment"
                class="w-full h-48 object-cover"
              />
              <span class="absolute top-3 left-3 bg-red-500 text-white text-xs font-bold px-3 py-1 rounded-full">
                BEFORE
              </span>
            </div>
            <div class="relative">
              <img
                src={story.afterImg}
                alt="After treatment"
                class="w-full h-48 object-cover"
              />
              <span class="absolute top-3 right-3 bg-green-500 text-white text-xs font-bold px-3 py-1 rounded-full">
                AFTER
              </span>
            </div>
          </div>

          <!-- Details -->
          <div class="p-5">
            <div class="flex items-center justify-between mb-2">
              <h3 class="font-bold text-gray-900">{story.condition}</h3>
              <span class="text-sm text-primary-600 font-medium">
                <i class="fas fa-clock mr-1"></i>{story.duration}
              </span>
            </div>
            <p class="text-gray-600 text-sm">{story.description}</p>
          </div>
        </div>
      {/each}
    </div>

    <div class="text-center mt-10">
      <p class="text-gray-500 text-sm">
        <i class="fas fa-shield-alt mr-1"></i>
        All images shared with patient consent
      </p>
    </div>
  </div>
</section>

<!-- Testimonials Preview -->
<section class="section">
  <div class="container-custom">
    <div class="text-center mb-12">
      <h2 class="section-title">What Our Patients Say</h2>
      <p class="section-subtitle">Real stories from real patients</p>
    </div>

    <div class="grid md:grid-cols-3 gap-8">
      {#each testimonials as testimonial}
        <div class="bg-white p-6 rounded-2xl shadow-sm">
          <div class="flex items-center gap-4 mb-4">
            <div class="w-12 h-12 bg-primary-100 rounded-full flex items-center justify-center">
              <i class="fas fa-user text-primary-500"></i>
            </div>
            <div>
              <p class="font-semibold text-gray-900">{testimonial.name}</p>
              <p class="text-sm text-primary-600">{testimonial.condition}</p>
            </div>
          </div>

          <p class="text-gray-600 mb-4">"{testimonial.text}"</p>

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
      <a href="tel:{contact.phoneRaw}" class="flex items-center gap-2 text-gray-600 hover:text-primary-600 transition-colors">
        <i class="fas fa-phone text-primary-500"></i>
        <span>{contact.phone}</span>
      </a>
      <a href="https://wa.me/{contact.whatsapp}" target="_blank" rel="noopener" class="flex items-center gap-2 text-gray-600 hover:text-primary-600 transition-colors">
        <i class="fab fa-whatsapp text-primary-500"></i>
        <span>WhatsApp Us</span>
      </a>
      <a href="mailto:{contact.email}" class="flex items-center gap-2 text-gray-600 hover:text-primary-600 transition-colors">
        <i class="fas fa-envelope text-primary-500"></i>
        <span>{contact.email}</span>
      </a>
    </div>
  </div>
</section>
