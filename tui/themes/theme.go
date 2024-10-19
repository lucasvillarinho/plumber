package themes

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/spf13/viper"

	"github.com/lucasvillarinho/plumber/helpers"
)

type ThemeFile struct {
	BackgroundColor string `mapstructure:"background_color"`

	BorderHeaderInfoColor string `mapstructure:"border_header_info_color"`
	BorderOutputColor     string `mapstructure:"border_output_color"`

	TextPrimaryColor string `mapstructure:"text_primary_color"`
}

type Theme struct {
	BackgroundColor tcell.Color

	BorderHeaderInfoColor tcell.Color
	BorderOutputColor     tcell.Color

	TextLogoPlumberColor tcell.Color
	TextPrimaryColor     tcell.Color
}

func GetTheme(tn string) (*Theme, error) {
	viper.SetConfigName(tn)
	viper.SetConfigType("toml")
	viper.AddConfigPath("./tui/themes")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var theme ThemeFile
	if err := viper.Unmarshal(&theme); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return parseThemeToTcellColors(theme)
}

func parseThemeToTcellColors(th ThemeFile) (*Theme, error) {
	bgColor, err := helpers.ParseHexToTcellColor(th.BackgroundColor)
	if err != nil {
		return nil, fmt.Errorf("invalid background color: %w", err)
	}

	// Border colors
	borderHeaderInfoColor, err := helpers.ParseHexToTcellColor(th.BorderHeaderInfoColor)
	if err != nil {
		return nil, fmt.Errorf("invalid header border color: %w", err)
	}

	borderOutputColor, err := helpers.ParseHexToTcellColor(th.BorderOutputColor)
	if err != nil {
		return nil, fmt.Errorf("invalid output border color: %w", err)
	}

	// Text colors
	textPrimaryColor, err := helpers.ParseHexToTcellColor(th.TextPrimaryColor)
	if err != nil {
		return nil, fmt.Errorf("invalid text primary color: %w", err)
	}

	return &Theme{
		BackgroundColor: bgColor,

		BorderHeaderInfoColor: borderHeaderInfoColor,
		BorderOutputColor:     borderOutputColor,

		TextPrimaryColor: textPrimaryColor,
	}, nil
}
