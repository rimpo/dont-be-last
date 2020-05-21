package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type gameWindow struct {
	cells  map[string]*widget.Button
	window fyne.Window
}

func newGameWindow() *gameWindow {
	// a.window.SetContent(widget.NewVBox(
	// 	widget.NewLabel("Hello Fyne!"),
	// 	widget.NewButton("Quit", func() {
	// 		a.app.Quit()
	// 	}),
	// ))
	return &gameWindow{}
}

func (gw *gameWindow) Load(app fyne.App) {
	gw.window = app.NewWindow("Play")
	gw.window.Show()
}
