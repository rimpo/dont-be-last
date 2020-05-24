package app

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/rimpo/dont-be-last/pkg/service/cli"
)

type loginWindow struct {
	window        fyne.Window
	authService   cli.AuthService
	userNameEntry *widget.Entry
	passwordEntry *widget.Entry
	loginBtn      *widget.Button
	app           fyne.App
	token         string
	playerId      int
}

func newLoginWindow(authService cli.AuthService, app fyne.App) *loginWindow {
	return &loginWindow{
		authService: authService,
		app:         app,
		window:      app.NewWindow("Login"),
	}
}

func (lw *loginWindow) onLoginPress() {
	token, playerId, err := lw.authService.Login(lw.userNameEntry.Text, lw.passwordEntry.Text)
	if err != nil {
		fmt.Println("error")
		return
	}
	lw.token = token
	lw.playerId = playerId
	lw.window.Hide()

	lw.userNameEntry.SetText("")
	lw.passwordEntry.SetText("")
}

func (lw *loginWindow) CreateAndShow(onSuccess func(), onFailure func()) {
	lw.userNameEntry = widget.NewEntry()
	lw.passwordEntry = widget.NewPasswordEntry()
	lw.loginBtn = widget.NewButton("login", func() {
		token, playerId, err := lw.authService.Login(lw.userNameEntry.Text, lw.passwordEntry.Text)
		if err != nil {
			dialog.ShowInformation("Login failed", "Incorrect username or password", lw.window)
			onFailure()
			return
		}
		lw.token = token
		lw.playerId = playerId

		lw.userNameEntry.SetText("")
		lw.passwordEntry.SetText("")

		//hide login window
		lw.window.Hide()
		// call the success callback
		onSuccess()
	})
	lw.window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("username:"), lw.userNameEntry),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("password:"), lw.passwordEntry),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewButton("quit", func() { lw.window.Close() }), lw.loginBtn),
		))
	lw.window.Show()
}
