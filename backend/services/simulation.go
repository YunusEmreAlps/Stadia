package services

import (
	"math"
	"math/rand"
	"stadia-backend/models"
	"time"
)

// SimulationService handles match simulation logic
type SimulationService struct {
	rand *rand.Rand
}

// NewSimulationService creates a new simulation service
func NewSimulationService() *SimulationService {
	return &SimulationService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// SimulateMatch simulates a match between two teams based on their power
// Factors considered:
// 1. Team power difference
// 2. Home advantage (home team gets a boost)
// 3. Randomness (for unpredictability)
func (s *SimulationService) SimulateMatch(homeTeam, awayTeam *models.Team) (homeScore, awayScore int) {
	// Calculate effective power with home advantage
	homeAdvantage := 10.0 // Home team gets 10% boost
	homePower := float64(homeTeam.Power) * (1 + homeAdvantage/100)
	awayPower := float64(awayTeam.Power)

	// Calculate expected goals based on power
	homeExpectedGoals := s.calculateExpectedGoals(homePower, awayPower)
	awayExpectedGoals := s.calculateExpectedGoals(awayPower, homePower)

	// Generate actual goals using Poisson-like distribution
	homeScore = s.generateGoals(homeExpectedGoals)
	awayScore = s.generateGoals(awayExpectedGoals)

	return homeScore, awayScore
}

// calculateExpectedGoals calculates expected goals based on team power
func (s *SimulationService) calculateExpectedGoals(attackPower, defensePower float64) float64 {
	// Normalize power values (0-100) to reasonable goal expectations (0-4)
	powerRatio := attackPower / defensePower

	// Base expected goals
	baseGoals := 1.5

	// Adjust based on power ratio
	// Strong team vs weak team: higher expected goals
	// Equal teams: around base goals
	expectedGoals := baseGoals * math.Pow(powerRatio, 0.4)

	// Add some randomness
	randomFactor := 0.8 + s.rand.Float64()*0.4 // 0.8 to 1.2
	expectedGoals *= randomFactor

	// Cap maximum expected goals
	if expectedGoals > 4.5 {
		expectedGoals = 4.5
	}

	return expectedGoals
}

// generateGoals generates actual goals using a Poisson-like distribution
func (s *SimulationService) generateGoals(expectedGoals float64) int {
	// Simple Poisson approximation
	L := math.Exp(-expectedGoals)
	k := 0
	p := 1.0

	for p > L {
		k++
		p *= s.rand.Float64()
	}

	return k - 1
}

// SimulateGoalsProbabilistic generates goals with more controlled distribution
func (s *SimulationService) SimulateGoalsProbabilistic(teamPower int) int {
	// Probability distribution for goals based on team power
	roll := s.rand.Float64() * 100

	powerFactor := float64(teamPower) / 100.0

	// Adjust probabilities based on team power
	if roll < 15*powerFactor {
		return s.rand.Intn(4) + 3 // 3-6 goals (rare)
	} else if roll < 40*powerFactor {
		return 2 // 2 goals
	} else if roll < 70*powerFactor {
		return 1 // 1 goal
	}
	return 0 // 0 goals
}

// CalculateWinProbability calculates the probability of team1 winning against team2
func (s *SimulationService) CalculateWinProbability(team1Power, team2Power int) float64 {
	// Use logistic function to calculate win probability
	powerDiff := float64(team1Power - team2Power)

	// Sigmoid function scaled for football matches
	// A 20-point difference gives about 70% win probability
	// A 40-point difference gives about 90% win probability
	scalingFactor := 0.05
	probability := 1.0 / (1.0 + math.Exp(-scalingFactor*powerDiff))

	return probability
}
