import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8000/api'

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

export const leagueApi = {
  // Initialize league with teams
  initialize(teams) {
    return api.post('/league/initialize', { teams })
  },

  // Get league state
  getLeague() {
    return api.get('/league')
  },

  // Get standings
  getStandings() {
    return api.get('/league/standings')
  },

  // Play next week
  playNextWeek() {
    return api.post('/league/play-next-week')
  },

  // Play all weeks
  playAllWeeks() {
    return api.post('/league/play-all-weeks')
  },

  // Update match result
  updateMatch(matchId, homeScore, awayScore) {
    return api.put(`/league/match/${matchId}`, {
      homeScore,
      awayScore
    })
  },

  // Reset league
  reset() {
    return api.post('/league/reset')
  },

  // Get predictions
  getPredictions() {
    return api.get('/league/predictions')
  }
}

export default api
