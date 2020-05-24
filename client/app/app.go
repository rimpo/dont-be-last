package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/rimpo/dont-be-last/pkg/service/cli"
)

type application struct {
	app         fyne.App
	authService cli.AuthService
	gameService cli.GameService
}

func New(authService cli.AuthService, gameService cli.GameService) *application {
	a := &application{
		app:         app.New(),
		authService: authService,
		gameService: gameService,
	}
	lw := newLoginWindow(authService, a.app)

	lw.CreateAndShow(
		func() {
			gw := newGameWindow(a.app)
			gw.CreateAndShow(func() {
				lw.window.Show()
			})
		},
		func() {},
	)
	return a
}

func (a *application) Run() {

	a.app.Run()
}
