<template>
  <div class="card">
    <div class="card-header">
      <h2>📅 Match Schedule</h2>
    </div>

    <div v-if="league && league.fixtures" class="weeks-container">
      <div 
        v-for="(weekMatches, weekIndex) in league.fixtures" 
        :key="weekIndex"
        class="week-section"
        :class="{ 'current-week': weekIndex + 1 === league.currentWeek, 'played': weekIndex < league.currentWeek }"
      >
        <div class="week-header">
          <h3>Week {{ weekIndex + 1 }}</h3>
          <span v-if="weekIndex < league.currentWeek" class="week-status played">✓ Played</span>
          <span v-else-if="weekIndex + 1 === league.currentWeek" class="week-status current">▶️ Current</span>
          <span v-else class="week-status upcoming">⏱️ Upcoming</span>
        </div>

        <div class="matches-list">
          <div 
            v-for="match in weekMatches" 
            :key="match.id"
            class="match-card"
            :class="{ played: match.status === 'played' }"
          >
            <div class="match-content">
              <div class="team home-team">
                <span class="team-name">{{ match.homeTeamName }}</span>
              </div>

              <div class="match-score">
                <template v-if="match.status === 'played'">
                  <input 
                    v-if="editingMatch === match.id"
                    type="number" 
                    v-model.number="editScores[match.id].home"
                    min="0"
                    class="score-input"
                  />
                  <span v-else class="score">{{ match.homeScore }}</span>
                  
                  <span class="separator">-</span>
                  
                  <input 
                    v-if="editingMatch === match.id"
                    type="number" 
                    v-model.number="editScores[match.id].away"
                    min="0"
                    class="score-input"
                  />
                  <span v-else class="score">{{ match.awayScore }}</span>
                </template>
                <template v-else>
                  <span class="vs">VS</span>
                </template>
              </div>

              <div class="team away-team">
                <span class="team-name">{{ match.awayTeamName }}</span>
              </div>
            </div>

            <div v-if="match.status === 'played'" class="match-actions">
              <button 
                v-if="editingMatch !== match.id"
                @click="startEdit(match)"
                class="btn-edit"
              >
                ✏️ Edit
              </button>
              <template v-else>
                <button 
                  @click="saveEdit(match.id)"
                  class="btn-save"
                  :disabled="loading"
                >
                  ✓ Save
                </button>
                <button 
                  @click="cancelEdit"
                  class="btn-cancel"
                >
                  ✕ Cancel
                </button>
              </template>
            </div>

            <div v-if="match.status === 'played'" class="match-result">
              <span 
                class="result-badge"
                :class="getResultClass(match)"
              >
                {{ getResultText(match) }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useLeagueStore } from '@/stores/league'

const leagueStore = useLeagueStore()

const league = computed(() => leagueStore.league)
const loading = computed(() => leagueStore.loading)

const editingMatch = ref(null)
const editScores = ref({})

const startEdit = (match) => {
  editingMatch.value = match.id
  editScores.value[match.id] = {
    home: match.homeScore,
    away: match.awayScore
  }
}

const cancelEdit = () => {
  editingMatch.value = null
  editScores.value = {}
}

const saveEdit = async (matchId) => {
  const scores = editScores.value[matchId]
  if (scores.home < 0 || scores.away < 0) {
    alert('Scores cannot be negative')
    return
  }

  try {
    await leagueStore.updateMatchResult(matchId, scores.home, scores.away)
    editingMatch.value = null
    editScores.value = {}
  } catch (error) {
    console.error('Failed to update match:', error)
  }
}

const getResultClass = (match) => {
  if (match.homeScore > match.awayScore) return 'home-win'
  if (match.homeScore < match.awayScore) return 'away-win'
  return 'draw'
}

const getResultText = (match) => {
  if (match.homeScore > match.awayScore) return 'Home Win'
  if (match.homeScore < match.awayScore) return 'Away Win'
  return 'Draw'
}
</script>

<style scoped>
.weeks-container {
  display: flex;
  flex-direction: column;
  gap: 24px;
  margin-top: 20px;
}

.week-section {
  border: 2px solid var(--border-color);
  border-radius: 12px;
  padding: 20px;
  transition: all 0.3s ease;
}

.week-section.current-week {
  border-color: var(--primary-color);
  background-color: #e8f0fe;
}

.week-section.played {
  opacity: 0.8;
}

.week-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 2px solid var(--border-color);
}

.week-header h3 {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-primary);
}

.week-status {
  padding: 6px 12px;
  border-radius: 16px;
  font-size: 12px;
  font-weight: 600;
}

.week-status.played {
  background-color: #d4edda;
  color: #155724;
}

.week-status.current {
  background-color: #cce5ff;
  color: #004085;
}

.week-status.upcoming {
  background-color: #fff3cd;
  color: #856404;
}

.matches-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.match-card {
  background-color: white;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  padding: 16px;
  transition: all 0.3s ease;
}

.match-card:hover {
  box-shadow: var(--shadow-sm);
  transform: translateX(4px);
}

.match-card.played {
  background-color: #f8f9fa;
}

.match-content {
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  align-items: center;
  gap: 16px;
}

.team {
  display: flex;
  align-items: center;
}

.home-team {
  justify-content: flex-end;
}

.away-team {
  justify-content: flex-start;
}

.team-name {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-primary);
}

.match-score {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 24px;
  font-weight: 700;
  min-width: 120px;
  justify-content: center;
}

.score {
  color: var(--primary-color);
  min-width: 32px;
  text-align: center;
}

.separator {
  color: var(--text-secondary);
}

.vs {
  color: var(--text-secondary);
  font-size: 14px;
  font-weight: 600;
}

.score-input {
  width: 60px;
  padding: 8px;
  text-align: center;
  font-size: 20px;
  font-weight: 700;
  border: 2px solid var(--primary-color);
  border-radius: 6px;
  outline: none;
}

.match-actions {
  display: flex;
  gap: 8px;
  justify-content: center;
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid var(--border-color);
}

.btn-edit,
.btn-save,
.btn-cancel {
  padding: 6px 16px;
  font-size: 12px;
  border-radius: 6px;
  border: none;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s ease;
}

.btn-edit {
  background-color: var(--primary-color);
  color: white;
}

.btn-save {
  background-color: var(--secondary-color);
  color: white;
}

.btn-cancel {
  background-color: #6c757d;
  color: white;
}

.btn-edit:hover,
.btn-save:hover,
.btn-cancel:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

.match-result {
  margin-top: 8px;
  text-align: center;
}

.result-badge {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 11px;
  font-weight: 700;
  text-transform: uppercase;
}

.result-badge.home-win {
  background-color: #d4edda;
  color: #155724;
}

.result-badge.away-win {
  background-color: #f8d7da;
  color: #721c24;
}

.result-badge.draw {
  background-color: #fff3cd;
  color: #856404;
}

@media (max-width: 768px) {
  .match-content {
    grid-template-columns: 1fr;
    gap: 8px;
    text-align: center;
  }

  .home-team,
  .away-team {
    justify-content: center;
  }

  .team-name {
    font-size: 14px;
  }

  .match-score {
    font-size: 20px;
  }
}
</style>
