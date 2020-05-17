package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/rimpo/dont-be-last/pkg/service/cli"
)

type application struct {
	app         fyne.App
	window      fyne.Window
	authService cli.AuthService
	gameService cli.GameService
}

func New(authService cli.AuthService, gameService cli.GameService) *application {
	app := app.New()
	window := app.NewWindow("Login")

	a := &application{
		app:         app,
		window:      window,
		authService: authService,
		gameService: gameService,
	}

	a.window.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			a.app.Quit()
		}),
	))
	return a
}

func (a *application) Run() {
	a.window.ShowAndRun()
}
