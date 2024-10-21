package component

import (
	"fmt"

	"github.com/rivo/tview"

	cfg "github.com/lucasvillarinho/plumber/config"
	ijt "github.com/lucasvillarinho/plumber/internal/injector"
)

type OutputComponent struct {
	theme *cfg.Theme
}

func NewOutputComponent(injector *ijt.Injector) (*OutputComponent, error) {
	theme, err := ijt.Get[*cfg.Theme](injector)
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
