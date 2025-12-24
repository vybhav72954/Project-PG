import { superValidate } from 'sveltekit-superforms/server'
import { fail } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'
import { zod } from 'sveltekit-superforms/adapters'
import { appointmentSchema } from '@/AppointmentSchema'

export const load: PageServerLoad = async () => {
  const form = await superValidate(zod(appointmentSchema), {
    strict: false
  })

  return {
    form
  }
}

export const actions: Actions = {
  default: async ({ request }) => {
    const form = await superValidate(request, zod(appointmentSchema), {
      strict: true
    })

    if (!form.valid) {
      return fail(400, { form })
    }

    try {
      // Here you would normally process the form data
      // e.g., save to database, send confirmation emails, etc.
      // For demonstration, we'll just return the form data
      return { form }
    } catch (error) {
      console.error(error)

      return fail(500, {
        form,
        error: 'An error occurred while processing your appointment'
      })
    }
  }
}
