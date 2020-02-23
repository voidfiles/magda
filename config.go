package main

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config is the config for magda
type Config struct {
	GoogleApplicationCredentials string
}

func readConfig() (Config, error) {
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, errors.Wrap(err, "Failed to read config")
		}
	}

	viper.SetDefault("GOOGLE_APPLICATION_CREDENTIALS", "./service_account_key.json")
	config := Config{
		GoogleApplicationCredentials: viper.GetString("GOOGLE_APPLICATION_CREDENTIALS"),
	}

	return config, nil
}
