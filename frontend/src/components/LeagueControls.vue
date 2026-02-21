<template>
  <div class="card controls-card">
    <div class="card-header">
      <div>
        <h2>League Controls</h2>
        <p class="week-info">Week {{ currentWeek }} of {{ totalWeeks }}</p>
      </div>
    </div>

    <div class="controls-grid">
      <button 
        class="btn-primary"
        @click="playNextWeek"
        :disabled="!canPlayNext || loading || isAutoPlaying"
      >
        ▶️ Play Next Week
      </button>

      <button 
        class="btn-secondary"
        @click="toggleAutoPlay"
        :disabled="!canPlayNext || loading"
      >
        {{ isAutoPlaying ? '⏸️ Stop Auto Play' : '⏩ Auto Play All' }}
      </button>

      <button 
        class="btn-danger"
        @click="resetLeague"
        :disabled="loading || isAutoPlaying"
      >
        🔄 Reset League
      </button>
    </div>

    <div v-if="loading" class="loading-bar">
      <div class="loading-progress"></div>
    </div>

    <div v-if="isFinished" class="finished-message">
      🎉 League Finished! Check the final standings above.
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useLeagueStore } from '@/stores/league'

const leagueStore = useLeagueStore()

const currentWeek = computed(() => leagueStore.currentWeek)
const totalWeeks = computed(() => leagueStore.totalWeeks)
const canPlayNext = computed(() => leagueStore.canPlayNext)
const isFinished = computed(() => leagueStore.isFinished)
const loading = computed(() => leagueStore.loading)
const isAutoPlaying = computed(() => !!leagueStore.autoPlayInterval)

const playNextWeek = async () => {
  await leagueStore.playNextWeek()
}

const toggleAutoPlay = () => {
  if (isAutoPlaying.value) {
    leagueStore.stopAutoPlay()
  } else {
    leagueStore.startAutoPlay(2000) // 2 seconds between weeks
  }
}

const resetLeague = async () => {
  if (confirm('Are you sure you want to reset the league? All progress will be lost.')) {
    await leagueStore.resetLeague()
  }
}
</script>

<style scoped>
.controls-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.card-header h2 {
  color: white;
}

.week-info {
  margin-top: 8px;
  font-size: 14px;
  opacity: 0.9;
}

.controls-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
  margin-top: 20px;
}

.controls-grid button {
  padding: 14px 20px;
  font-size: 15px;
}

.loading-bar {
  width: 100%;
  height: 4px;
  background-color: rgba(255, 255, 255, 0.3);
  border-radius: 2px;
  margin-top: 20px;
  overflow: hidden;
}

.loading-progress {
  height: 100%;
  background-color: white;
  animation: loading 1.5s ease-in-out infinite;
}

@keyframes loading {
  0% {
    width: 0;
    transform: translateX(0);
  }
  50% {
    width: 100%;
    transform: translateX(0);
  }
  100% {
    width: 100%;
    transform: translateX(100%);
  }
}

.finished-message {
  background-color: rgba(255, 255, 255, 0.2);
  padding: 16px;
  border-radius: 8px;
  text-align: center;
  font-size: 18px;
  font-weight: 600;
  margin-top: 20px;
}

@media (max-width: 768px) {
  .controls-grid {
    grid-template-columns: 1fr;
  }
}
</style>
