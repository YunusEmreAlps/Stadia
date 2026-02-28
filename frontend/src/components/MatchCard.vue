<template>
  <div 
    class="match-card"
    :class="{ 'match-played': isPlayed, 'match-editing': isEditing }"
  >
    <div class="match-content">
      <div class="team home-team">
        <span class="team-name">{{ match.homeTeamName }}</span>
      </div>

      <div class="match-score">
        <template v-if="isPlayed">
          <input 
            v-if="isEditing"
            type="number" 
            v-model.number="editedScores.home"
            min="0"
            class="score-input"
            @keyup.enter="$emit('save')"
            @keyup.esc="$emit('cancel')"
          />
          <span v-else class="score">{{ match.homeScore }}</span>
          
          <span class="separator">-</span>
          
          <input 
            v-if="isEditing"
            type="number" 
            v-model.number="editedScores.away"
            min="0"
            class="score-input"
            @keyup.enter="$emit('save')"
            @keyup.esc="$emit('cancel')"
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

    <div v-if="isPlayed" class="match-footer">
      <span class="result-badge" :class="resultClass">
        {{ resultText }}
      </span>
      
      <div class="match-actions">
        <button 
          v-if="!isEditing"
          @click="$emit('edit')"
          class="btn-edit"
        >
          Edit
        </button>
        <template v-else>
          <button 
            @click="handleSave"
            class="btn-save"
            :disabled="!isValid"
          >
            Save
          </button>
          <button 
            @click="$emit('cancel')"
            class="btn-cancel"
          >
            Cancel
          </button>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, watch } from 'vue'

const props = defineProps({
  match: {
    type: Object,
    required: true
  },
  isEditing: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['edit', 'save', 'cancel'])

const editedScores = ref({
  home: props.match.homeScore,
  away: props.match.awayScore
})

watch(() => props.isEditing, (newVal) => {
  if (newVal) {
    editedScores.value = {
      home: props.match.homeScore,
      away: props.match.awayScore
    }
  }
})

const isPlayed = computed(() => props.match.status === 'played')

const isValid = computed(() => {
  return editedScores.value.home >= 0 && editedScores.value.away >= 0
})

const resultClass = computed(() => {
  if (props.match.homeScore > props.match.awayScore) return 'home-win'
  if (props.match.homeScore < props.match.awayScore) return 'away-win'
  return 'draw'
})

const resultText = computed(() => {
  if (props.match.homeScore > props.match.awayScore) return 'Home Win'
  if (props.match.homeScore < props.match.awayScore) return 'Away Win'
  return 'Draw'
})

const handleSave = () => {
  if (isValid.value) {
    emit('save', editedScores.value)
  }
}
</script>

<style scoped>
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

.match-card.match-played {
  background-color: #f8f9fa;
}

.match-card.match-editing {
  border-color: var(--primary-color);
  background-color: #e8f0fe;
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

.match-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid var(--border-color);
}

.match-actions {
  display: flex;
  gap: 8px;
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

.btn-save:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-cancel {
  background-color: #6c757d;
  color: white;
}

.btn-edit:hover:not(:disabled),
.btn-save:hover:not(:disabled),
.btn-cancel:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

.result-badge {
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
