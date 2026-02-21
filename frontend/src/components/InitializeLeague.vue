<template>
  <div class="card">
    <div class="card-header">
      <h2>⚙️ Initialize League</h2>
    </div>

    <div class="team-selection">
      <p class="mb-2">Select teams to participate in the Champions League group stage:</p>
      
      <div class="teams-grid">
        <div 
          v-for="team in availableTeams" 
          :key="team.name"
          class="team-card"
          :class="{ selected: isSelected(team) }"
          @click="toggleTeam(team)"
        >
          <div class="team-logo">{{ team.logo }}</div>
          <div class="team-info">
            <h3>{{ team.name }}</h3>
            <div class="power-bar">
              <div class="power-fill" :style="{ width: team.power + '%' }"></div>
              <span class="power-label">Power: {{ team.power }}</span>
            </div>
          </div>
          <div class="check-icon" v-if="isSelected(team)">✓</div>
        </div>
      </div>

      <div class="selection-info">
        <p>Selected: {{ selectedTeams.length }} / 4 teams</p>
        <p v-if="selectedTeams.length < 4" class="text-secondary">
          Please select {{ 4 - selectedTeams.length }} more team(s)
        </p>
      </div>

      <button 
        class="btn-primary btn-large"
        @click="initializeLeague"
        :disabled="selectedTeams.length !== 4 || loading"
      >
        {{ loading ? 'Initializing...' : 'Start League' }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useLeagueStore } from '@/stores/league'

const emit = defineEmits(['initialized'])
const leagueStore = useLeagueStore()

const availableTeams = ref([
  { name: 'Manchester City', power: 92, logo: '🏴' },
  { name: 'Real Madrid', power: 91, logo: '🇪🇸' },
  { name: 'Bayern Munich', power: 90, logo: '🇩🇪' },
  { name: 'Paris Saint-Germain', power: 88, logo: '🇫🇷' },
  { name: 'Liverpool', power: 87, logo: '🏴' },
  { name: 'Barcelona', power: 86, logo: '🇪🇸' },
  { name: 'Inter Milan', power: 84, logo: '🇮🇹' },
  { name: 'Arsenal', power: 83, logo: '🏴' },
  { name: 'Napoli', power: 82, logo: '🇮🇹' },
  { name: 'Atletico Madrid', power: 81, logo: '🇪🇸' },
  { name: 'Chelsea', power: 80, logo: '🏴' },
  { name: 'Borussia Dortmund', power: 79, logo: '🇩🇪' },
  { name: 'Ajax', power: 78, logo: '🇳🇱' },
  { name: 'Porto', power: 77, logo: '🇵🇹' },
  { name: 'RB Leipzig', power: 76, logo: '🇩🇪' },
  { name: 'AC Milan', power: 75, logo: '🇮🇹' },
  { name: 'Beşiktaş', power: 74, logo: '🇹🇷' },
  { name: 'Celtic', power: 73, logo: '🏴' },
  { name: 'Qarabağ', power: 70, logo: '🇦🇿' },
  { name: 'Konyaspor', power: 68, logo: '🇹🇷' },
])

const selectedTeams = ref([])
const loading = ref(false)

const isSelected = (team) => {
  return selectedTeams.value.some(t => t.name === team.name)
}

const toggleTeam = (team) => {
  const index = selectedTeams.value.findIndex(t => t.name === team.name)
  
  if (index !== -1) {
    selectedTeams.value.splice(index, 1)
  } else if (selectedTeams.value.length < 4) {
    selectedTeams.value.push(team)
  }
}

const initializeLeague = async () => {
  if (selectedTeams.value.length !== 4) return

  loading.value = true
  try {
    await leagueStore.initializeLeague(selectedTeams.value)
    emit('initialized')
  } catch (error) {
    console.error('Failed to initialize league:', error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.team-selection {
  padding: 20px 0;
}

.teams-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
  margin: 24px 0;
}

.team-card {
  position: relative;
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px;
  border: 2px solid var(--border-color);
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  background-color: white;
}

.team-card:hover {
  border-color: var(--primary-color);
  transform: translateY(-4px);
  box-shadow: var(--shadow-md);
}

.team-card.selected {
  border-color: var(--primary-color);
  background-color: #e8f0fe;
}

.team-logo {
  font-size: 40px;
  flex-shrink: 0;
}

.team-info {
  flex: 1;
  min-width: 0;
}

.team-info h3 {
  font-size: 16px;
  margin-bottom: 8px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.power-bar {
  position: relative;
  height: 24px;
  background-color: #e0e0e0;
  border-radius: 12px;
  overflow: hidden;
}

.power-fill {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  background: linear-gradient(90deg, #4caf50, #8bc34a);
  transition: width 0.3s ease;
}

.power-label {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  font-size: 12px;
  font-weight: 600;
  color: #333;
  z-index: 1;
}

.check-icon {
  position: absolute;
  top: 8px;
  right: 8px;
  width: 28px;
  height: 28px;
  background-color: var(--primary-color);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
}

.selection-info {
  text-align: center;
  margin: 24px 0;
  font-size: 16px;
  font-weight: 600;
}

.text-secondary {
  color: var(--text-secondary);
  font-weight: normal;
  font-size: 14px;
  margin-top: 4px;
}

.btn-large {
  width: 100%;
  padding: 16px;
  font-size: 16px;
  margin-top: 16px;
}

@media (max-width: 768px) {
  .teams-grid {
    grid-template-columns: 1fr;
  }
}
</style>
