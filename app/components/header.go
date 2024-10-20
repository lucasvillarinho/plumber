package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/lucasvillarinho/plumber/app/themes"
	pkg "github.com/lucasvillarinho/plumber/pkg/injector"
)

type HeaderComponent struct {
	theme *themes.Theme
}

func NewHeaderComponent(injector *pkg.Injector) (*HeaderComponent, error) {
	theme, err := pkg.Get[*themes.Theme](injector)
	if err != nil || theme == nil {
		return nil, fmt.Errorf("failed to inject Theme instance: %w", err)
	}

	return &HeaderComponent{
		theme: *theme,
	}, nil
}

func (hec *HeaderComponent) CreateHeaderPanel() *tview.Flex {
	headerInfoPanel := tview.NewFlex()
	metaInfoArea := tview.
		NewTextView().
		SetDynamicColors(true).
		SetRegions(true)
	metaInfoArea.
		SetBorder(true).
		SetBorderColor(hec.theme.BorderHeaderInfoColor).
		SetTitle("🛁[::b] Plumber").
		SetTitleColor(tcell.Color(hec.theme.TextPrimaryColor)).
		SetTitleAlign(tview.AlignLeft)

	infoText := "Welcome to Plumber!\n" +
		"Version: 1.0.0\n" +
		"Go Version: 1.22\n" +
		"Status: Running\n" +
		"Users Connected: 5\n"

	metaInfoArea.SetText(infoText)
	metaInfoArea.SetBackgroundColor(hec.theme.BackgroundColor)

	headerInfoPanel.AddItem(metaInfoArea, 0, 1, false)

	return headerInfoPanel
}
