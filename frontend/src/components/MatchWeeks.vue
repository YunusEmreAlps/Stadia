<template>
  <div class="card">
    <div class="card-header">
      <h2>Match Schedule</h2>
    </div>

    <div v-if="league && league.fixtures" class="weeks-container">
      <WeekSection
        v-for="(weekMatches, weekIndex) in league.fixtures"
        :key="weekIndex"
        :weekNumber="weekIndex + 1"
        :matches="weekMatches"
        :currentWeek="league.currentWeek"
        @updateMatch="handleUpdateMatch"
      />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useLeagueStore } from '@/stores/league'
import WeekSection from './WeekSection.vue'

const leagueStore = useLeagueStore()

const league = computed(() => leagueStore.league)

const handleUpdateMatch = async (matchId, homeScore, awayScore) => {
  if (homeScore < 0 || awayScore < 0) {
    alert('Scores cannot be negative')
    return
  }

  try {
    await leagueStore.updateMatchResult(matchId, homeScore, awayScore)
  } catch (error) {
    console.error('Failed to update match:', error)
  }
}
</script>

<style scoped>
.weeks-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
  margin-top: 20px;
}
</style>
