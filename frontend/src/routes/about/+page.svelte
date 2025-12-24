<script lang="ts">
  const qualifications = [
    {
      degree: 'BHMS',
      institution: 'Homoeopathic Medical College & Hospital, Chandigarh',
      description: 'Bachelor of Homeopathic Medicine and Surgery - comprehensive training in homeopathic principles, materia medica, and clinical practice.'
    },
    {
      degree: 'MD (Homoeopathy)',
      institution: 'National Institute of Homoeopathy, New Delhi',
      description: 'Advanced specialization in homeopathic medicine with focus on chronic diseases and constitutional treatment.'
    }
  ];

  const expertise = [
    'Chronic Skin Disorders',
    'Digestive Problems',
    'Respiratory Conditions',
    'Hair & Scalp Issues',
    'Mental Health',
    "Women's Health",
    'Joint & Muscle Pain',
    'Pediatric Care'
  ];

  const approach = [
    {
      title: 'Detailed Case Taking',
      description: 'Understanding your complete health history, lifestyle, and symptoms.',
      icon: 'fa-clipboard-list'
    },
    {
      title: 'Constitutional Analysis',
      description: 'Identifying your unique constitution for personalized treatment.',
      icon: 'fa-user-check'
    },
    {
      title: 'Natural Remedies',
      description: 'Prescribing gentle, side-effect free homeopathic medicines.',
      icon: 'fa-leaf'
    },
    {
      title: 'Follow-up Care',
      description: 'Regular monitoring and adjustment of treatment for best results.',
      icon: 'fa-calendar-check'
    }
  ];
</script>

<svelte:head>
  <title>About Dr. Aditi Singh | Friends2health Homoeo Clinic</title>
</svelte:head>

<!-- Hero Section -->
<section class="bg-gradient-hero py-16">
  <div class="container-custom">
    <div class="grid lg:grid-cols-2 gap-12 items-center">
      <!-- Image -->
      <div class="flex justify-center order-2 lg:order-1">
        <div class="relative">
          <div class="w-80 h-96 bg-gradient-green rounded-2xl flex items-center justify-center">
            <div class="text-center">
              <div class="w-40 h-40 bg-white rounded-full mx-auto mb-4 flex items-center justify-center shadow-lg">
                <i class="fas fa-user-doctor text-6xl text-primary-500"></i>
              </div>
              <p class="text-primary-700 font-bold text-xl">Dr. Aditi Singh</p>
              <p class="text-primary-600">BHMS, MD (Hom)</p>
            </div>
          </div>
          <!-- Decorative elements -->
          <div class="absolute -top-4 -left-4 w-20 h-20 bg-cta/20 rounded-full"></div>
          <div class="absolute -bottom-4 -right-4 w-16 h-16 bg-primary-300/40 rounded-full"></div>
        </div>
      </div>
      
      <!-- Content -->
      <div class="order-1 lg:order-2">
        <span class="badge mb-4">About the Doctor</span>
        <h1 class="text-4xl md:text-5xl font-bold text-gray-900 mb-4">
          Dr. Aditi Singh
        </h1>
        <p class="text-xl text-primary-600 mb-6">
          Homeopathic Physician | 3+ Years Experience
        </p>
        <p class="text-gray-600 mb-6">
          Dr. Aditi Singh is a passionate and dedicated homeopathic physician committed to providing 
          natural, gentle, and effective healthcare solutions. With her patient-centric approach, 
          she believes in treating the person as a whole, not just the disease.
        </p>
        <p class="text-gray-600 mb-8">
          Her journey in homeopathy began with a deep belief in the body's innate healing ability 
          and the power of natural remedies to restore health without side effects.
        </p>
        <a href="/appointment" class="btn-primary">
          <i class="fas fa-calendar-plus mr-2"></i>
          Book Consultation
        </a>
      </div>
    </div>
  </div>
</section>

<!-- Qualifications -->
<section class="section">
  <div class="container-custom">
    <div class="text-center mb-12">
      <h2 class="section-title">Education & Qualifications</h2>
      <p class="section-subtitle">Trained at India's premier homeopathic institutions</p>
    </div>
    
    <div class="grid md:grid-cols-2 gap-8 max-w-4xl mx-auto">
      {#each qualifications as qual}
        <div class="card-hover p-8">
          <div class="flex items-start gap-4">
            <div class="w-14 h-14 bg-primary-100 rounded-xl flex items-center justify-center flex-shrink-0">
              <i class="fas fa-graduation-cap text-2xl text-primary-600"></i>
            </div>
            <div>
              <h3 class="text-xl font-bold text-gray-900 mb-1">{qual.degree}</h3>
              <p class="text-primary-600 font-medium mb-2">{qual.institution}</p>
              <p class="text-gray-600 text-sm">{qual.description}</p>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- Treatment Approach -->
<section class="section-alt">
  <div class="container-custom">
    <div class="text-center mb-12">
      <h2 class="section-title">Treatment Approach</h2>
      <p class="section-subtitle">A systematic approach to holistic healing</p>
    </div>
    
    <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
      {#each approach as item, index}
        <div class="bg-white p-6 rounded-xl shadow-sm text-center">
          <div class="relative inline-block mb-4">
            <div class="w-16 h-16 bg-cta/20 rounded-full flex items-center justify-center">
              <i class="fas {item.icon} text-2xl text-cta-hover"></i>
            </div>
            <span class="absolute -top-1 -right-1 w-6 h-6 bg-primary-500 text-white text-xs font-bold rounded-full flex items-center justify-center">
              {index + 1}
            </span>
          </div>
          <h3 class="font-bold text-gray-900 mb-2">{item.title}</h3>
          <p class="text-gray-600 text-sm">{item.description}</p>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- Areas of Expertise -->
<section class="section">
  <div class="container-custom">
    <div class="grid lg:grid-cols-2 gap-12 items-center">
      <div>
        <h2 class="section-title">Areas of Expertise</h2>
        <p class="text-gray-600 mb-8">
          Dr. Aditi specializes in treating a wide range of acute and chronic conditions using 
          classical homeopathic principles. Her expertise spans across various health domains, 
          with a particular focus on conditions that have limited success with conventional treatment.
        </p>
        <div class="grid grid-cols-2 gap-4">
          {#each expertise as item}
            <div class="flex items-center gap-3">
              <i class="fas fa-check-circle text-primary-500"></i>
              <span class="text-gray-700">{item}</span>
            </div>
          {/each}
        </div>
      </div>
      
      <div class="bg-primary-50 rounded-2xl p-8">
        <h3 class="text-xl font-bold text-gray-900 mb-4">Why Choose Homeopathy?</h3>
        <ul class="space-y-4">
          <li class="flex items-start gap-3">
            <i class="fas fa-leaf text-primary-500 mt-1"></i>
            <div>
              <p class="font-medium text-gray-900">100% Natural</p>
              <p class="text-sm text-gray-600">Made from natural substances with no chemicals</p>
            </div>
          </li>
          <li class="flex items-start gap-3">
            <i class="fas fa-shield-heart text-primary-500 mt-1"></i>
            <div>
              <p class="font-medium text-gray-900">No Side Effects</p>
              <p class="text-sm text-gray-600">Safe for all ages including children and elderly</p>
            </div>
          </li>
          <li class="flex items-start gap-3">
            <i class="fas fa-bullseye text-primary-500 mt-1"></i>
            <div>
              <p class="font-medium text-gray-900">Treats Root Cause</p>
              <p class="text-sm text-gray-600">Addresses underlying issues, not just symptoms</p>
            </div>
          </li>
          <li class="flex items-start gap-3">
            <i class="fas fa-pills text-primary-500 mt-1"></i>
            <div>
              <p class="font-medium text-gray-900">Non-Addictive</p>
              <p class="text-sm text-gray-600">Can be safely discontinued without dependency</p>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</section>

<!-- CTA Section -->
<section class="py-16 bg-primary-500">
  <div class="container-custom text-center">
    <h2 class="text-3xl font-bold text-white mb-4">
      Start Your Healing Journey Today
    </h2>
    <p class="text-primary-100 mb-8 max-w-xl mx-auto">
      Book a consultation with Dr. Aditi Singh and experience the gentle power of homeopathy.
    </p>
    <div class="flex flex-col sm:flex-row gap-4 justify-center">
      <a href="/appointment" class="inline-flex items-center justify-center rounded-md bg-cta px-8 py-4 font-semibold text-gray-900 hover:bg-cta-hover transition-all">
        <i class="fas fa-calendar-plus mr-2"></i>
        Book Appointment
      </a>
      <a href="https://wa.me/919877505344" target="_blank" rel="noopener" class="inline-flex items-center justify-center rounded-md bg-white/10 border-2 border-white px-8 py-4 font-semibold text-white hover:bg-white/20 transition-all">
        <i class="fab fa-whatsapp mr-2"></i>
        WhatsApp Us
      </a>
    </div>
  </div>
</section>
