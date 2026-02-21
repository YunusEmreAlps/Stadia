package models

// Prediction represents the championship probability for a team
type Prediction struct {
	TeamID      string  `json:"teamId"`
	TeamName    string  `json:"teamName"`
	Probability float64 `json:"probability"` // Percentage (0-100)
}

// PredictionResponse contains predictions for all teams
type PredictionResponse struct {
	Week        int          `json:"week"`
	Predictions []Prediction `json:"predictions"`
}
