<script lang="ts">
  import { siteConfig } from '$lib/config';

  const { clinicName, doctor, contact } = siteConfig;

  // Last updated date
  const lastUpdated = 'December 25, 2024';

  // For navigation within the page
  let activeSection = 'privacy';

  const sections = [
    { id: 'privacy', label: 'Privacy Policy', icon: 'fa-shield-halved' },
    { id: 'terms', label: 'Terms of Service', icon: 'fa-file-contract' },
    { id: 'refund', label: 'Refund Policy', icon: 'fa-rotate-left' },
    { id: 'disclaimer', label: 'Medical Disclaimer', icon: 'fa-stethoscope' }
  ];
</script>

<svelte:head>
  <title>Privacy Policy & Terms | {clinicName}</title>
  <meta name="description" content="Privacy Policy, Terms of Service, and Refund Policy for {clinicName}. Learn how we protect your data and our service terms." />
</svelte:head>

<!-- Hero Section -->
<section class="bg-gradient-hero py-12">
  <div class="container-custom">
    <h1 class="text-3xl md:text-4xl font-bold text-gray-900 mb-2">
      Legal & Policies
    </h1>
    <p class="text-gray-600">
      Last updated: {lastUpdated}
    </p>
  </div>
</section>

<!-- Navigation Tabs -->
<section class="bg-white border-b border-gray-200 sticky top-[72px] z-40">
  <div class="container-custom">
    <div class="flex gap-1 py-2 overflow-x-auto">
      {#each sections as section}
        <a
          href="#{section.id}"
          on:click={() => activeSection = section.id}
          class="flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium transition-colors whitespace-nowrap {activeSection === section.id ? 'bg-primary-500 text-white' : 'text-gray-600 hover:bg-gray-100'}"
        >
          <i class="fas {section.icon}"></i>
          {section.label}
        </a>
      {/each}
    </div>
  </div>
</section>

<!-- Content -->
<section class="section">
  <div class="container-custom">
    <div class="max-w-4xl mx-auto">

      <!-- Privacy Policy -->
      <article id="privacy" class="mb-16 scroll-mt-32">
        <div class="flex items-center gap-3 mb-6">
          <div class="w-12 h-12 bg-primary-100 rounded-xl flex items-center justify-center">
            <i class="fas fa-shield-halved text-xl text-primary-600"></i>
          </div>
          <h2 class="text-2xl font-bold text-gray-900">Privacy Policy</h2>
        </div>

        <div class="prose prose-gray max-w-none space-y-6">
          <p class="text-gray-600">
            {clinicName} ("we", "us", or "our") is committed to protecting your privacy. This Privacy Policy explains how we collect, use, disclose, and safeguard your information when you use our website and services.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">1. Information We Collect</h3>
          <p class="text-gray-600">We collect information that you provide directly to us, including:</p>
          <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
            <li><strong>Personal Information:</strong> Name, email address, phone number</li>
            <li><strong>Health Information:</strong> Medical history, symptoms, and health concerns shared during consultations</li>
            <li><strong>Payment Information:</strong> Payment details processed securely through Razorpay (we do not store card details)</li>
            <li><strong>Communication Data:</strong> Records of your consultations and communications with us</li>
          </ul>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">2. How We Use Your Information</h3>
          <p class="text-gray-600">We use the collected information to:</p>
          <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
            <li>Provide homeopathic consultation services</li>
            <li>Schedule and manage appointments</li>
            <li>Process payments for consultations</li>
            <li>Send appointment reminders and follow-up communications</li>
            <li>Maintain medical records as required by law</li>
            <li>Improve our services and patient experience</li>
          </ul>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">3. Information Sharing</h3>
          <p class="text-gray-600">
            We do not sell, trade, or rent your personal information to third parties. We may share your information only in the following circumstances:
          </p>
          <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
            <li><strong>Payment Processing:</strong> With Razorpay for secure payment processing</li>
            <li><strong>Legal Requirements:</strong> When required by law or to protect our rights</li>
            <li><strong>With Your Consent:</strong> When you explicitly authorize us to share information</li>
          </ul>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">4. Data Security</h3>
          <p class="text-gray-600">
            We implement appropriate technical and organizational measures to protect your personal information against unauthorized access, alteration, disclosure, or destruction. However, no method of transmission over the Internet is 100% secure.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">5. Data Retention</h3>
          <p class="text-gray-600">
            We retain your medical records and personal information for a minimum of 7 years as required by healthcare regulations in India, or longer if required by law.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">6. Your Rights</h3>
          <p class="text-gray-600">You have the right to:</p>
          <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
            <li>Access your personal information</li>
            <li>Request correction of inaccurate data</li>
            <li>Request deletion of your data (subject to legal retention requirements)</li>
            <li>Opt-out of marketing communications</li>
          </ul>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">7. Contact Us</h3>
          <p class="text-gray-600">
            For privacy-related questions or requests, contact us at:
          </p>
          <div class="bg-gray-50 rounded-lg p-4 mt-2">
            <p class="text-gray-700"><strong>{clinicName}</strong></p>
            <p class="text-gray-600">Email: {contact.email}</p>
            <p class="text-gray-600">Phone: {contact.phone}</p>
          </div>
        </div>
      </article>

      <!-- Terms of Service -->
      <article id="terms" class="mb-16 scroll-mt-32">
        <div class="flex items-center gap-3 mb-6">
          <div class="w-12 h-12 bg-primary-100 rounded-xl flex items-center justify-center">
            <i class="fas fa-file-contract text-xl text-primary-600"></i>
          </div>
          <h2 class="text-2xl font-bold text-gray-900">Terms of Service</h2>
        </div>

        <div class="prose prose-gray max-w-none space-y-6">
          <p class="text-gray-600">
            By using the services of {clinicName}, you agree to these Terms of Service. Please read them carefully.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">1. Services Provided</h3>
          <p class="text-gray-600">
            {clinicName} provides online homeopathic consultation services by {doctor.name} ({doctor.title}). Our services include:
          </p>
          <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
            <li>Video and voice consultations</li>
            <li>Homeopathic treatment recommendations</li>
            <li>Follow-up consultations</li>
            <li>Prescription of homeopathic remedies</li>
          </ul>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">2. Appointment Booking</h3>
          <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
            <li>Appointments must be booked through our website</li>
            <li>Payment is required at the time of booking</li>
            <li>You will receive a confirmation email with consultation details</li>
            <li>Please join the consultation on time; late arrivals may result in shortened consultation time</li>
          </ul>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">3. Patient Responsibilities</h3>
          <p class="text-gray-600">As a patient, you agree to:</p>
          <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
            <li>Provide accurate and complete health information</li>
            <li>Inform the doctor of all current medications and treatments</li>
            <li>Follow the prescribed treatment plan</li>
            <li>Attend scheduled appointments or cancel/reschedule in advance</li>
            <li>Use the consultation services for legitimate healthcare purposes only</li>
          </ul>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">4. Consultation Fees</h3>
          <p class="text-gray-600">
            Consultation fees are displayed on our website and are subject to change. All fees are in Indian Rupees (INR) and are inclusive of applicable taxes.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">5. Intellectual Property</h3>
          <p class="text-gray-600">
            All content on this website, including text, graphics, logos, and images, is the property of {clinicName} and is protected by copyright laws.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">6. Limitation of Liability</h3>
          <p class="text-gray-600">
            {clinicName} and {doctor.name} shall not be liable for any indirect, incidental, or consequential damages arising from the use of our services. Our total liability shall not exceed the amount paid for the consultation.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">7. Governing Law</h3>
          <p class="text-gray-600">
            These Terms shall be governed by the laws of India. Any disputes shall be subject to the exclusive jurisdiction of the courts in the state where the clinic is registered.
          </p>
        </div>
      </article>

      <!-- Refund Policy -->
      <article id="refund" class="mb-16 scroll-mt-32">
        <div class="flex items-center gap-3 mb-6">
          <div class="w-12 h-12 bg-primary-100 rounded-xl flex items-center justify-center">
            <i class="fas fa-rotate-left text-xl text-primary-600"></i>
          </div>
          <h2 class="text-2xl font-bold text-gray-900">Refund & Cancellation Policy</h2>
        </div>

        <div class="prose prose-gray max-w-none space-y-6">
          <h3 class="text-lg font-semibold text-gray-900">Cancellation by Patient</h3>

          <div class="bg-green-50 border border-green-200 rounded-lg p-4">
            <p class="text-green-800 font-medium mb-2">
              <i class="fas fa-check-circle mr-2"></i>Full Refund
            </p>
            <p class="text-green-700 text-sm">
              Cancellation made <strong>24 hours or more</strong> before the scheduled appointment.
            </p>
          </div>

          <div class="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
            <p class="text-yellow-800 font-medium mb-2">
              <i class="fas fa-exclamation-circle mr-2"></i>50% Refund
            </p>
            <p class="text-yellow-700 text-sm">
              Cancellation made <strong>between 12-24 hours</strong> before the scheduled appointment.
            </p>
          </div>

          <div class="bg-red-50 border border-red-200 rounded-lg p-4">
            <p class="text-red-800 font-medium mb-2">
              <i class="fas fa-times-circle mr-2"></i>No Refund
            </p>
            <p class="text-red-700 text-sm">
              Cancellation made <strong>less than 12 hours</strong> before the appointment, or no-show without prior notice.
            </p>
          </div>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">Cancellation by Clinic</h3>
          <p class="text-gray-600">
            In the rare event that we need to cancel your appointment due to unforeseen circumstances, you will receive:
          </p>
          <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
            <li>Full refund, OR</li>
            <li>Free rescheduling to another available slot of your choice</li>
          </ul>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">Rescheduling</h3>
          <p class="text-gray-600">
            You may reschedule your appointment free of charge if done at least 12 hours before the scheduled time. Contact us via WhatsApp or email to reschedule.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">Refund Processing</h3>
          <p class="text-gray-600">
            Approved refunds will be processed within <strong>5-7 business days</strong> to the original payment method. Bank processing times may vary.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">How to Request a Refund</h3>
          <p class="text-gray-600">To request a cancellation or refund:</p>
          <div class="bg-gray-50 rounded-lg p-4 mt-2">
            <p class="text-gray-600 mb-2">
              <i class="fab fa-whatsapp text-green-500 mr-2"></i>
              WhatsApp: <a href="https://wa.me/{contact.whatsapp}" class="text-primary-600 hover:underline">{contact.phone}</a>
            </p>
            <p class="text-gray-600">
              <i class="fas fa-envelope text-primary-500 mr-2"></i>
              Email: <a href="mailto:{contact.email}" class="text-primary-600 hover:underline">{contact.email}</a>
            </p>
          </div>
        </div>
      </article>

      <!-- Medical Disclaimer -->
      <article id="disclaimer" class="mb-16 scroll-mt-32">
        <div class="flex items-center gap-3 mb-6">
          <div class="w-12 h-12 bg-primary-100 rounded-xl flex items-center justify-center">
            <i class="fas fa-stethoscope text-xl text-primary-600"></i>
          </div>
          <h2 class="text-2xl font-bold text-gray-900">Medical Disclaimer</h2>
        </div>

        <div class="prose prose-gray max-w-none space-y-6">
          <div class="bg-amber-50 border border-amber-200 rounded-lg p-6">
            <p class="text-amber-800 font-medium mb-3">
              <i class="fas fa-exclamation-triangle mr-2"></i>Important Notice
            </p>
            <p class="text-amber-700">
              The information provided on this website and during consultations is for general informational and educational purposes only. It is not intended to be a substitute for professional medical advice, diagnosis, or treatment.
            </p>
          </div>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">Nature of Homeopathic Treatment</h3>
          <p class="text-gray-600">
            Homeopathy is a complementary system of medicine. {doctor.name} is a qualified homeopathic physician registered with the appropriate medical council. The treatments provided are based on homeopathic principles and practices.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">Emergency Situations</h3>
          <p class="text-gray-600">
            <strong>Our online consultation service is NOT suitable for medical emergencies.</strong> If you are experiencing a medical emergency, please:
          </p>
          <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
            <li>Call emergency services immediately (112 in India)</li>
            <li>Visit the nearest hospital emergency department</li>
            <li>Contact your local emergency healthcare provider</li>
          </ul>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">Existing Treatments</h3>
          <p class="text-gray-600">
            Do not discontinue any existing medication or treatment without consulting your primary healthcare provider. Homeopathic treatment can be used alongside conventional medicine, but your doctor should be informed of all treatments you are receiving.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">Individual Results</h3>
          <p class="text-gray-600">
            Results from homeopathic treatment vary from person to person. The success stories and testimonials on our website represent individual experiences and are not guarantees of specific outcomes.
          </p>

          <h3 class="text-lg font-semibold text-gray-900 mt-8">Practitioner Credentials</h3>
          <div class="bg-gray-50 rounded-lg p-4">
            <p class="text-gray-700 font-medium">{doctor.name}</p>
            <p class="text-gray-600">{doctor.title}</p>
            <p class="text-gray-600 text-sm mt-2">BHMS - {doctor.bhmsCollege}</p>
            <p class="text-gray-600 text-sm">MD - {doctor.mdCollege}</p>
          </div>
        </div>
      </article>

    </div>
  </div>
</section>

<!-- Contact Section -->
<section class="bg-gray-50 py-12 border-t border-gray-200">
  <div class="container-custom text-center">
    <h3 class="text-xl font-bold text-gray-900 mb-4">Questions About Our Policies?</h3>
    <p class="text-gray-600 mb-6">
      If you have any questions or concerns, please don't hesitate to contact us.
    </p>
    <div class="flex flex-col sm:flex-row gap-4 justify-center">
      <a href="mailto:{contact.email}" class="btn-secondary">
        <i class="fas fa-envelope mr-2"></i>
        Email Us
      </a>
      <a href="https://wa.me/{contact.whatsapp}" target="_blank" rel="noopener" class="btn-primary">
        <i class="fab fa-whatsapp mr-2"></i>
        WhatsApp Us
      </a>
    </div>
  </div>
</section>
