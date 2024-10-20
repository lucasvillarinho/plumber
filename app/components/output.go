package components

import (
	"fmt"

	"github.com/rivo/tview"

	"github.com/lucasvillarinho/plumber/app/themes"
	pkg "github.com/lucasvillarinho/plumber/pkg/injector"
)

type OutputComponent struct {
	theme *themes.Theme
}

func NewOutputComponent(injector *pkg.Injector) (*OutputComponent, error) {
	theme, err := pkg.Get[*themes.Theme](injector)
	if err != nil || theme == nil {
		return nil, fmt.Errorf("failed to inject Theme instance: %w", err)
	}

	return &OutputComponent{
		theme: *theme,
	}, nil
}

func (oc *OutputComponent) CreateOutputPanel() *tview.List {
	outputPanel := tview.NewList()
	outputPanel.
		SetBorder(true).
		SetBorderColor(oc.theme.BorderOutputColor).
		SetTitle("[::b]🔎 Output").
		SetTitleColor(oc.theme.TextPrimaryColor).
		SetTitleAlign(tview.AlignLeft)

	outputPanel.SetBackgroundColor(oc.theme.BackgroundColor)

	return outputPanel
}
