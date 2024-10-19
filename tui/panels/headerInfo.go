package panels

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/lucasvillarinho/plumber/tui"
)

func CreateHeaderInfoPanel(cfg tui.Config) *tview.Flex {
	headerInfoPanel := tview.NewFlex()
	metaInfoArea := tview.NewTextView().SetDynamicColors(true).SetRegions(true)
	metaInfoArea.
		SetBorder(true).
		SetBorderColor(cfg.Theme.BorderHeaderInfoColor).
		SetTitle("🛁[::b] Plumber").
		SetTitleColor(tcell.Color(cfg.Theme.TextPrimaryColor)).
		SetTitleAlign(tview.AlignLeft)

	infoText := "Welcome to Plumber!\n" +
		"Version: 1.0.0\n" +
		"Go Version: 1.22\n" +
		"Status: Running\n" +
		"Users Connected: 5\n"

	metaInfoArea.SetText(infoText)
	metaInfoArea.SetBackgroundColor(cfg.Theme.BackgroundColor)

	headerInfoPanel.AddItem(metaInfoArea, 0, 1, false)

	return headerInfoPanel
}
