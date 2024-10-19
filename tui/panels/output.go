package panels

import (
	"github.com/rivo/tview"

	"github.com/lucasvillarinho/plumber/tui"
)

func CreateOutputPanel(cfg tui.Config) *tview.List {
	outputPanel := tview.NewList()
	outputPanel.
		SetBorder(true).
		SetBorderColor(cfg.Theme.BorderOutputColor).
		SetTitle("[::b]🔎 Output").
		SetTitleColor(cfg.Theme.TextPrimaryColor).
		SetTitleAlign(tview.AlignLeft)

	outputPanel.SetBackgroundColor(cfg.Theme.BackgroundColor)

	return outputPanel
}
