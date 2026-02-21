package services

import (
	"stadia-backend/models"
)

// FixtureService handles fixture generation
type FixtureService struct{}

// NewFixtureService creates a new fixture service
func NewFixtureService() *FixtureService {
	return &FixtureService{}
}

// GenerateRoundRobinFixtures generates a round-robin fixture for all teams
// Each team plays every other team twice (home and away)
func (f *FixtureService) GenerateRoundRobinFixtures(teams []*models.Team) [][]*models.Match {
	if len(teams) < 2 {
		return nil
	}

	fixtures := make([][]*models.Match, 0)
	week := 1

	// First half: Each team plays each other once
	for i := 0; i < len(teams); i++ {
		for j := i + 1; j < len(teams); j++ {
			weekMatches := []*models.Match{
				models.NewMatch(
					teams[i].ID,
					teams[j].ID,
					teams[i].Name,
					teams[j].Name,
					week,
				),
			}
			fixtures = append(fixtures, weekMatches)
			week++
		}
	}

	// Second half: Return fixtures (away becomes home and vice versa)
	firstHalfCount := len(fixtures)
	for i := 0; i < firstHalfCount; i++ {
		originalMatch := fixtures[i][0]
		returnMatch := models.NewMatch(
			originalMatch.AwayTeamID,
			originalMatch.HomeTeamID,
			originalMatch.AwayTeamName,
			originalMatch.HomeTeamName,
			week,
		)
		weekMatches := []*models.Match{returnMatch}
		fixtures = append(fixtures, weekMatches)
		week++
	}

	return fixtures
}

// GenerateFixturesOptimized generates fixtures with better week distribution
// Uses the circle method (Round Robin algorithm) to create balanced fixtures
// For n teams: (n-1) weeks in first half, (n-1) weeks in second half
// Each week has n/2 matches played simultaneously
func (f *FixtureService) GenerateFixturesOptimized(teams []*models.Team) [][]*models.Match {
	n := len(teams)
	if n < 2 {
		return nil
	}

	// Create a copy of teams to work with
	teamList := make([]*models.Team, n)
	copy(teamList, teams)

	// If odd number of teams, add a dummy team (bye week)
	hasDummy := false
	if n%2 != 0 {
		hasDummy = true
		teamList = append(teamList, nil) // nil represents a bye
		n++
	}

	fixtures := make([][]*models.Match, 0)
	week := 1

	// Generate first half fixtures using circle method
	// Keep first team fixed, rotate others clockwise
	for round := 0; round < n-1; round++ {
		weekMatches := make([]*models.Match, 0)

		for i := 0; i < n/2; i++ {
			home := (round + i) % (n - 1)
			away := (n - 1 - i + round) % (n - 1)

			// Adjust for the fixed team
			if i == 0 {
				away = n - 1
			}

			// Skip if either team is dummy (bye week)
			if hasDummy && (teamList[home] == nil || teamList[away] == nil) {
				continue
			}

			// Alternate home/away to balance home advantage
			if round%2 == 0 {
				weekMatches = append(weekMatches, models.NewMatch(
					teamList[home].ID,
					teamList[away].ID,
					teamList[home].Name,
					teamList[away].Name,
					week,
				))
			} else {
				weekMatches = append(weekMatches, models.NewMatch(
					teamList[away].ID,
					teamList[home].ID,
					teamList[away].Name,
					teamList[home].Name,
					week,
				))
			}
		}

		if len(weekMatches) > 0 {
			fixtures = append(fixtures, weekMatches)
			week++
		}
	}

	// Generate second half fixtures (return matches)
	firstHalfCount := len(fixtures)
	for i := 0; i < firstHalfCount; i++ {
		weekMatches := make([]*models.Match, 0)
		for _, match := range fixtures[i] {
			// Swap home and away teams
			returnMatch := models.NewMatch(
				match.AwayTeamID,
				match.HomeTeamID,
				match.AwayTeamName,
				match.HomeTeamName,
				week,
			)
			weekMatches = append(weekMatches, returnMatch)
		}
		fixtures = append(fixtures, weekMatches)
		week++
	}

	return fixtures
}
