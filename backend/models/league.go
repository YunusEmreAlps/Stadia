package models

// League represents the football league
type League struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Teams       map[string]*Team   `json:"teams"`
	Fixtures    [][]*Match         `json:"fixtures"` // Fixtures grouped by week
	CurrentWeek int                `json:"currentWeek"`
	TotalWeeks  int                `json:"totalWeeks"`
	Predictions map[string]float64 `json:"predictions,omitempty"` // Team ID -> Win probability
}

// NewLeague creates a new league
func NewLeague(name string) *League {
	return &League{
		ID:          "champions-league-2026",
		Name:        name,
		Teams:       make(map[string]*Team),
		Fixtures:    make([][]*Match, 0),
		CurrentWeek: 0,
		TotalWeeks:  0,
		Predictions: make(map[string]float64),
	}
}

// AddTeam adds a team to the league
func (l *League) AddTeam(team *Team) {
	l.Teams[team.ID] = team
}

// GetTeam retrieves a team by ID
func (l *League) GetTeam(id string) *Team {
	return l.Teams[id]
}

// GetTeamsList returns a slice of all teams
func (l *League) GetTeamsList() []*Team {
	teams := make([]*Team, 0, len(l.Teams))
	for _, team := range l.Teams {
		teams = append(teams, team)
	}
	return teams
}

// GetMatchesByWeek returns matches for a specific week
func (l *League) GetMatchesByWeek(week int) []*Match {
	if week < 1 || week > len(l.Fixtures) {
		return nil
	}
	return l.Fixtures[week-1]
}

// GetAllMatches returns all matches flattened
func (l *League) GetAllMatches() []*Match {
	allMatches := make([]*Match, 0)
	for _, weekMatches := range l.Fixtures {
		allMatches = append(allMatches, weekMatches...)
	}
	return allMatches
}

// IsFinished checks if all matches have been played
func (l *League) IsFinished() bool {
	return l.CurrentWeek >= l.TotalWeeks
}
