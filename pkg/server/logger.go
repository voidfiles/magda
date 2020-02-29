package server

import (
	"os"

	"github.com/rs/zerolog"
)

// MustNewLogger returns a logger for the app
func MustNewLogger() zerolog.Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return logger
}
