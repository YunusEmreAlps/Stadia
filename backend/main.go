package main

import (
	"log"
	"stadia-backend/config"
	"stadia-backend/handlers"
	"stadia-backend/services"

	docs "stadia-backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Stadia
// @description: API for simulating and predicting football championship outcomes based on team performance, fixtures, and historical data.
// @version 1.0.0
// @schemes http https

// @contact.name   Yunus Emre Alpu
// @contact.url    https://yunusemrealpu.netlify.app
// @contact.email  YunusAlpu@icloud.com

// @BasePath /api

func main() {
	// Load configuration
	config.MustLoad(".")

	mode := config.AppConfig.App.Mode
	port := config.GetPort()
	log.Printf("Starting in %s mode on port %s", mode, port)

	// Set Gin mode
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if mode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize services
	leagueService := services.NewLeagueService()

	// Initialize handlers
	leagueHandler := handlers.NewLeagueHandler(leagueService)

	// Setup Gin router
	router := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.AppConfig.App.AllowedOrigins
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	log.Printf("CORS enabled for origins: %v", corsConfig.AllowOrigins)
	router.Use(cors.New(corsConfig))

	// API routes
	api := router.Group("/api")
	{
		league := api.Group("/league")
		{
			league.POST("/initialize", leagueHandler.Initialize)
			league.GET("", leagueHandler.GetLeague)
			league.GET("/standings", leagueHandler.GetStandings)
			league.POST("/play-next-week", leagueHandler.PlayNextWeek)
			league.POST("/play-all-weeks", leagueHandler.PlayAllWeeks)
			league.PUT("/match/:id", leagueHandler.UpdateMatch)
			league.POST("/reset", leagueHandler.ResetLeague)
			league.GET("/predictions", leagueHandler.GetPredictions)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": config.AppConfig.App.Name,
		})
	})

	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
