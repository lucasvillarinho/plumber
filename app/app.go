package app

import (
	"fmt"
	"time"

	"github.com/rivo/tview"

	"github.com/lucasvillarinho/plumber/app/component"
	configs "github.com/lucasvillarinho/plumber/config"
	di "github.com/lucasvillarinho/plumber/internal/injector"
	psb "github.com/lucasvillarinho/plumber/internal/pubsub"
)

type App struct {
	headerPanel     *tview.Flex
	headerComponent *component.HeaderComponent

	outputPanel     *tview.List
	outputComponent *component.OutputComponent

	layout      *tview.Flex
	pages       *tview.Pages
	Application *tview.Application

	pubSub *psb.PubSub[string]
}

func NewApplication(_ *di.Injector) (*tview.Application, error) {
	return tview.NewApplication(), nil
}

func NewApp() (*App, error) {
	app := &App{
		headerPanel: tview.NewFlex(),
		outputPanel: tview.NewList(),
		layout:      tview.NewFlex(),
	}

	injector, err := di.NewInjector()
	if err != nil {
		return nil, err
	}

	di.Register(injector, configs.NewThemeConfig)
	di.Register(injector, configs.NewPlumberConfig)
	di.Register(injector, psb.NewPubSub[string])
	di.Register(injector, NewApplication)

	app.headerComponent, err = component.NewHeaderComponent(injector)
	if err != nil {
		return nil, err
	}
	app.outputComponent, err = component.NewOutputComponent(injector)
	if err != nil {
		return nil, err
	}

	application, err := di.Get[*tview.Application](injector)
	if err != nil || application == nil {
		return nil, fmt.Errorf("failed to inject Theme instance: %w", err)
	}

	pubsub, err := di.Get[*psb.PubSub[string]](injector)
	if err != nil || pubsub == nil {
		return nil, fmt.Errorf("failed to inject PubSub instance: %w", err)
	}

	app.Application = *application
	app.outputPanel = app.outputComponent.CreateOutputPanel()
	app.headerPanel = app.headerComponent.CreateHeaderPanel()
	app.pubSub = *pubsub

	return app, nil
}

func (ap *App) updateOutputPanel() {
	counter := 0
	for {
		time.Sleep(1 * time.Second)
		counter++

		ap.Application.QueueUpdateDraw(func() {
			newMessage := fmt.Sprintf("Log entry %d: This is a new log message", counter)
			ap.outputPanel.AddItem(newMessage, "", 0, nil)
		})
	}
}

func (ap *App) Setup() {
	ap.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(ap.headerPanel, 0, 2, false).
		AddItem(ap.outputPanel, 0, 10, false)

	go ap.updateOutputPanel()
}

func (ap *App) Run() error {
	ap.Setup()

	ap.pages = tview.NewPages()
	ap.pages.AddPage("base", ap.layout, true, true)

	ap.startLogPublishing()

	return ap.Application.SetRoot(ap.pages, true).Run()
}

func (app *App) startLogPublishing() {
	go func() {
		counter := 0
		for {
			time.Sleep(1 * time.Second)
			counter++

			newMessage := fmt.Sprintf("Log entry %d: This is a new log message", counter)
			app.pubSub.Publish(newMessage)
		}
	}()
}
