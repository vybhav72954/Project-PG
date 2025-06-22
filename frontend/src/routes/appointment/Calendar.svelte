<script lang="ts">
  import { createEventDispatcher, onMount } from 'svelte'

  export let selectedDate: Date | null = null

  const dispatch = createEventDispatcher()
  let currentMonth = new Date()
  let days: Array<{
    date: Date | null
    isCurrentMonth: boolean
    isDisabled: boolean
    isSelected: boolean
    isToday: boolean
  }> = []

  // Initialize with current date
  onMount(() => {
    generateCalendar()
  })

  // Generate calendar days
  function generateCalendar() {
    days = []

    const firstDay = new Date(
      currentMonth.getFullYear(),
      currentMonth.getMonth(),
      1
    )
    const lastDay = new Date(
      currentMonth.getFullYear(),
      currentMonth.getMonth() + 1,
      0
    )

    // Get the day of the week for the first day (0 = Sunday, 6 = Saturday)
    const firstDayOfWeek = firstDay.getDay()

    // Fill in days from previous month
    for (let i = 0; i < firstDayOfWeek; i++) {
      const prevMonthDay = new Date(firstDay)
      prevMonthDay.setDate(prevMonthDay.getDate() - (firstDayOfWeek - i))
      days.push({
        date: prevMonthDay,
        isCurrentMonth: false,
        isDisabled: true,
        isSelected: false,
        isToday: isSameDay(prevMonthDay, new Date())
      })
    }

    // Fill in days for current month
    const today = new Date()
    for (let i = 1; i <= lastDay.getDate(); i++) {
      const date = new Date(
        currentMonth.getFullYear(),
        currentMonth.getMonth(),
        i
      )

      // Check against available dates from server (dummy implementation)
      // In a real app, this would check against dates returned from the API
      const isAvailable = Math.random() > 0.3 // 70% of dates are available
      const isDisabled = date < today || isWeekend(date) || !isAvailable

      days.push({
        date,
        isCurrentMonth: true,
        isDisabled,
        isSelected: selectedDate ? isSameDay(date, selectedDate) : false,
        isToday: isSameDay(date, today)
      })
    }

    // Fill in days from next month
    const daysNeeded = 42 - days.length // 6 rows × 7 days
    for (let i = 1; i <= daysNeeded; i++) {
      const nextMonthDay = new Date(lastDay)
      nextMonthDay.setDate(nextMonthDay.getDate() + i)
      days.push({
        date: nextMonthDay,
        isCurrentMonth: false,
        isDisabled: true,
        isSelected: false,
        isToday: isSameDay(nextMonthDay, new Date())
      })
    }
  }

  // Check if a date is a weekend
  function isWeekend(date: Date): boolean {
    const day = date.getDay()
    return day === 0 || day === 6 // 0 = Sunday, 6 = Saturday
  }

  // Check if two dates are the same day
  function isSameDay(date1: Date, date2: Date): boolean {
    return (
      date1.getDate() === date2.getDate() &&
      date1.getMonth() === date2.getMonth() &&
      date1.getFullYear() === date2.getFullYear()
    )
  }

  // Format month name
  function formatMonth(date: Date): string {
    return date.toLocaleString('default', {
      month: 'long',
      year: 'numeric'
    })
  }

  // Go to previous month
  function prevMonth() {
    currentMonth = new Date(
      currentMonth.getFullYear(),
      currentMonth.getMonth() - 1,
      1
    )
    generateCalendar()
  }

  // Go to next month
  function nextMonth() {
    currentMonth = new Date(
      currentMonth.getFullYear(),
      currentMonth.getMonth() + 1,
      1
    )
    generateCalendar()
  }

  // Handle date selection
  function selectDate(day: (typeof days)[0]) {
    if (day.isDisabled || !day.date) return

    selectedDate = day.date
    days = days.map(d => ({
      ...d,
      isSelected:
        d.date && selectedDate ? isSameDay(d.date, selectedDate) : false
    }))

    dispatch('select', { date: selectedDate })
  }

  // Watch for changes in selected date
  $: if (selectedDate) {
    // If selected date is in a different month, update the current month view
    if (
      selectedDate.getMonth() !== currentMonth.getMonth() ||
      selectedDate.getFullYear() !== currentMonth.getFullYear()
    ) {
      currentMonth = new Date(
        selectedDate.getFullYear(),
        selectedDate.getMonth(),
        1
      )
      generateCalendar()
    }
  }

  // Update calendar when month changes
  $: if (currentMonth) {
    generateCalendar()
  }
</script>

<div class="calendar">
  <!-- Calendar header -->
  <div class="mb-4 flex items-center justify-between">
    <button
      type="button"
      on:click={prevMonth}
      class="flex h-8 w-8 items-center justify-center rounded-full text-gray-600 hover:bg-gray-100"
      aria-label="Previous month"
    >
      <i class="fa-solid fa-chevron-left"></i>
    </button>
    <h3 class="text-lg font-medium">{formatMonth(currentMonth)}</h3>
    <button
      type="button"
      on:click={nextMonth}
      class="flex h-8 w-8 items-center justify-center rounded-full text-gray-600 hover:bg-gray-100"
      aria-label="Next month"
    >
      <i class="fa-solid fa-chevron-right"></i>
    </button>
  </div>

  <!-- Calendar grid -->
  <div class="grid grid-cols-7 gap-1">
    <!-- Day names -->
    {#each ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'] as dayName}
      <div class="p-2 text-center text-sm font-medium text-gray-500">
        {dayName}
      </div>
    {/each}

    <!-- Calendar days -->
    {#each days as day}
      <button
        type="button"
        class="flex aspect-square items-center justify-center rounded-full text-sm
        {day.isSelected ? 'bg-[#d5c455] text-white' : ''}
        {!day.isCurrentMonth
          ? 'text-gray-300'
          : day.isToday
            ? 'border border-[#d5c455] font-bold'
            : 'text-gray-700'}
        {day.isDisabled
          ? 'cursor-not-allowed opacity-40'
          : 'hover:bg-gray-100'}"
        disabled={day.isDisabled}
        on:click={() => selectDate(day)}
        aria-label={day.date
          ? day.date.toLocaleDateString('en-US', {
              day: 'numeric',
              month: 'long',
              year: 'numeric'
            })
          : ''}
        aria-pressed={day.isSelected}
        aria-disabled={day.isDisabled}
      >
        {day.date?.getDate()}
      </button>
    {/each}
  </div>

  <!-- Selected date display -->
  {#if selectedDate}
    <div class="mt-4 text-center text-sm text-gray-600">
      Selected: {selectedDate.toLocaleDateString('en-US', {
        weekday: 'long',
        year: 'numeric',
        month: 'long',
        day: 'numeric'
      })}
    </div>
  {/if}
</div>
