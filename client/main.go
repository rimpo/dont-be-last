package main

import (
	"github.com/rimpo/dont-be-last/client/app"
	"github.com/rimpo/dont-be-last/pkg/service/cli/auth"
	"github.com/rimpo/dont-be-last/pkg/service/cli/game"
)

/*type diagonal struct {
}

func (d *diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height-d.MinSize(objects).Height)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}

func (d *diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := 0, 0
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func main() {
	a := app.New()
	w := a.NewWindow("Diagonal")

	text1 := widget.NewLabel("topleft")
	text2 := widget.NewLabel("Middle Label")
	text3 := widget.NewLabel("bottomright")

	w.SetContent(fyne.NewContainerWithLayout(&diagonal{}, text1, text2, text3))
	w.ShowAndRun()
}*/

func main() {
	authService := auth.NewService("http://127.0.0.1:5000/auth")
	gameService := game.NewService()

	application := app.New(authService, gameService)

	application.Run()

}
