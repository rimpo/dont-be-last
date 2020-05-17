package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/rimpo/dont-be-last/pkg/service/cli"
	"github.com/rimpo/dont-be-last/pkg/service/cli/auth"
	"github.com/rimpo/dont-be-last/pkg/service/cli/game"
)

type GameWindow struct {
	window      fyne.Window
	authService cli.AuthService
	gameService cli.GameService
}

func NewGameWindow(authService cli.AuthService, gameService cli.GameService) *GameWindow {
	gw := &GameWindow{
		window:      app.NewWindow("Hello"),
		authService: authService,
		gameService: gameService,
	}
	gw.window = app.NewWindow("Hello")

	gw.window.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))
	return gw
}
func (gw *GameWindow) Run() {
	gw.window.ShowAndRun()
}

func main() {
	app := app.New()

	authService := auth.NewService("http://127.0.0.1:5000/auth")
	gameService := game.NewService()

	gameWindow = NewGameWindow(authService, gameService)

	gameWindow.Run()
}
