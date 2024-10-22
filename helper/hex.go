package helpers

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

func ParseHexToTcellColor(hex string) (tcell.Color, error) {
	// Remove o '#' se estiver presente
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	colorValue, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return tcell.ColorDefault, fmt.Errorf("failed to parse hex color: %w", err)
	}

	return tcell.NewHexColor(int32(colorValue)), nil
}
