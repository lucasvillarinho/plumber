package component

import (
	"fmt"

	"github.com/rivo/tview"

	cfg "github.com/lucasvillarinho/plumber/config"
	di "github.com/lucasvillarinho/plumber/internal/injector"
	psb "github.com/lucasvillarinho/plumber/internal/pubsub"
)

type OutputComponent struct {
	theme       *cfg.ThemeConfig
	pubsub      *psb.PubSub[string]
	application tview.Application
}

func NewOutputComponent(injector *di.Injector) (*OutputComponent, error) {
	theme, err := di.Get[*cfg.ThemeConfig](injector)
	if err != nil || theme == nil {
		return nil, fmt.Errorf("failed to inject Theme instance: %w", err)
	}

	pubsub, err := di.Get[*psb.PubSub[string]](injector)
	if err != nil || pubsub == nil {
		return nil, fmt.Errorf("failed to inject PubSub instance: %w", err)
	}

	return &OutputComponent{
		theme:  *theme,
		pubsub: *pubsub,
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
