package services

import (
	"math"
	"stadia-backend/models"
)

// PredictionService handles championship prediction calculations
type PredictionService struct {
	simulationService *SimulationService
}

// NewPredictionService creates a new prediction service
func NewPredictionService() *PredictionService {
	return &PredictionService{
		simulationService: NewSimulationService(),
	}
}

// CalculatePredictions calculates championship probabilities for all teams
// This uses Monte Carlo simulation to predict outcomes based on:
// 1. Current points and standings
// 2. Remaining matches
// 3. Team strengths
// 4. Historical performance in played matches
func (ps *PredictionService) CalculatePredictions(
	teams []*models.Team,
	fixtures [][]*models.Match,
	currentWeek int,
	totalWeeks int,
) map[string]float64 {
	predictions := make(map[string]float64)

	// Initialize all teams with 0 probability
	for _, team := range teams {
		predictions[team.ID] = 0.0
	}

	// If league is finished, winner has 100% probability
	if currentWeek >= totalWeeks {
		winner := ps.getLeagueLeader(teams)
		if winner != nil {
			predictions[winner.ID] = 1.0
		}
		return predictions
	}

	// Run Monte Carlo simulations
	numSimulations := 10000
	wins := make(map[string]int)

	for _, team := range teams {
		wins[team.ID] = 0
	}

	for i := 0; i < numSimulations; i++ {
		winner := ps.simulateRemainingMatches(teams, fixtures, currentWeek, totalWeeks)
		if winner != nil {
			wins[winner.ID]++
		}
	}

	// Calculate probabilities
	for teamID, winCount := range wins {
		predictions[teamID] = float64(winCount) / float64(numSimulations)
	}

	return predictions
}

// simulateRemainingMatches simulates all remaining matches and returns the winner
func (ps *PredictionService) simulateRemainingMatches(
	teams []*models.Team,
	fixtures [][]*models.Match,
	currentWeek int,
	totalWeeks int,
) *models.Team {
	// Create a deep copy of teams to avoid modifying original data
	teamsCopy := make(map[string]*models.Team)
	for _, team := range teams {
		teamsCopy[team.ID] = &models.Team{
			ID:           team.ID,
			Name:         team.Name,
			Power:        team.Power,
			Played:       team.Played,
			Won:          team.Won,
			Drawn:        team.Drawn,
			Lost:         team.Lost,
			GoalsFor:     team.GoalsFor,
			GoalsAgainst: team.GoalsAgainst,
			Points:       team.Points,
		}
	}

	// Simulate remaining matches
	for week := currentWeek + 1; week <= totalWeeks; week++ {
		weekMatches := fixtures[week-1]
		for _, match := range weekMatches {
			if match.IsPlayed() {
				// Use actual result
				homeTeam := teamsCopy[match.HomeTeamID]
				awayTeam := teamsCopy[match.AwayTeamID]
				homeTeam.UpdateStats(match.HomeScore, match.AwayScore)
				awayTeam.UpdateStats(match.AwayScore, match.HomeScore)
			} else {
				// Simulate the match
				homeTeam := teamsCopy[match.HomeTeamID]
				awayTeam := teamsCopy[match.AwayTeamID]
				homeScore, awayScore := ps.simulationService.SimulateMatch(homeTeam, awayTeam)
				homeTeam.UpdateStats(homeScore, awayScore)
				awayTeam.UpdateStats(awayScore, homeScore)
			}
		}
	}

	// Find the winner (team with most points)
	var winner *models.Team
	maxPoints := -1
	teamsList := make([]*models.Team, 0, len(teamsCopy))
	for _, team := range teamsCopy {
		teamsList = append(teamsList, team)
	}

	// Sort teams to get the winner
	for _, team := range teamsList {
		if team.Points > maxPoints {
			maxPoints = team.Points
			winner = team
		} else if team.Points == maxPoints && winner != nil {
			// Tiebreaker: goal difference
			if team.GoalDifference() > winner.GoalDifference() {
				winner = team
			} else if team.GoalDifference() == winner.GoalDifference() {
				// Second tiebreaker: goals for
				if team.GoalsFor > winner.GoalsFor {
					winner = team
				}
			}
		}
	}

	return winner
}

// getLeagueLeader returns the current league leader
func (ps *PredictionService) getLeagueLeader(teams []*models.Team) *models.Team {
	if len(teams) == 0 {
		return nil
	}

	leader := teams[0]
	for _, team := range teams[1:] {
		if team.Points > leader.Points {
			leader = team
		} else if team.Points == leader.Points {
			if team.GoalDifference() > leader.GoalDifference() {
				leader = team
			} else if team.GoalDifference() == leader.GoalDifference() {
				if team.GoalsFor > leader.GoalsFor {
					leader = team
				}
			}
		}
	}

	return leader
}

// CalculateSimplePrediction calculates a simpler prediction based on current form
// This is a faster alternative to Monte Carlo simulation
func (ps *PredictionService) CalculateSimplePrediction(
	teams []*models.Team,
	currentWeek int,
	totalWeeks int,
) map[string]float64 {
	predictions := make(map[string]float64)
	totalProbability := 0.0

	// Calculate a score for each team based on:
	// 1. Current points (60% weight)
	// 2. Goal difference (20% weight)
	// 3. Team power (20% weight)
	scores := make(map[string]float64)

	// remainingWeeks := totalWeeks - currentWeek
	// maxPossiblePoints := float64(remainingWeeks * 3) // Max points from remaining matches

	for _, team := range teams {
		// Normalize current points (0-1)
		currentPointsScore := float64(team.Points) / float64(totalWeeks*3)

		// Normalize goal difference (-1 to 1, capped at ±10)
		gdScore := math.Max(-10, math.Min(10, float64(team.GoalDifference()))) / 10.0
		gdScore = (gdScore + 1) / 2 // Convert to 0-1 range

		// Normalize team power (0-1)
		powerScore := float64(team.Power) / 100.0

		// Combine scores with weights
		score := currentPointsScore*0.6 + gdScore*0.2 + powerScore*0.2

		// Adjust for remaining matches - more weight on current position if fewer matches remain
		matchesPlayedRatio := float64(currentWeek) / float64(totalWeeks)
		score = score * (0.5 + 0.5*matchesPlayedRatio)

		scores[team.ID] = math.Max(0.01, score) // Minimum score to avoid division by zero
		totalProbability += scores[team.ID]
	}

	// Normalize to probabilities (sum to 1)
	for teamID, score := range scores {
		predictions[teamID] = score / totalProbability
	}

	return predictions
}
