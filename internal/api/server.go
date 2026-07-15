package api

import (
	//	"dynamic_runner_subsystem/internal/client"
	//	"net/http"

	"dynamic_runner_subsystem/internal/config"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// StartServer initializes and starts the Gin server
func StartServer() {
	// Enable Gin release mode when requested
	if config.GetEnv("GIN_MODE", "debug") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Standard middleware
	r.Use(gin.Recovery())

	// Request logging middleware
	r.Use(func(c *gin.Context) {
		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Msg("API Request")
		c.Next()
	})

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Debug API route
	r.POST("/api/*param", func(c *gin.Context) {
		param := c.Param("param")
		log.Info().Msgf("API Request: %s", param)
		json, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Error().Msg("Failed to read request body")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read request body",
			})
			return
		}

		log.Info().Msgf("API Request Body: %s", json)
		c.JSON(http.StatusOK, gin.H{
			"Status": "OK",
		})
	})

	port := config.GetEnv("PORT", "8080")
	log.Info().Msgf("Server starts on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
