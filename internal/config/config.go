package config

import (
	"github.com/Vaserty/go-authorization/internal/customtypes"
	"github.com/go-playground/validator/v10"
)

const SettingsFolderPath string = "settings"

type Config struct {
	AppInfo struct {
		Environment customtypes.Environment `validate:"required,oneof=development test production"`
		Name        string                  `validate:"required"`
		Version     string                  `validate:"required"`
	}
	Logger struct {
		Level string `validate:"required,oneof=debug info warn error"`
	}
}

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
