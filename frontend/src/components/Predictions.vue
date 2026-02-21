<template>
  <div class="card predictions-card">
    <div class="card-header">
      <h2>🔮 Championship Predictions</h2>
      <span class="badge badge-warning">Week {{ predictions?.week }}</span>
    </div>

    <p class="description">
      Based on current standings, remaining fixtures, and team strengths, 
      here are the predicted chances of winning the group:
    </p>

    <div v-if="predictions && predictions.predictions.length > 0" class="predictions-list">
      <div 
        v-for="(pred, index) in predictions.predictions" 
        :key="pred.teamId"
        class="prediction-item"
        :class="getPredictionClass(index)"
      >
        <div class="prediction-header">
          <div class="team-info">
            <span class="position">{{ index + 1 }}</span>
            <span class="team-name">{{ pred.teamName }}</span>
          </div>
          <span class="probability">{{ formatProbability(pred.probability) }}%</span>
        </div>
        <div class="probability-bar">
          <div 
            class="probability-fill" 
            :style="{ width: pred.probability + '%' }"
            :class="getBarClass(pred.probability)"
          ></div>
        </div>
      </div>
    </div>

    <div v-else class="no-predictions">
      <p>Predictions will be available from Week 4 onwards...</p>
    </div>

    <div class="prediction-info">
      <p>
        💡 Predictions are calculated using Monte Carlo simulation with 10,000 iterations, 
        considering current points, goal difference, remaining fixtures, and team strength.
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useLeagueStore } from '@/stores/league'

const leagueStore = useLeagueStore()

const predictions = computed(() => leagueStore.predictions)

const formatProbability = (prob) => {
  return prob.toFixed(1)
}

const getPredictionClass = (index) => {
  if (index === 0) return 'favorite'
  return ''
}

const getBarClass = (prob) => {
  if (prob >= 70) return 'high'
  if (prob >= 40) return 'medium'
  return 'low'
}
</script>

<style scoped>
.predictions-card {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  color: white;
}

.card-header h2 {
  color: white;
}

.description {
  margin: 16px 0 24px;
  opacity: 0.95;
  font-size: 14px;
  line-height: 1.6;
}

.predictions-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.prediction-item {
  background-color: rgba(255, 255, 255, 0.15);
  border-radius: 12px;
  padding: 16px;
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
}

.prediction-item:hover {
  background-color: rgba(255, 255, 255, 0.25);
  transform: translateX(4px);
}

.prediction-item.favorite {
  background-color: rgba(255, 255, 255, 0.25);
  border: 2px solid rgba(255, 255, 255, 0.5);
}

.prediction-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.team-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.position {
  display: inline-block;
  width: 32px;
  height: 32px;
  line-height: 32px;
  text-align: center;
  background-color: rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  font-weight: 700;
  font-size: 14px;
}

.team-name {
  font-size: 16px;
  font-weight: 600;
}

.probability {
  font-size: 24px;
  font-weight: 700;
}

.probability-bar {
  height: 12px;
  background-color: rgba(255, 255, 255, 0.3);
  border-radius: 6px;
  overflow: hidden;
}

.probability-fill {
  height: 100%;
  border-radius: 6px;
  transition: width 0.5s ease;
  position: relative;
}

.probability-fill.high {
  background: linear-gradient(90deg, #4caf50, #8bc34a);
}

.probability-fill.medium {
  background: linear-gradient(90deg, #ff9800, #ffc107);
}

.probability-fill.low {
  background: linear-gradient(90deg, #f44336, #e91e63);
}

.no-predictions {
  text-align: center;
  padding: 40px 20px;
  background-color: rgba(255, 255, 255, 0.15);
  border-radius: 12px;
  font-size: 16px;
}

.prediction-info {
  margin-top: 24px;
  padding-top: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.3);
  font-size: 13px;
  opacity: 0.9;
  line-height: 1.6;
}

@media (max-width: 768px) {
  .team-name {
    font-size: 14px;
  }

  .probability {
    font-size: 20px;
  }
}
</style>
