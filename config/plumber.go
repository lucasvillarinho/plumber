package configs

import (
	"fmt"

	"github.com/spf13/viper"

	di "github.com/lucasvillarinho/plumber/internal/injector"
)

type Process struct {
	Name string `mapstructure:"name"`
	Cmd  string `mapstructure:"cmd"`
}

type PlumberConfig struct {
	Process []Process `mapstructure:"process"`
}

func NewPlumberConfig(injector *di.Injector) (*PlumberConfig, error) {
	viper.SetConfigName("plumber")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error to read config: %w", err)
	}

	var pumberConfig PlumberConfig
	if err := viper.Unmarshal(&pumberConfig); err != nil {
		return nil, fmt.Errorf("error to unmarshal config: %w", err)
	}

	return &pumberConfig, nil
}
