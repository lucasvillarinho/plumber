package component

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	cfg "github.com/lucasvillarinho/plumber/config"
	inj "github.com/lucasvillarinho/plumber/internal/injector"
)

type HeaderComponent struct {
	theme *cfg.ThemeConfig
}

func NewHeaderComponent(injector *inj.Injector) (*HeaderComponent, error) {
	theme, err := inj.Get[*cfg.ThemeConfig](injector)
	if err != nil || theme == nil {
		return nil, fmt.Errorf("failed to inject Theme instance: %w", err)
	}

	return &HeaderComponent{
		theme: *theme,
	}, nil
}

func (hdc *HeaderComponent) CreateHeaderPanel() *tview.Flex {
	headerInfoPanel := tview.NewFlex()
	metaInfoArea := tview.
		NewTextView().
		SetDynamicColors(true).
		SetRegions(true)
	metaInfoArea.
		SetBorder(true).
		SetBorderColor(hdc.theme.BorderHeaderInfoColor).
		SetTitle("🛁[::b] Plumber").
		SetTitleColor(tcell.Color(hdc.theme.TextPrimaryColor)).
		SetTitleAlign(tview.AlignLeft)

	infoText := "Welcome to Plumber!\n" +
		"Version: 1.0.0\n" +
		"Go Version: 1.22\n" +
		"Status: Running\n" +
		"Users Connected: 5\n"

	metaInfoArea.SetText(infoText)
	metaInfoArea.SetBackgroundColor(hdc.theme.BackgroundColor)

	headerInfoPanel.AddItem(metaInfoArea, 0, 1, false)

	return headerInfoPanel
}
