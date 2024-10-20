package app

import (
	"fmt"
	"time"

	"github.com/rivo/tview"

	"github.com/lucasvillarinho/plumber/app/components"
	"github.com/lucasvillarinho/plumber/app/themes"
	pkg "github.com/lucasvillarinho/plumber/pkg/injector"
)

type App struct {
	headerPanel     *tview.Flex
	headerComponent *components.HeaderComponent

	outputPanel     *tview.List
	outputComponent *components.OutputComponent

	layout      *tview.Flex
	pages       *tview.Pages
	Application *tview.Application
}

func NewApp() (*App, error) {
	app := &App{
		headerPanel: tview.NewFlex(),
		outputPanel: tview.NewList(),

		layout:      tview.NewFlex(),
		Application: tview.NewApplication(),
	}

	injector, err := pkg.NewInjector()
	if err != nil {
		return nil, err
	}

	pkg.Register(injector, themes.NewTheme)

	app.headerComponent, err = components.NewHeaderComponent(injector)
	if err != nil {
		return nil, err
	}
	app.headerPanel = app.headerComponent.CreateHeaderPanel()

	app.outputComponent, err = components.NewOutputComponent(injector)
	if err != nil {
		return nil, err
	}
	app.outputPanel = app.outputComponent.CreateOutputPanel()

	return app, nil
}

func (app *App) updateOutputPanel() {
	counter := 0
	for {
		time.Sleep(1 * time.Second)
		counter++

		app.Application.QueueUpdateDraw(func() {
			newMessage := fmt.Sprintf("Log entry %d: This is a new log message", counter)
			app.outputPanel.AddItem(newMessage, "", 0, nil)
		})
	}
}

func (app *App) Setup() {
	app.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(app.headerPanel, 0, 2, false).
		AddItem(app.outputPanel, 0, 10, false)

	go app.updateOutputPanel()
}

func (app *App) Run() error {
	app.Setup()

	app.pages = tview.NewPages()
	app.pages.AddPage("base", app.layout, true, true)

	return app.Application.SetRoot(app.pages, true).Run()
}
