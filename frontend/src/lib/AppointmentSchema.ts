import { z } from 'zod'

export const appointmentSchema = z.object({
  name: z.string().min(4, 'Name must be at least 4 characters'),
  contact: z.union(
    [
      z.string().email({ message: 'Invalid Email Address' }),
      z
        .string()
        .min(10, 'Mobile number must be 10 digits')
        .max(10, 'Mobile number must be 10 digits')
        .regex(/^\d{10}$/, { message: 'Must be a 10-Digit Mobile Number' })
    ],
    {
      errorMap: () => ({
        message: 'Please enter a Valid Email or 10-Digit Mobile Number'
      })
    }
  ),
  otp: z.string().length(6, 'OTP must be 6 digits').optional(),
  isVerified: z.boolean().default(false),
  consultationType: z.enum(['video', 'voice']).default('voice'),
  appointmentDate: z.string(),
  appointmentTime: z.string()
})
