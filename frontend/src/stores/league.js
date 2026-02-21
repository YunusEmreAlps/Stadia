import { defineStore } from 'pinia'
import { leagueApi } from '@/services/api'

export const useLeagueStore = defineStore('league', {
  state: () => ({
    league: null,
    standings: [],
    predictions: null,
    loading: false,
    error: null,
    autoPlayInterval: null
  }),

  getters: {
    currentWeek: (state) => state.league?.currentWeek || 0,
    totalWeeks: (state) => state.league?.totalWeeks || 0,
    isFinished: (state) => state.league?.currentWeek >= state.league?.totalWeeks,
    canPlayNext: (state) => state.league && !state.isFinished,
    showPredictions: (state) => state.league?.currentWeek >= 3
  },

  actions: {
    async initializeLeague(teams) {
      this.loading = true
      this.error = null
      try {
        const response = await leagueApi.initialize(teams)
        this.league = response.data.league
        await this.refreshData()
      } catch (error) {
        this.error = error.response?.data?.error || 'Failed to initialize league'
        throw error
      } finally {
        this.loading = false
      }
    },

    async refreshData() {
      try {
        const [leagueRes, standingsRes] = await Promise.all([
          leagueApi.getLeague(),
          leagueApi.getStandings()
        ])
        
        this.league = leagueRes.data
        this.standings = standingsRes.data.standings

        // Fetch predictions if available
        if (this.showPredictions) {
          const predRes = await leagueApi.getPredictions()
          this.predictions = predRes.data
        }
      } catch (error) {
        this.error = error.response?.data?.error || 'Failed to refresh data'
      }
    },

    async playNextWeek() {
      this.loading = true
      this.error = null
      try {
        await leagueApi.playNextWeek()
        await this.refreshData()
      } catch (error) {
        this.error = error.response?.data?.error || 'Failed to play next week'
        throw error
      } finally {
        this.loading = false
      }
    },

    async playAllWeeks() {
      this.loading = true
      this.error = null
      try {
        await leagueApi.playAllWeeks()
        await this.refreshData()
      } catch (error) {
        this.error = error.response?.data?.error || 'Failed to play all weeks'
        throw error
      } finally {
        this.loading = false
      }
    },

    async updateMatchResult(matchId, homeScore, awayScore) {
      this.loading = true
      this.error = null
      try {
        await leagueApi.updateMatch(matchId, homeScore, awayScore)
        await this.refreshData()
      } catch (error) {
        this.error = error.response?.data?.error || 'Failed to update match'
        throw error
      } finally {
        this.loading = false
      }
    },

    async resetLeague() {
      this.loading = true
      this.error = null
      try {
        await leagueApi.reset()
        await this.refreshData()
      } catch (error) {
        this.error = error.response?.data?.error || 'Failed to reset league'
        throw error
      } finally {
        this.loading = false
      }
    },

    startAutoPlay(delay = 2000) {
      if (this.autoPlayInterval) return

      this.autoPlayInterval = setInterval(async () => {
        if (this.canPlayNext && !this.loading) {
          try {
            await this.playNextWeek()
          } catch (error) {
            this.stopAutoPlay()
          }
        } else {
          this.stopAutoPlay()
        }
      }, delay)
    },

    stopAutoPlay() {
      if (this.autoPlayInterval) {
        clearInterval(this.autoPlayInterval)
        this.autoPlayInterval = null
      }
    }
  }
})
