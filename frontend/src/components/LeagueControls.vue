<template>
  <div class="card controls-card">
    <div class="card-header">
      <div class="header-content">
        <div class="icon-wrapper">
          <div v-if="isFinished" class="control-icon">🏆</div>
          <div v-if="!isFinished" class="progress-ring">
            <svg width="60" height="60">
              <circle
                cx="30"
                cy="30"
                r="25"
                stroke="rgba(255, 255, 255, 0.2)"
                stroke-width="4"
                fill="none"
              />
              <circle
                cx="30"
                cy="30"
                r="25"
                stroke="#fff"
                stroke-width="4"
                fill="none"
                stroke-dasharray="157"
                :stroke-dashoffset="157 - (157 * currentWeek / totalWeeks)"
                stroke-linecap="round"
                transform="rotate(-90 30 30)"
              />
            </svg>
            <span class="progress-percent">{{ Math.round((currentWeek / totalWeeks) * 100) }}%</span>
          </div>
        </div>
        <div class="header-text">
          <h2>League Controls</h2>
          <p class="week-info">
            <span class="week-badge">Week {{ currentWeek }}</span>
            <span class="separator">/</span>
            <span>{{ totalWeeks }}</span>
          </p>
        </div>
      </div>
    </div>

    <div class="controls-grid">
      <button 
        class="control-btn primary-btn"
        @click="playNextWeek"
        :disabled="!canPlayNext || loading || isAutoPlaying"
      >
        <span class="btn-icon">▶️</span>
        <span class="btn-text">Play Next Week</span>
      </button>

      <button 
        class="control-btn secondary-btn"
        @click="toggleAutoPlay"
        :disabled="!canPlayNext || loading"
      >
        <span class="btn-icon">{{ isAutoPlaying ? '⏸️' : '⏩' }}</span>
        <span class="btn-text">{{ isAutoPlaying ? 'Stop Auto Play' : 'Auto Play' }}</span>
      </button>

      <button 
        class="control-btn danger-btn"
        @click="resetLeague"
        :disabled="loading || isAutoPlaying"
      >
        <span class="btn-icon">🔄</span>
        <span class="btn-text">Reset League</span>
      </button>
    </div>

    <div v-if="loading" class="loading-bar">
      <div class="loading-progress"></div>
    </div>

    <div v-if="isFinished" class="finished-message">
      <h3>League Finished!</h3>
      <p>Check the final standings above</p>
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
  background: linear-gradient(135deg, #11998e 0%, #1a7a5e 50%, #38ef7d 100%);
  color: white;
  box-shadow: 0 10px 40px rgba(102, 126, 234, 0.3);
  position: relative;
  overflow: hidden;
}

.controls-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23ffffff' fill-opacity='0.05'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
  opacity: 0.1;
  z-index: 0;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
  z-index: 1;
}

.header-content {
  display: flex;
  align-items: center;
  gap: 16px;
}

.icon-wrapper {
  width: 56px;
  height: 56px;
  backdrop-filter: blur(10px);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.control-icon {
  font-size: 28px;
  width: 56px;
  height: 56px;
  backdrop-filter: blur(10px);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.2);
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
}

.header-text h2 {
  color: white;
  margin: 0;
  font-size: 24px;
  font-weight: 700;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.week-info {
  margin-top: 6px;
  font-size: 15px;
  opacity: 0.95;
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
  color: white;
}

.week-badge {
  background: rgba(255, 255, 255, 0.25);
  padding: 4px 12px;
  border-radius: 20px;
  font-weight: 600;
  backdrop-filter: blur(10px);
}

.separator {
  opacity: 0.6;
}

.progress-ring {
  position: relative;
  width: 60px;
  height: 60px;
}

.progress-ring svg {
  transform: rotate(0deg);
  transition: all 0.3s ease;
}

.progress-ring circle {
  transition: stroke-dashoffset 0.5s ease;
}

.progress-percent {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-size: 13px;
  font-weight: 700;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
  color:white;
}

.controls-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 12px;
  margin-top: 24px;
  position: relative;
  z-index: 1;
}

.control-btn {
  padding: 16px 24px;
  font-size: 15px;
  font-weight: 600;
  border-radius: 12px;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
  position: relative;
  overflow: hidden;
}

.control-btn::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 0;
  height: 0;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.3);
  transform: translate(-50%, -50%);
  transition: width 0.6s, height 0.6s;
}

.primary-btn::before {
  background: rgba(102, 126, 234, 0.2);
}

.secondary-btn::before,
.danger-btn::before {
  background: rgba(255, 255, 255, 0.2);
}

.control-btn:hover::before {
  width: 300px;
  height: 300px;
}

.btn-icon {
  font-size: 18px;
  position: relative;
  z-index: 1;
}

.btn-text {
  position: relative;
  z-index: 1;
}

.primary-btn {
  background: rgba(255, 255, 255, 0.95);
  color: #667eea;
}

.primary-btn:hover:not(:disabled) {
  background: white;
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.2);
}

.secondary-btn {
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.9) 0%, rgba(5, 150, 105, 0.9) 100%);
  color: white;
  border: 2px solid rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(10px);
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.secondary-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, rgba(16, 185, 129, 1) 0%, rgba(5, 150, 105, 1) 100%);
  border-color: rgba(255, 255, 255, 0.8);
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(16, 185, 129, 0.4);
}

.danger-btn {
  background: linear-gradient(135deg, rgba(239, 68, 68, 0.9) 0%, rgba(220, 38, 38, 0.9) 100%);
  color: white;
  border: 2px solid rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(10px);
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.danger-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, rgba(239, 68, 68, 1) 0%, rgba(220, 38, 38, 1) 100%);
  border-color: rgba(255, 255, 255, 0.8);
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(239, 68, 68, 0.5);
}

.control-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
  filter: grayscale(20%);
}

.loading-bar {
  width: 100%;
  height: 3px;
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 2px;
  margin-top: 24px;
  overflow: hidden;
  position: relative;
  z-index: 1;
}

.loading-progress {
  height: 100%;
  background: linear-gradient(90deg, transparent, white, transparent);
  animation: loading 1.5s ease-in-out infinite;
}

@keyframes loading {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(200%);
  }
}

.finished-message {
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  padding: 24px;
  border-radius: 16px;
  text-align: center;
  margin-top: 24px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  position: relative;
  z-index: 1;
}

@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

.finished-message h3 {
  margin: 0 0 8px 0;
  font-size: 24px;
  font-weight: 700;
}

.finished-message p {
  margin: 0;
  font-size: 15px;
  opacity: 0.9;
}

@media (max-width: 768px) {
  .controls-grid {
    grid-template-columns: 1fr;
  }

  .card-header {
    flex-direction: column;
    gap: 16px;
  }

  .progress-ring {
    order: -1;
  }

  .icon-wrapper {
    width: 48px;
    height: 48px;
  }

  .control-icon {
    font-size: 24px;
  }
}
</style>
