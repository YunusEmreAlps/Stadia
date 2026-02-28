<template>
  <div id="app">
    <header class="app-header">
      <div class="container">
        <h1>Stadia - Champions League Simulation</h1>
        <p class="subtitle">Experience the thrill of the Champions League group stage</p>
      </div>
    </header>

    <main class="container">
      <InitializeLeague v-if="!league" @initialized="onLeagueInitialized" />
      
      <div v-else>
        <LeagueControls />
        <LeagueTable />
        <Predictions v-if="showPredictions" />
        <MatchWeeks />
      </div>
    </main>

    <div v-if="error" class="error-toast">
      {{ error }}
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useLeagueStore } from '@/stores/league'
import InitializeLeague from '@/components/InitializeLeague.vue'
import LeagueControls from '@/components/LeagueControls.vue'
import LeagueTable from '@/components/LeagueTable.vue'
import Predictions from '@/components/Predictions.vue'
import MatchWeeks from '@/components/MatchWeeks.vue'

const leagueStore = useLeagueStore()

const league = computed(() => leagueStore.league)
const error = computed(() => leagueStore.error)
const showPredictions = computed(() => leagueStore.showPredictions)

const onLeagueInitialized = () => {
  // League has been initialized
}

onMounted(() => {
  // Could load saved state from localStorage here
})
</script>

<style scoped>
.app-header {
  background: linear-gradient(135deg, #11998e 0%, #1a7a5e 50%, #38ef7d 100%);
  color: white;
  padding: 40px 0;
  margin-bottom: 30px;
  box-shadow: var(--shadow-md);
}

.app-header h1 {
  font-size: 36px;
  margin-bottom: 8px;
  font-weight: 800;
}

.subtitle {
  font-size: 16px;
  opacity: 0.9;
}

.error-toast {
  position: fixed;
  bottom: 20px;
  right: 20px;
  background-color: var(--danger-color);
  color: white;
  padding: 16px 24px;
  border-radius: 8px;
  box-shadow: var(--shadow-lg);
  animation: slideIn 0.3s ease;
  z-index: 1000;
}

@keyframes slideIn {
  from {
    transform: translateX(400px);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

@media (max-width: 768px) {
  .app-header h1 {
    font-size: 24px;
  }

  .subtitle {
    font-size: 14px;
  }
}
</style>
