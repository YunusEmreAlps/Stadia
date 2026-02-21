package services

import (
	"stadia-backend/models"
	"testing"
)

func TestGenerateRoundRobinFixtures(t *testing.T) {
	service := NewFixtureService()

	teams := []*models.Team{
		models.NewTeam("Team A", 80, ""),
		models.NewTeam("Team B", 75, ""),
		models.NewTeam("Team C", 70, ""),
		models.NewTeam("Team D", 65, ""),
	}

	fixtures := service.GenerateRoundRobinFixtures(teams)

	// 4 teams playing home and away: 4C2 * 2 = 12 matches total
	expectedMatches := 12
	actualMatches := 0
	for _, weekMatches := range fixtures {
		actualMatches += len(weekMatches)
	}

	if actualMatches != expectedMatches {
		t.Errorf("Expected %d matches, got %d", expectedMatches, actualMatches)
	}

	// Verify each team plays each other team twice
	matchCount := make(map[string]map[string]int)
	for _, team := range teams {
		matchCount[team.ID] = make(map[string]int)
	}

	for _, weekMatches := range fixtures {
		for _, match := range weekMatches {
			matchCount[match.HomeTeamID][match.AwayTeamID]++
			matchCount[match.AwayTeamID][match.HomeTeamID]++
		}
	}

	for i, team1 := range teams {
		for j, team2 := range teams {
			if i != j {
				count := matchCount[team1.ID][team2.ID]
				if count != 2 {
					t.Errorf("Team %s should play Team %s twice, but played %d times",
						team1.Name, team2.Name, count)
				}
			}
		}
	}
}

func TestGenerateFixturesOptimized(t *testing.T) {
	service := NewFixtureService()

	teams := []*models.Team{
		models.NewTeam("Team A", 80, ""),
		models.NewTeam("Team B", 75, ""),
		models.NewTeam("Team C", 70, ""),
		models.NewTeam("Team D", 65, ""),
	}

	fixtures := service.GenerateFixturesOptimized(teams)

	// Should have 6 weeks
	if len(fixtures) != 6 {
		t.Errorf("Expected 6 weeks, got %d", len(fixtures))
	}

	// Count total matches
	totalMatches := 0
	for _, weekMatches := range fixtures {
		totalMatches += len(weekMatches)

		// Each week should have 2 matches for 4 teams
		if len(weekMatches) != 2 {
			t.Errorf("Expected 2 matches per week, got %d", len(weekMatches))
		}
	}

	// Total should be 12 matches
	if totalMatches != 12 {
		t.Errorf("Expected 12 total matches, got %d", totalMatches)
	}

	// Verify week numbers are correct
	for weekIdx, weekMatches := range fixtures {
		expectedWeek := weekIdx + 1
		for _, match := range weekMatches {
			if match.Week != expectedWeek {
				t.Errorf("Match in week %d has incorrect week number: %d",
					expectedWeek, match.Week)
			}
		}
	}
}

func TestGenerateFixturesWithTwoTeams(t *testing.T) {
	service := NewFixtureService()

	teams := []*models.Team{
		models.NewTeam("Team A", 80, ""),
		models.NewTeam("Team B", 75, ""),
	}

	fixtures := service.GenerateRoundRobinFixtures(teams)

	// 2 teams playing home and away: 2 matches total
	totalMatches := 0
	for _, weekMatches := range fixtures {
		totalMatches += len(weekMatches)
	}

	if totalMatches != 2 {
		t.Errorf("Expected 2 matches for 2 teams, got %d", totalMatches)
	}
}

func TestGenerateFixturesEmpty(t *testing.T) {
	service := NewFixtureService()

	teams := []*models.Team{}
	fixtures := service.GenerateRoundRobinFixtures(teams)

	if fixtures != nil {
		t.Errorf("Expected nil fixtures for empty teams, got %v", fixtures)
	}
}
