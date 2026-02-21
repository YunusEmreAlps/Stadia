package models

import "github.com/google/uuid"

// Team represents a football team in the league
type Team struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Power        int    `json:"power"`          // Team strength (0-100)
	Played       int    `json:"played"`         // Matches played
	Won          int    `json:"won"`            // Matches won
	Drawn        int    `json:"drawn"`          // Matches drawn
	Lost         int    `json:"lost"`           // Matches lost
	GoalsFor     int    `json:"goalsFor"`       // Goals scored
	GoalsAgainst int    `json:"goalsAgainst"`   // Goals conceded
	Points       int    `json:"points"`         // Total points
	Logo         string `json:"logo,omitempty"` // Team logo URL
}

// NewTeam creates a new team with a unique ID
func NewTeam(name string, power int, logo string) *Team {
	return &Team{
		ID:           uuid.New().String(),
		Name:         name,
		Power:        power,
		Played:       0,
		Won:          0,
		Drawn:        0,
		Lost:         0,
		GoalsFor:     0,
		GoalsAgainst: 0,
		Points:       0,
		Logo:         logo,
	}
}

// GoalDifference calculates the goal difference
func (t *Team) GoalDifference() int {
	return t.GoalsFor - t.GoalsAgainst
}

// UpdateStats updates team statistics after a match
func (t *Team) UpdateStats(goalsFor, goalsAgainst int) {
	t.Played++
	t.GoalsFor += goalsFor
	t.GoalsAgainst += goalsAgainst

	if goalsFor > goalsAgainst {
		t.Won++
		t.Points += 3
	} else if goalsFor == goalsAgainst {
		t.Drawn++
		t.Points += 1
	} else {
		t.Lost++
	}
}

// ResetStats resets all team statistics
func (t *Team) ResetStats() {
	t.Played = 0
	t.Won = 0
	t.Drawn = 0
	t.Lost = 0
	t.GoalsFor = 0
	t.GoalsAgainst = 0
	t.Points = 0
}
