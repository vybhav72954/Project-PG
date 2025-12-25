<script lang="ts">
  import { createEventDispatcher, onMount } from 'svelte'

  export let selectedTimeSlot: string | null = null
  export let selectedDate: Date | null = null

  const dispatch = createEventDispatcher()
  let availableSlots = []

  onMount(() => {
    // Ensure this component initializes properly
    if (selectedDate) {
      fetchTimeSlotsFromServer(selectedDate)
    }
  })

  // Watch for changes in selected date
  $: if (selectedDate) {
    fetchTimeSlotsFromServer(selectedDate)

    // Reset time slot if date changes
    if (selectedTimeSlot) {
      selectedTimeSlot = null
    }
  }

  function fetchTimeSlotsFromServer(date: Date) {
    if (!date) return

    console.log(
      `Fetching time slots for date: ${date.toISOString().split('T')[0]}`
    )

    // Hardcoded time slots based on the day of the week
    // In a real app, this would be an API call
    const dayOfWeek = date.getDay()

    const morningSlots = ['09:00 AM', '10:00 AM', '11:00 AM']
    const afternoonSlots = ['01:00 PM', '02:00 PM', '03:00 PM', '04:00 PM']
    const eveningSlots = ['05:00 PM', '06:00 PM', '07:00 PM']

    // Adjust slot availability based on day of week (just for demo variety)
    if (dayOfWeek === 1) {
      // Monday
      morningSlots.pop() // Remove last morning slot
    } else if (dayOfWeek === 5) {
      // Friday
      eveningSlots.shift() // Remove first evening slot
    }

    // Generate some random availability
    const randomAvailability = () => Math.random() > 0.3 // 70% available

    availableSlots = [
      {
        title: 'Morning',
        slots: morningSlots.map(time => ({
          time,
          available: randomAvailability(),
          booked: Math.random() > 0.7 // 30% booked
        }))
      },
      {
        title: 'Afternoon',
        slots: afternoonSlots.map(time => ({
          time,
          available: randomAvailability(),
          booked: Math.random() > 0.7
        }))
      },
      {
        title: 'Evening',
        slots: eveningSlots.map(time => ({
          time,
          available: randomAvailability(),
          booked: Math.random() > 0.7
        }))
      }
    ]

    console.log('Available time slots:', availableSlots)
  }

  function selectTimeSlot(time: string) {
    selectedTimeSlot = time
    dispatch('select', { time })
  }
</script>

<div class="time-slots">
  {#if !selectedDate}
    <p class="text-center text-gray-500">Please select a date first</p>
  {:else if availableSlots.length === 0}
    <p class="text-center text-gray-500">
      Loading available time slots...
    </p>
  {:else}
    <div class="space-y-4">
      {#each availableSlots as section}
        <div>
          <h4 class="mb-2 text-sm font-medium text-gray-700">
            {section.title}
          </h4>
          <div class="grid grid-cols-3 gap-2">
            {#each section.slots as slot}
              <button
                type="button"
                class="rounded-md border p-2 text-sm
                  {slot.available && !slot.booked
                  ? selectedTimeSlot === slot.time
                    ? 'border-[#d5c455] bg-[#d5c455] text-white'
                    : 'border-gray-300 hover:border-[#d5c455]'
                  : 'cursor-not-allowed border-gray-200 bg-gray-100 text-gray-400'}"
                disabled={!slot.available || slot.booked}
                on:click={() => selectTimeSlot(slot.time)}
              >
                {slot.time}
                {#if slot.booked}
                  <span class="block text-xs">(Booked)</span>
                {/if}
              </button>
            {/each}
          </div>
        </div>
      {/each}
    </div>
  {/if}

  {#if selectedTimeSlot}
    <div class="mt-4 text-center text-sm text-gray-600">
      Selected time: {selectedTimeSlot}
    </div>
  {/if}
</div>
