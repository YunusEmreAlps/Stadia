<template>
  <div 
    class="week-section"
    :class="weekStatusClass"
  >
    <div class="week-header">
      <h3>Week {{ weekNumber }}</h3>
      <span class="week-status" :class="statusBadgeClass">
        {{ statusText }}
      </span>
    </div>

    <div class="matches-list">
      <MatchCard
        v-for="match in matches"
        :key="match.id"
        :match="match"
        :isEditing="editingMatchId === match.id"
        @edit="handleEdit(match.id)"
        @save="handleSave(match.id, $event)"
        @cancel="handleCancelEdit"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import MatchCard from './MatchCard.vue'

const props = defineProps({
  weekNumber: {
    type: Number,
    required: true
  },
  matches: {
    type: Array,
    required: true
  },
  currentWeek: {
    type: Number,
    required: true
  }
})

const emit = defineEmits(['updateMatch'])

const editingMatchId = ref(null)

const weekStatus = computed(() => {
  if (props.weekNumber < props.currentWeek) return 'played'
  if (props.weekNumber === props.currentWeek) return 'current'
  return 'upcoming'
})

const weekStatusClass = computed(() => {
  return {
    'current-week': weekStatus.value === 'current',
    'played': weekStatus.value === 'played'
  }
})

const statusBadgeClass = computed(() => weekStatus.value)

const statusText = computed(() => {
  const statusMap = {
    played: '✓ Played',
    current: '▶️ Current',
    upcoming: '⏱️ Upcoming'
  }
  return statusMap[weekStatus.value] || ''
})

const handleEdit = (matchId) => {
  editingMatchId.value = matchId
}

const handleSave = (matchId, scores) => {
  emit('updateMatch', matchId, scores.home, scores.away)
  editingMatchId.value = null
}

const handleCancelEdit = () => {
  editingMatchId.value = null
}
</script>

<style scoped>
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
</style>
