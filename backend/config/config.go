package config

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	App App `mapstructure:"app"`
	DB  DB  `mapstructure:"db"`
}

// App contains application-specific configuration
type App struct {
	Mode           string   `mapstructure:"mode"`
	Port           string   `mapstructure:"port"`
	Version        string   `mapstructure:"version"`
	Name           string   `mapstructure:"name"`
	AllowedOrigins []string `mapstructure:"allowed_origins"`
}

// DB contains database configuration
type DB struct {
	URL string `mapstructure:"url"`
}

// AppConfig is the global configuration instance
var AppConfig Config

// Load reads and parses the configuration file
func Load(processCwdir string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(processCwdir, "config"))
	viper.AddConfigPath(".") // Also check current directory

	// Set defaults
	viper.SetDefault("app.mode", "debug")
	viper.SetDefault("app.port", "8000")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.name", "stadia-backend")
	viper.SetDefault("app.allowed_origins", []string{"http://localhost", "http://localhost:8080", "http://localhost:5173", "http://localhost:3000", "https://stadiaa.netlify.app", "https://stadia-xex6.onrender.com"})

	// Read environment variables
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, using defaults and environment variables")
		} else {
			return fmt.Errorf("error reading config file: %w", err)
		}
	} else {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}

	// Unmarshal config
	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("error unmarshaling config: %w", err)
	}

	return nil
}

// MustLoad loads configuration and panics on error
func MustLoad(processCwdir string) {
	if err := Load(processCwdir); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
}

// GetPort returns the configured port or default
func GetPort() string {
	if AppConfig.App.Port != "" {
		return AppConfig.App.Port
	}
	return "8000"
}
