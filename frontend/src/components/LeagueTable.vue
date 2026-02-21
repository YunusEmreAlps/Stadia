<template>
  <div class="card">
    <div class="card-header">
      <h2>📊 League Table</h2>
      <span class="badge badge-info">Week {{ currentWeek }}</span>
    </div>

    <div class="table-wrapper">
      <table class="league-table">
        <thead>
          <tr>
            <th class="text-center">#</th>
            <th>Team</th>
            <th class="text-center" title="Played">P</th>
            <th class="text-center" title="Won">W</th>
            <th class="text-center" title="Drawn">D</th>
            <th class="text-center" title="Lost">L</th>
            <th class="text-center" title="Goals For">GF</th>
            <th class="text-center" title="Goals Against">GA</th>
            <th class="text-center" title="Goal Difference">GD</th>
            <th class="text-center" title="Points">Pts</th>
          </tr>
        </thead>
        <tbody>
          <tr 
            v-for="(team, index) in standings" 
            :key="team.id"
            :class="getRowClass(index)"
          >
            <td class="text-center position-cell">
              <span class="position-badge" :class="getPositionBadgeClass(index)">
                {{ index + 1 }}
              </span>
            </td>
            <td class="team-name-cell">
              <span class="team-logo" v-if="team.logo">{{ team.logo }}</span>
              <strong>{{ team.name }}</strong>
            </td>
            <td class="text-center">{{ team.played }}</td>
            <td class="text-center">{{ team.won }}</td>
            <td class="text-center">{{ team.drawn }}</td>
            <td class="text-center">{{ team.lost }}</td>
            <td class="text-center">{{ team.goalsFor }}</td>
            <td class="text-center">{{ team.goalsAgainst }}</td>
            <td class="text-center" :class="getGDClass(team.goalsFor - team.goalsAgainst)">
              {{ formatGoalDifference(team.goalsFor - team.goalsAgainst) }}
            </td>
            <td class="text-center points-cell">
              <strong>{{ team.points }}</strong>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="legend">
      <div class="legend-item">
        <span class="legend-badge qualified"></span>
        <span>Top 2 qualify for knockout stage</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useLeagueStore } from '@/stores/league'

const leagueStore = useLeagueStore()

const standings = computed(() => leagueStore.standings)
const currentWeek = computed(() => leagueStore.currentWeek)

const getRowClass = (index) => {
  if (index < 2) return 'qualified'
  return ''
}

const getPositionBadgeClass = (index) => {
  if (index === 0) return 'gold'
  if (index === 1) return 'silver'
  return ''
}

const getGDClass = (gd) => {
  if (gd > 0) return 'positive'
  if (gd < 0) return 'negative'
  return ''
}

const formatGoalDifference = (gd) => {
  if (gd > 0) return `+${gd}`
  return gd.toString()
}
</script>

<style scoped>
.table-wrapper {
  overflow-x: auto;
}

.league-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

.league-table thead {
  background-color: #f8f9fa;
}

.league-table th {
  padding: 12px 8px;
  font-weight: 600;
  text-transform: uppercase;
  font-size: 11px;
  color: var(--text-secondary);
  border-bottom: 2px solid var(--border-color);
}

.league-table td {
  padding: 14px 8px;
  border-bottom: 1px solid var(--border-color);
}

.league-table tbody tr {
  transition: background-color 0.2s ease;
}

.league-table tbody tr:hover {
  background-color: #f8f9fa;
}

.league-table tbody tr.qualified {
  background-color: #e8f5e9;
}

.league-table tbody tr.qualified:hover {
  background-color: #c8e6c9;
}

.position-cell {
  font-weight: 700;
}

.position-badge {
  display: inline-block;
  width: 28px;
  height: 28px;
  line-height: 28px;
  border-radius: 50%;
  background-color: #e0e0e0;
  font-weight: 700;
}

.position-badge.gold {
  background: linear-gradient(135deg, #ffd700, #ffed4e);
  color: #875a00;
}

.position-badge.silver {
  background: linear-gradient(135deg, #c0c0c0, #e8e8e8);
  color: #555;
}

.team-name-cell {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 15px;
}

.team-logo {
  font-size: 24px;
}

.points-cell {
  font-size: 16px;
  color: var(--primary-color);
}

.positive {
  color: #2e7d32;
  font-weight: 600;
}

.negative {
  color: #c62828;
  font-weight: 600;
}

.legend {
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid var(--border-color);
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: var(--text-secondary);
}

.legend-badge {
  width: 20px;
  height: 20px;
  border-radius: 4px;
}

.legend-badge.qualified {
  background-color: #a5d6a7;
}

@media (max-width: 768px) {
  .league-table {
    font-size: 12px;
  }

  .league-table th,
  .league-table td {
    padding: 8px 4px;
  }

  .team-logo {
    font-size: 20px;
  }

  .team-name-cell {
    font-size: 13px;
  }
}
</style>
