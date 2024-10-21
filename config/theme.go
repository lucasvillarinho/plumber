package configs

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/spf13/viper"

	"github.com/lucasvillarinho/plumber/helpers"
	inj "github.com/lucasvillarinho/plumber/internal/injector"
)

type ThemeFile struct {
	BackgroundColor string `mapstructure:"backgroundColor" yaml:"backgroundColor"`

	Border struct {
		HeaderInfoColor string `mapstructure:"headerInfoColor" yaml:"headerInfoColor"`
		OutputColor     string `mapstructure:"outputColor" yaml:"outputColor"`
	} `mapstructure:"border" yaml:"border"`

	Text struct {
		PrimaryColor string `mapstructure:"primaryColor" yaml:"primaryColor"`
	} `mapstructure:"text" yaml:"text"`
}

type ThemeConfig struct {
	BackgroundColor tcell.Color

	BorderHeaderInfoColor tcell.Color
	BorderOutputColor     tcell.Color

	TextLogoPlumberColor tcell.Color
	TextPrimaryColor     tcell.Color
}

func NewThemeConfig(injector *inj.Injector) (*ThemeConfig, error) {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./app/themes")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error to read config: %w", err)
	}

	var theme ThemeFile
	if err := viper.Unmarshal(&theme); err != nil {
		return nil, fmt.Errorf("error to unmarshal config: %w", err)
	}

	return parseThemeToTcellColors(theme)
}

func parseThemeToTcellColors(th ThemeFile) (*ThemeConfig, error) {
	bgColor, err := helpers.ParseHexToTcellColor(th.BackgroundColor)
	if err != nil {
		return nil, fmt.Errorf("invalid background color: %w", err)
	}

	borderHeaderInfoColor, err := helpers.ParseHexToTcellColor(th.Border.HeaderInfoColor)
	if err != nil {
		return nil, fmt.Errorf("invalid header border color: %w", err)
	}

	borderOutputColor, err := helpers.ParseHexToTcellColor(th.Border.OutputColor)
	if err != nil {
		return nil, fmt.Errorf("invalid output border color: %w", err)
	}

	textPrimaryColor, err := helpers.ParseHexToTcellColor(th.Text.PrimaryColor)
	if err != nil {
		return nil, fmt.Errorf("invalid text primary color: %w", err)
	}

	return &ThemeConfig{
		BackgroundColor:       bgColor,
		BorderHeaderInfoColor: borderHeaderInfoColor,
		BorderOutputColor:     borderOutputColor,
		TextPrimaryColor:      textPrimaryColor,
	}, nil
}
