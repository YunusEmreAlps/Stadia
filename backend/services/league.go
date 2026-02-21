package services

import (
	"errors"
	"sort"
	"stadia-backend/models"
)

// LeagueService manages league operations
type LeagueService struct {
	league            *models.League
	simulationService *SimulationService
	fixtureService    *FixtureService
	predictionService *PredictionService
}

// NewLeagueService creates a new league service
func NewLeagueService() *LeagueService {
	return &LeagueService{
		league:            models.NewLeague("Champions League Group Stage"),
		simulationService: NewSimulationService(),
		fixtureService:    NewFixtureService(),
		predictionService: NewPredictionService(),
	}
}

// InitializeLeague initializes the league with teams
func (ls *LeagueService) InitializeLeague(teams []*models.Team) error {
	if len(teams) < 2 {
		return errors.New("at least 2 teams are required")
	}

	// Reset league
	ls.league = models.NewLeague("Champions League Group Stage")

	// Add teams
	for _, team := range teams {
		ls.league.AddTeam(team)
	}

	// Generate fixtures
	ls.league.Fixtures = ls.fixtureService.GenerateFixturesOptimized(teams)
	ls.league.TotalWeeks = len(ls.league.Fixtures)
	ls.league.CurrentWeek = 0

	return nil
}

// GetLeague returns the current league state
func (ls *LeagueService) GetLeague() *models.League {
	return ls.league
}

// GetStandings returns the sorted league table
func (ls *LeagueService) GetStandings() []*models.Team {
	teams := ls.league.GetTeamsList()

	// Sort by: Points (desc), Goal Difference (desc), Goals For (desc), Name (asc)
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Points != teams[j].Points {
			return teams[i].Points > teams[j].Points
		}
		if teams[i].GoalDifference() != teams[j].GoalDifference() {
			return teams[i].GoalDifference() > teams[j].GoalDifference()
		}
		if teams[i].GoalsFor != teams[j].GoalsFor {
			return teams[i].GoalsFor > teams[j].GoalsFor
		}
		return teams[i].Name < teams[j].Name
	})

	return teams
}

// PlayNextWeek simulates all matches in the next week
func (ls *LeagueService) PlayNextWeek() error {
	if ls.league.CurrentWeek >= ls.league.TotalWeeks {
		return errors.New("all weeks have been played")
	}

	ls.league.CurrentWeek++
	matches := ls.league.GetMatchesByWeek(ls.league.CurrentWeek)

	for _, match := range matches {
		if match.IsPlayed() {
			continue
		}

		homeTeam := ls.league.GetTeam(match.HomeTeamID)
		awayTeam := ls.league.GetTeam(match.AwayTeamID)

		homeScore, awayScore := ls.simulationService.SimulateMatch(homeTeam, awayTeam)

		match.SetResult(homeScore, awayScore)
		homeTeam.UpdateStats(homeScore, awayScore)
		awayTeam.UpdateStats(awayScore, homeScore)
	}

	// Update predictions if we're past week 3
	if ls.league.CurrentWeek >= 3 {
		ls.updatePredictions()
	}

	return nil
}

// PlayAllWeeks simulates all remaining weeks
func (ls *LeagueService) PlayAllWeeks() error {
	for ls.league.CurrentWeek < ls.league.TotalWeeks {
		err := ls.PlayNextWeek()
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateMatchResult manually updates a match result
func (ls *LeagueService) UpdateMatchResult(matchID string, homeScore, awayScore int) error {
	if homeScore < 0 || awayScore < 0 {
		return errors.New("scores cannot be negative")
	}

	// Find the match
	var targetMatch *models.Match
	for _, weekMatches := range ls.league.Fixtures {
		for _, match := range weekMatches {
			if match.ID == matchID {
				targetMatch = match
				break
			}
		}
		if targetMatch != nil {
			break
		}
	}

	if targetMatch == nil {
		return errors.New("match not found")
	}

	// Get teams
	homeTeam := ls.league.GetTeam(targetMatch.HomeTeamID)
	awayTeam := ls.league.GetTeam(targetMatch.AwayTeamID)

	// If match was already played, revert the old stats
	if targetMatch.IsPlayed() {
		ls.revertMatchStats(homeTeam, awayTeam, targetMatch.HomeScore, targetMatch.AwayScore)
	}

	// Update match
	targetMatch.SetResult(homeScore, awayScore)

	// Update team stats
	homeTeam.UpdateStats(homeScore, awayScore)
	awayTeam.UpdateStats(awayScore, homeScore)

	// Update predictions if applicable
	if ls.league.CurrentWeek >= 3 {
		ls.updatePredictions()
	}

	return nil
}

// revertMatchStats reverts team statistics for a match
func (ls *LeagueService) revertMatchStats(homeTeam, awayTeam *models.Team, homeScore, awayScore int) {
	// Revert home team stats
	homeTeam.Played--
	homeTeam.GoalsFor -= homeScore
	homeTeam.GoalsAgainst -= awayScore
	if homeScore > awayScore {
		homeTeam.Won--
		homeTeam.Points -= 3
	} else if homeScore == awayScore {
		homeTeam.Drawn--
		homeTeam.Points -= 1
	} else {
		homeTeam.Lost--
	}

	// Revert away team stats
	awayTeam.Played--
	awayTeam.GoalsFor -= awayScore
	awayTeam.GoalsAgainst -= homeScore
	if awayScore > homeScore {
		awayTeam.Won--
		awayTeam.Points -= 3
	} else if awayScore == homeScore {
		awayTeam.Drawn--
		awayTeam.Points -= 1
	} else {
		awayTeam.Lost--
	}
}

// ResetLeague resets the league to its initial state
func (ls *LeagueService) ResetLeague() error {
	if ls.league == nil {
		return errors.New("league not initialized")
	}

	// Reset all team stats
	for _, team := range ls.league.Teams {
		team.ResetStats()
	}

	// Reset all matches
	for _, weekMatches := range ls.league.Fixtures {
		for _, match := range weekMatches {
			match.HomeScore = 0
			match.AwayScore = 0
			match.Status = models.StatusNotPlayed
		}
	}

	ls.league.CurrentWeek = 0
	ls.league.Predictions = make(map[string]float64)

	return nil
}

// updatePredictions updates championship predictions
func (ls *LeagueService) updatePredictions() {
	teams := ls.league.GetTeamsList()
	predictions := ls.predictionService.CalculatePredictions(
		teams,
		ls.league.Fixtures,
		ls.league.CurrentWeek,
		ls.league.TotalWeeks,
	)

	ls.league.Predictions = predictions
}

// GetPredictions returns current predictions
func (ls *LeagueService) GetPredictions() *models.PredictionResponse {
	predictions := make([]models.Prediction, 0)

	for teamID, probability := range ls.league.Predictions {
		team := ls.league.GetTeam(teamID)
		if team != nil {
			predictions = append(predictions, models.Prediction{
				TeamID:      teamID,
				TeamName:    team.Name,
				Probability: probability * 100, // Convert to percentage
			})
		}
	}

	// Sort by probability descending
	sort.Slice(predictions, func(i, j int) bool {
		return predictions[i].Probability > predictions[j].Probability
	})

	return &models.PredictionResponse{
		Week:        ls.league.CurrentWeek,
		Predictions: predictions,
	}
}
