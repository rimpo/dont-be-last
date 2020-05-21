package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
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

	w := a.app.NewWindow("Hello")

	loginContent := widget.NewVBox()
	boardContent := widget.NewVBox()

	userNameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	loginContent.Append(widget.NewButton("Quit", func() {
		a.app.Quit()
	}))
	loginContent.Append(fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("username:"), userNameEntry),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("password:"), passwordEntry),
		widget.NewButton("login", func() {
			w.SetContent(boardContent)
		}),
	))

	player1 := widget.NewEntry()
	player2 := widget.NewPasswordEntry()

	boardContent.Append(widget.NewLabel("Game Board!"))
	boardContent.Append(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("player 1:"), player1),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("player 2:"), player2),
		),
	)
	boardContent.Append(
		widget.NewButton("Logout", func() {
			w.SetContent(loginContent)
		}),
	)

	w.SetContent(loginContent)

	w.ShowAndRun()

	// lw := newLoginWindow(authService)
	// lw.Load(a.app)
	// gw := newGameWindow()
	// gw.Load(a.app)
	return a
}

func (a *application) Run() {
	//a.app.Run()
}
