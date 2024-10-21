package configs

import (
	"fmt"

	"github.com/spf13/viper"

	inj "github.com/lucasvillarinho/plumber/internal/injector"
)

type Process struct {
	Name string `mapstructure:"name" yaml:"name"`
	Cmd  string `mapstructure:"cmd" yaml:"cmd"`
}

type PlumberConfig struct {
	Process []Process `mapstructure:"process" yaml:"process"`
}

func NewPlumberConfig(injector *inj.Injector) (*PlumberConfig, error) {
	viper.SetConfigName("plumber")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error to read config: %w", err)
	}

	var pumberConfig PlumberConfig
	if err := viper.Unmarshal(&pumberConfig); err != nil {
		return nil, fmt.Errorf("error to unmarshal config: %w", err)
	}

	return &pumberConfig, nil
}
