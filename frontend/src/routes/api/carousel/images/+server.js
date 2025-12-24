import { json } from '@sveltejs/kit'
import fs from 'fs'
import path from 'path'

/** @type {import('./$types').RequestHandler} */
export async function GET() {
  try {
    // Path to the carousel images directory relative to the static folder
    const carouselDir = path.join(process.cwd(), 'static/images/carousel')

    // Read the directory
    const files = fs.readdirSync(carouselDir)

    // Filter for image files (optional, can be expanded)
    const imageFiles = files.filter(file =>
      /\.(jpg|jpeg|png|gif|webp|svg)$/i.test(file)
    )

    // Create proper paths for the images
    const imagePaths = imageFiles.map(file => `/images/carousel/${file}`)

    return json(imagePaths)
  } catch (error) {
    console.error('Error reading carousel directory:', error)
    return json([])
  }
}
