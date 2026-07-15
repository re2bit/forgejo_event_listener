package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init initializes the global Zerolog logger
func Init() {
	// Use console output for readability during development
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()

	// Use info level unless configured otherwise
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
