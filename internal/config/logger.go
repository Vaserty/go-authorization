package config

import (
	"fmt"

	logger "github.com/sirupsen/logrus"
)

func ConfigLogger(config *Config) error {
	level, err := logger.ParseLevel(config.Logger.Level)
	if err != nil {
		return fmt.Errorf("error parsing log level: %w", err)
	}
	logger.SetLevel(level)
	logger.Info("Logger configured successfully.")
	return nil
}
