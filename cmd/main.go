package main

import (
	"fmt"

	"github.com/Vaserty/go-authorization/internal/config"
	"github.com/Vaserty/go-authorization/internal/customtypes"
	"github.com/samber/do/v2"
	logger "github.com/sirupsen/logrus"
)

func initializeConfig(environment customtypes.Environment) (*config.Config, error) {
	logger.Info("Starting to load configuration...")
	buildConfig, err := config.NewConfig(environment)

	if err != nil {
		return nil, fmt.Errorf("error creating configuration: %w", err)
	}

	if validationErr := buildConfig.Validate(); validationErr != nil {
		return nil, fmt.Errorf("configuration validation error: %w", validationErr)
	}

	logger.Info("Configuration successfully loaded.")
	return buildConfig, nil
}

func provideConfig(
	injector *do.RootScope,
	environment customtypes.Environment,
) (*config.Config, error) {
	do.Provide(injector, func(do.Injector) (*config.Config, error) {
		return initializeConfig(environment)
	})
	return do.Invoke[*config.Config](injector)
}

func main() {
	logger.Info("Initializing application...")
	injector := do.New()
	environment := config.GetEnvironment()
	logger.Infof("Running in environment: '%v'", environment)
	buildedConfig, err := provideConfig(injector, environment)
	if err != nil {
		logger.Fatalf("Failed to initialize configuration: %v", err)
	}

	if err := config.ConfigLogger(buildedConfig); err != nil {
		logger.Fatalf("Failed to configure logger: %v", err)
	}

	logger.Infof(
		"Application name: '%v', Application version: '%v'",
		buildedConfig.AppInfo.Name,
		buildedConfig.AppInfo.Version,
	)
}
