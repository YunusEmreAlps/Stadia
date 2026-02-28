/**
 * Available teams for the Champions League simulation
 * Each team has a name, power rating (0-100), and country flag code
 */
export const AVAILABLE_TEAMS = [
  { name: 'Manchester City', power: 92, logo: 'GB-ENG' },
  { name: 'Real Madrid', power: 91, logo: 'ES' },
  { name: 'Bayern Munich', power: 90, logo: 'DE' },
  { name: 'Paris Saint-Germain', power: 88, logo: 'FR' },
  { name: 'Liverpool', power: 87, logo: 'GB-ENG' },
  { name: 'Barcelona', power: 86, logo: 'ES' },
  { name: 'Inter Milan', power: 84, logo: 'IT' },
  { name: 'Arsenal', power: 83, logo: 'GB-ENG' },
  { name: 'Napoli', power: 82, logo: 'IT' },
  { name: 'Atletico Madrid', power: 81, logo: 'ES' },
  { name: 'Chelsea', power: 80, logo: 'GB-ENG' },
  { name: 'Borussia Dortmund', power: 79, logo: 'DE' },
  { name: 'Ajax', power: 78, logo: 'NL' },
  { name: 'Porto', power: 77, logo: 'PT' },
  { name: 'RB Leipzig', power: 76, logo: 'DE' },
  { name: 'AC Milan', power: 75, logo: 'IT' },
  { name: 'Beşiktaş', power: 74, logo: 'TR' },
  { name: 'Celtic', power: 73, logo: 'GB-SCT' },
  { name: 'Qarabağ', power: 70, logo: 'AZ' },
  { name: 'Konyaspor', power: 68, logo: 'TR' },
]

/**
 * Minimum number of teams required to start a league
 */
export const MIN_TEAMS = 4

/**
 * Flag image base URL
 */
export const FLAG_BASE_URL = 'https://raw.githubusercontent.com/hampusborgos/country-flags/refs/heads/main/svg'
