package main

import (
	"os"

	"github.com/rs/zerolog"
)

// MustNewLogger returns a logger for the app
func MustNewLogger() (zerolog.Logger, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return logger, nil
}
