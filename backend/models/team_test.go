package models

import "testing"

func TestNewTeam(t *testing.T) {
	team := NewTeam("Test Team", 75, "logo.png")

	if team.Name != "Test Team" {
		t.Errorf("Expected name 'Test Team', got '%s'", team.Name)
	}

	if team.Power != 75 {
		t.Errorf("Expected power 75, got %d", team.Power)
	}

	if team.ID == "" {
		t.Error("Team ID should not be empty")
	}

	if team.Played != 0 || team.Won != 0 || team.Points != 0 {
		t.Error("Initial stats should be zero")
	}
}

func TestUpdateStats_Win(t *testing.T) {
	team := NewTeam("Test Team", 75, "")

	// Team wins 3-1
	team.UpdateStats(3, 1)

	if team.Played != 1 {
		t.Errorf("Expected 1 game played, got %d", team.Played)
	}

	if team.Won != 1 {
		t.Errorf("Expected 1 win, got %d", team.Won)
	}

	if team.Points != 3 {
		t.Errorf("Expected 3 points, got %d", team.Points)
	}

	if team.GoalsFor != 3 {
		t.Errorf("Expected 3 goals for, got %d", team.GoalsFor)
	}

	if team.GoalsAgainst != 1 {
		t.Errorf("Expected 1 goal against, got %d", team.GoalsAgainst)
	}
}

func TestUpdateStats_Draw(t *testing.T) {
	team := NewTeam("Test Team", 75, "")

	// Team draws 2-2
	team.UpdateStats(2, 2)

	if team.Drawn != 1 {
		t.Errorf("Expected 1 draw, got %d", team.Drawn)
	}

	if team.Points != 1 {
		t.Errorf("Expected 1 point, got %d", team.Points)
	}
}

func TestUpdateStats_Loss(t *testing.T) {
	team := NewTeam("Test Team", 75, "")

	// Team loses 1-3
	team.UpdateStats(1, 3)

	if team.Lost != 1 {
		t.Errorf("Expected 1 loss, got %d", team.Lost)
	}

	if team.Points != 0 {
		t.Errorf("Expected 0 points, got %d", team.Points)
	}
}

func TestGoalDifference(t *testing.T) {
	team := NewTeam("Test Team", 75, "")

	team.UpdateStats(5, 2) // Win 5-2
	team.UpdateStats(1, 3) // Loss 1-3

	expectedGD := 1 // (5-2) + (1-3) = 3 - 4 = -1... wait: (5+1) - (2+3) = 6 - 5 = 1
	actualGD := team.GoalDifference()

	if actualGD != expectedGD {
		t.Errorf("Expected goal difference %d, got %d", expectedGD, actualGD)
	}
}

func TestResetStats(t *testing.T) {
	team := NewTeam("Test Team", 75, "")

	// Play some matches
	team.UpdateStats(3, 1)
	team.UpdateStats(2, 2)
	team.UpdateStats(1, 3)

	// Reset
	team.ResetStats()

	if team.Played != 0 || team.Won != 0 || team.Drawn != 0 || team.Lost != 0 {
		t.Error("All match stats should be zero after reset")
	}

	if team.Points != 0 || team.GoalsFor != 0 || team.GoalsAgainst != 0 {
		t.Error("All points and goals should be zero after reset")
	}

	// Name and power should remain
	if team.Name != "Test Team" || team.Power != 75 {
		t.Error("Name and power should not change after reset")
	}
}
