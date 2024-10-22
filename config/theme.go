package configs

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/spf13/viper"

	hlp "github.com/lucasvillarinho/plumber/helper"
	inj "github.com/lucasvillarinho/plumber/internal/injector"
)

type ThemeFile struct {
	Colors struct {
		BackgroundColor string `mapstructure:"backgroundColor"`

		Border struct {
			HeaderInfoColor string `mapstructure:"headerInfoColor"`
			OutputColor     string `mapstructure:"outputColor"`
		} `mapstructure:"border"`

		Text struct {
			PrimaryColor string `mapstructure:"primaryColor"`
		} `mapstructure:"text"`
	} `mapstructure:"colors"`
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
	viper.SetConfigType("toml")
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

func parseThemeToTcellColors(themeFile ThemeFile) (*ThemeConfig, error) {
	bgColor, err := hlp.ParseHexToTcellColor(themeFile.Colors.BackgroundColor)
	if err != nil {
		return nil, fmt.Errorf("invalid background color: %w", err)
	}

	borderHeaderInfoColor, err := hlp.ParseHexToTcellColor(themeFile.Colors.Border.HeaderInfoColor)
	if err != nil {
		return nil, fmt.Errorf("invalid header border color: %w", err)
	}

	borderOutputColor, err := hlp.ParseHexToTcellColor(themeFile.Colors.Border.OutputColor)
	if err != nil {
		return nil, fmt.Errorf("invalid output border color: %w", err)
	}

	textPrimaryColor, err := hlp.ParseHexToTcellColor(themeFile.Colors.Text.PrimaryColor)
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
