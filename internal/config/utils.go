package config

import (
	"os"

	"github.com/Vaserty/go-authorization/internal/customtypes"
)

func GetEnvironment() customtypes.Environment {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		return customtypes.EnvDev
	}
	return customtypes.Environment(environment)
}

func NewConfig(environment customtypes.Environment) (*Config, error) {
	return viperLoadConfig(environment)
}
