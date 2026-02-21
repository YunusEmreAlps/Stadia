package services

import (
	"stadia-backend/models"
	"testing"
)

func TestSimulateMatch(t *testing.T) {
	service := NewSimulationService()

	team1 := models.NewTeam("Strong Team", 90, "")
	team2 := models.NewTeam("Weak Team", 30, "")

	// Run multiple simulations
	strongWins := 0
	draws := 0
	weakWins := 0
	totalMatches := 1000

	for i := 0; i < totalMatches; i++ {
		homeScore, awayScore := service.SimulateMatch(team1, team2)

		if homeScore > awayScore {
			strongWins++
		} else if homeScore == awayScore {
			draws++
		} else {
			weakWins++
		}

		// Ensure scores are non-negative
		if homeScore < 0 || awayScore < 0 {
			t.Errorf("Negative score: home=%d, away=%d", homeScore, awayScore)
		}
	}

	// Strong team should win more than 60% of matches
	strongWinRate := float64(strongWins) / float64(totalMatches)
	if strongWinRate < 0.6 {
		t.Errorf("Strong team win rate too low: %.2f", strongWinRate)
	}

	t.Logf("Strong wins: %d, Draws: %d, Weak wins: %d", strongWins, draws, weakWins)
	t.Logf("Strong team win rate: %.2f%%", strongWinRate*100)
}

func TestCalculateWinProbability(t *testing.T) {
	service := NewSimulationService()

	tests := []struct {
		team1Power int
		team2Power int
		minProb    float64
		maxProb    float64
	}{
		{50, 50, 0.45, 0.55}, // Equal teams
		{70, 50, 0.65, 0.85}, // Moderate advantage
		{90, 30, 0.85, 1.0},  // Strong advantage
		{30, 90, 0.0, 0.15},  // Strong disadvantage
	}

	for _, tt := range tests {
		prob := service.CalculateWinProbability(tt.team1Power, tt.team2Power)

		if prob < tt.minProb || prob > tt.maxProb {
			t.Errorf("Win probability for %d vs %d = %.2f, expected between %.2f and %.2f",
				tt.team1Power, tt.team2Power, prob, tt.minProb, tt.maxProb)
		}
	}
}

func TestCalculateExpectedGoals(t *testing.T) {
	service := NewSimulationService()

	// Test that stronger teams get more expected goals
	strongAttack := 90.0
	weakDefense := 40.0
	expectedGoalsHigh := service.calculateExpectedGoals(strongAttack, weakDefense)

	weakAttack := 40.0
	strongDefense := 90.0
	expectedGoalsLow := service.calculateExpectedGoals(weakAttack, strongDefense)

	if expectedGoalsHigh <= expectedGoalsLow {
		t.Errorf("Expected goals for strong attack (%.2f) should be higher than weak attack (%.2f)",
			expectedGoalsHigh, expectedGoalsLow)
	}

	t.Logf("Strong attack expected goals: %.2f", expectedGoalsHigh)
	t.Logf("Weak attack expected goals: %.2f", expectedGoalsLow)
}

func TestGenerateGoals(t *testing.T) {
	service := NewSimulationService()

	// Generate many goals and check distribution
	expectedGoals := 2.0
	totalGoals := 0
	iterations := 1000

	for i := 0; i < iterations; i++ {
		goals := service.generateGoals(expectedGoals)
		if goals < 0 {
			t.Errorf("Generated negative goals: %d", goals)
		}
		totalGoals += goals
	}

	avgGoals := float64(totalGoals) / float64(iterations)

	// Average should be close to expected (within 20%)
	tolerance := expectedGoals * 0.3
	if avgGoals < expectedGoals-tolerance || avgGoals > expectedGoals+tolerance {
		t.Errorf("Average goals %.2f not close to expected %.2f", avgGoals, expectedGoals)
	}

	t.Logf("Expected goals: %.2f, Average goals: %.2f", expectedGoals, avgGoals)
}

func BenchmarkSimulateMatch(b *testing.B) {
	service := NewSimulationService()
	team1 := models.NewTeam("Team 1", 75, "")
	team2 := models.NewTeam("Team 2", 65, "")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.SimulateMatch(team1, team2)
	}
}
