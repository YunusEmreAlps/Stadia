package handlers

import (
	"net/http"
	"stadia-backend/models"
	"stadia-backend/services"

	"github.com/gin-gonic/gin"
)

// LeagueHandler handles league-related HTTP requests
type LeagueHandler struct {
	leagueService *services.LeagueService
}

// NewLeagueHandler creates a new league handler
func NewLeagueHandler(leagueService *services.LeagueService) *LeagueHandler {
	return &LeagueHandler{
		leagueService: leagueService,
	}
}

// InitializeRequest represents the request to initialize a league
type InitializeRequest struct {
	Teams []struct {
		Name  string `json:"name" binding:"required"`
		Power int    `json:"power" binding:"required,min=1,max=100"`
		Logo  string `json:"logo"`
	} `json:"teams" binding:"required,min=2"`
}

// UpdateMatchRequest represents the request to update a match result
type UpdateMatchRequest struct {
	HomeScore int `json:"homeScore" binding:"min=0"`
	AwayScore int `json:"awayScore" binding:"min=0"`
}

// Initialize initializes the league with teams
// @Summary Initialize league
// @Description Initialize a new league with the provided teams
// @Tags league
// @Accept json
// @Produce json
// @Param request body InitializeRequest true "Teams to initialize"
// @Success 200 {object} map[string]interface{} "League initialized successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /league/initialize [post]
func (h *LeagueHandler) Initialize(c *gin.Context) {
	var req InitializeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teams := make([]*models.Team, len(req.Teams))
	for i, teamReq := range req.Teams {
		teams[i] = models.NewTeam(teamReq.Name, teamReq.Power, teamReq.Logo)
	}

	err := h.leagueService.InitializeLeague(teams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "League initialized successfully",
		"league":  h.leagueService.GetLeague(),
	})
}

// GetLeague returns the current league state
// @Summary Get league
// @Description Get the current state of the league including all teams and matches
// @Tags league
// @Produce json
// @Success 200 {object} models.League "Current league state"
// @Router /league [get]
func (h *LeagueHandler) GetLeague(c *gin.Context) {
	league := h.leagueService.GetLeague()
	c.JSON(http.StatusOK, league)
}

// GetStandings returns the current league standings
// @Summary Get standings
// @Description Get the current league standings sorted by points
// @Tags league
// @Produce json
// @Success 200 {object} map[string]interface{} "League standings"
// @Router /league/standings [get]
func (h *LeagueHandler) GetStandings(c *gin.Context) {
	standings := h.leagueService.GetStandings()
	c.JSON(http.StatusOK, gin.H{"standings": standings})
}

// PlayNextWeek simulates the next week of matches
// @Summary Play next week
// @Description Simulate the next week of matches in the league
// @Tags league
// @Produce json
// @Success 200 {object} map[string]interface{} "Week played successfully"
// @Failure 400 {object} map[string]string "All weeks already played or league not initialized"
// @Router /league/play-next-week [post]
func (h *LeagueHandler) PlayNextWeek(c *gin.Context) {
	err := h.leagueService.PlayNextWeek()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Week played successfully",
		"league":  h.leagueService.GetLeague(),
	})
}

// PlayAllWeeks simulates all remaining weeks
// @Summary Play all weeks
// @Description Simulate all remaining weeks of matches in the league
// @Tags league
// @Produce json
// @Success 200 {object} map[string]interface{} "All weeks played successfully"
// @Failure 400 {object} map[string]string "League not initialized"
// @Router /league/play-all-weeks [post]
func (h *LeagueHandler) PlayAllWeeks(c *gin.Context) {
	err := h.leagueService.PlayAllWeeks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All weeks played successfully",
		"league":  h.leagueService.GetLeague(),
	})
}

// UpdateMatch updates a match result
// @Summary Update match
// @Description Update the result of a specific match
// @Tags league
// @Accept json
// @Produce json
// @Param id path string true "Match ID"
// @Param request body UpdateMatchRequest true "New match scores"
// @Success 200 {object} map[string]interface{} "Match updated successfully"
// @Failure 400 {object} map[string]string "Invalid request or match not found"
// @Router /league/match/{id} [put]
func (h *LeagueHandler) UpdateMatch(c *gin.Context) {
	matchID := c.Param("id")

	var req UpdateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.leagueService.UpdateMatchResult(matchID, req.HomeScore, req.AwayScore)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Match updated successfully",
		"league":  h.leagueService.GetLeague(),
	})
}

// ResetLeague resets the league
// @Summary Reset league
// @Description Reset the league to its initial state
// @Tags league
// @Produce json
// @Success 200 {object} map[string]interface{} "League reset successfully"
// @Failure 400 {object} map[string]string "League not initialized"
// @Router /league/reset [post]
func (h *LeagueHandler) ResetLeague(c *gin.Context) {
	err := h.leagueService.ResetLeague()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "League reset successfully",
		"league":  h.leagueService.GetLeague(),
	})
}

// GetPredictions returns championship predictions
// @Summary Get predictions
// @Description Get championship predictions using Monte Carlo simulation
// @Tags league
// @Produce json
// @Success 200 {object} []models.Prediction "Championship predictions"
// @Router /league/predictions [get]
func (h *LeagueHandler) GetPredictions(c *gin.Context) {
	predictions := h.leagueService.GetPredictions()
	c.JSON(http.StatusOK, predictions)
}
