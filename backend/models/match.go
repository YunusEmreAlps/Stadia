package models

import "github.com/google/uuid"

// MatchStatus represents the status of a match
type MatchStatus string

const (
	StatusNotPlayed MatchStatus = "not_played"
	StatusPlayed    MatchStatus = "played"
)

// Match represents a football match between two teams
type Match struct {
	ID           string      `json:"id"`
	HomeTeamID   string      `json:"homeTeamId"`
	AwayTeamID   string      `json:"awayTeamId"`
	HomeTeamName string      `json:"homeTeamName"`
	AwayTeamName string      `json:"awayTeamName"`
	HomeScore    int         `json:"homeScore"`
	AwayScore    int         `json:"awayScore"`
	Week         int         `json:"week"`
	Status       MatchStatus `json:"status"`
}

// NewMatch creates a new match
func NewMatch(homeTeamID, awayTeamID, homeTeamName, awayTeamName string, week int) *Match {
	return &Match{
		ID:           uuid.New().String(),
		HomeTeamID:   homeTeamID,
		AwayTeamID:   awayTeamID,
		HomeTeamName: homeTeamName,
		AwayTeamName: awayTeamName,
		HomeScore:    0,
		AwayScore:    0,
		Week:         week,
		Status:       StatusNotPlayed,
	}
}

// IsPlayed checks if the match has been played
func (m *Match) IsPlayed() bool {
	return m.Status == StatusPlayed
}

// SetResult sets the match result
func (m *Match) SetResult(homeScore, awayScore int) {
	m.HomeScore = homeScore
	m.AwayScore = awayScore
	m.Status = StatusPlayed
}
