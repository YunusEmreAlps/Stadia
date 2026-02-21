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
// For 4 teams: 3 matches per week in first round, 3 in second round (6 weeks total)
func (f *FixtureService) GenerateFixturesOptimized(teams []*models.Team) [][]*models.Match {
	if len(teams) != 4 {
		// Fall back to round-robin for other team counts
		return f.GenerateRoundRobinFixtures(teams)
	}

	fixtures := make([][]*models.Match, 0)

	// Week 1: Team 0 vs 1, Team 2 vs 3
	fixtures = append(fixtures, []*models.Match{
		models.NewMatch(teams[0].ID, teams[1].ID, teams[0].Name, teams[1].Name, 1),
		models.NewMatch(teams[2].ID, teams[3].ID, teams[2].Name, teams[3].Name, 1),
	})

	// Week 2: Team 0 vs 2, Team 1 vs 3
	fixtures = append(fixtures, []*models.Match{
		models.NewMatch(teams[0].ID, teams[2].ID, teams[0].Name, teams[2].Name, 2),
		models.NewMatch(teams[1].ID, teams[3].ID, teams[1].Name, teams[3].Name, 2),
	})

	// Week 3: Team 0 vs 3, Team 1 vs 2
	fixtures = append(fixtures, []*models.Match{
		models.NewMatch(teams[0].ID, teams[3].ID, teams[0].Name, teams[3].Name, 3),
		models.NewMatch(teams[1].ID, teams[2].ID, teams[1].Name, teams[2].Name, 3),
	})

	// Week 4: Return matches - Team 1 vs 0, Team 3 vs 2
	fixtures = append(fixtures, []*models.Match{
		models.NewMatch(teams[1].ID, teams[0].ID, teams[1].Name, teams[0].Name, 4),
		models.NewMatch(teams[3].ID, teams[2].ID, teams[3].Name, teams[2].Name, 4),
	})

	// Week 5: Return matches - Team 2 vs 0, Team 3 vs 1
	fixtures = append(fixtures, []*models.Match{
		models.NewMatch(teams[2].ID, teams[0].ID, teams[2].Name, teams[0].Name, 5),
		models.NewMatch(teams[3].ID, teams[1].ID, teams[3].Name, teams[1].Name, 5),
	})

	// Week 6: Return matches - Team 3 vs 0, Team 2 vs 1
	fixtures = append(fixtures, []*models.Match{
		models.NewMatch(teams[3].ID, teams[0].ID, teams[3].Name, teams[0].Name, 6),
		models.NewMatch(teams[2].ID, teams[1].ID, teams[2].Name, teams[1].Name, 6),
	})

	return fixtures
}
