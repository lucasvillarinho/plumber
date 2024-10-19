package main

import (
	"log"

	"github.com/rivo/tview"

	"github.com/lucasvillarinho/plumber/tui"
	"github.com/lucasvillarinho/plumber/tui/panels"
	"github.com/lucasvillarinho/plumber/tui/themes"
)

type App struct {
	headeinfoPanel *tview.Flex
	outputPanel    *tview.List

	layout *tview.Flex
	pages  *tview.Pages
	app    *tview.Application
	config *tui.Config
}

func NewApp() *App {
	ui := &App{
		headeinfoPanel: tview.NewFlex(),
		outputPanel:    tview.NewList(),

		layout: tview.NewFlex(),
		app:    tview.NewApplication(),
	}

	theme, err := themes.GetTheme("default")
	if err != nil {
		log.Fatal(err)
	}

	ui.config = &tui.Config{
		Theme: *theme,
	}

	ui.headeinfoPanel = panels.CreateHeaderInfoPanel(*ui.config)
	ui.outputPanel = panels.CreateOutputPanel(*ui.config)

	ui.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(ui.headeinfoPanel, 0, 2, false).
		AddItem(ui.outputPanel, 0, 10, false)

	ui.app.SetRoot(ui.layout, true)

	return ui
}

func (a *App) Run() error {

	a.pages = tview.NewPages()
	a.pages.AddPage("base", a.layout, true, true)

	return a.app.SetRoot(a.pages, true).Run()
}

func main() {
	app := NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
