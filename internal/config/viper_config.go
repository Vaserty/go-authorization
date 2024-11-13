package config

import (
	"github.com/Vaserty/go-authorization/internal/customtypes"
	"github.com/spf13/viper"
)

func viperSetConfig(environment customtypes.Environment) {
	viper.AddConfigPath(string(SettingsFolderPath))
	viper.SetConfigType("yaml")
	viper.SetConfigName(string(environment))
}

func viperLoadConfig(environment customtypes.Environment) (*Config, error) {
	viperSetConfig(environment)

	config := &Config{}

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return config, err
	}

	return config, nil
}
