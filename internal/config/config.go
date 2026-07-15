package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Load reads environment variables from a .env file
func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("No .env file found or failed to load it")
	} else {
		log.Info().Msg(".env file loaded successfully")
	}
}

// GetEnv returns an environment variable or a fallback value
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
