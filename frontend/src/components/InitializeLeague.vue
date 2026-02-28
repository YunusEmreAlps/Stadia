<template>
  <div class="card">
    <div class="card-header">
      <h2>Initialize League</h2>
    </div>

    <div class="team-selection">
      <p class="mb-2">Select teams to participate in the Champions League group stage <b>(minimum {{ MIN_TEAMS }} teams)</b>:</p>
      
      <div class="teams-grid">
        <div 
          v-for="team in availableTeams" 
          :key="team.name"
          class="team-card"
          :class="{ selected: isSelected(team) }"
          @click="toggleTeam(team)"
        >
          <div class="team-logo">
            <img 
              v-if="team.logo.length > 0" 
              :src="`${FLAG_BASE_URL}/${team.logo.toLowerCase()}.svg`"
              :alt="team.name + ' logo'"
              width="40"
              height="40"
              class="team-logo-img"
            />
            <span v-else>{{ team.logo }}</span>
          </div>
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
        <p>Selected: {{ selectedTeams.length }} teams</p>
        <p v-if="selectedTeams.length < MIN_TEAMS" class="text-warning">
          ⚠️ Please select at least {{ MIN_TEAMS - selectedTeams.length }} more team(s) (minimum {{ MIN_TEAMS }} required)
        </p>
        <p v-else class="text-success">
          ✓ Ready to start! {{ selectedTeams.length }} teams selected
        </p>
      </div>

      <button 
        class="btn-primary btn-large"
        @click="initializeLeague"
        :disabled="selectedTeams.length < MIN_TEAMS || loading"
      >
        {{ loading ? 'Initializing...' : `Start League with ${selectedTeams.length} Teams` }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useLeagueStore } from '@/stores/league'
import { AVAILABLE_TEAMS, MIN_TEAMS, FLAG_BASE_URL } from '@/constants/teams'

const emit = defineEmits(['initialized'])
const leagueStore = useLeagueStore()

const availableTeams = ref(AVAILABLE_TEAMS)
const selectedTeams = ref([])
const loading = ref(false)

const isSelected = (team) => {
  return selectedTeams.value.some(t => t.name === team.name)
}

const toggleTeam = (team) => {
  const index = selectedTeams.value.findIndex(t => t.name === team.name)
  
  if (index !== -1) {
    selectedTeams.value.splice(index, 1)
  } else {
    selectedTeams.value.push(team)
  }
}

const initializeLeague = async () => {
  if (selectedTeams.value.length < MIN_TEAMS) return

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

.text-warning {
  color: #f57c00;
  font-weight: normal;
  font-size: 14px;
  margin-top: 4px;
}

.text-success {
  color: #4caf50;
  font-weight: 600;
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
