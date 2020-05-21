package app

import (
	"fmt"

	"fyne.io/fyne"
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
	token         string
	playerId      int
}

func newLoginWindow(authService cli.AuthService) *loginWindow {
	return &loginWindow{
		authService: authService,
	}
}

func (lw *loginWindow) OnLoginPress() {
	token, playerId, err := lw.authService.Login(lw.userNameEntry.Text, lw.passwordEntry.Text)
	if err != nil {
		fmt.Println("error")
		return
	}
	lw.token = token
	lw.playerId = playerId
	lw.window.Close()
}

func (lw *loginWindow) Load(app fyne.App) {
	lw.window = app.NewWindow("Login")
	lw.userNameEntry = widget.NewEntry()
	lw.passwordEntry = widget.NewPasswordEntry()
	lw.loginBtn = widget.NewButton("login", func() {
		lw.OnLoginPress()
	})
	lw.window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("username:"), lw.userNameEntry),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("password:"), lw.passwordEntry),
			lw.loginBtn,
		))
	lw.window.Show()
}
