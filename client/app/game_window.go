package app

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const (
	MaxRows     = 5
	MaxCell int = 15
	ALIVE       = ""
	DEAD        = ""
)

type triangleLayout struct {
	parent fyne.CanvasObject
}

func (t *triangleLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	//fmt.Println("minSize:", t.parent.Canvas().Size())
	//all button of same size
	size := objects[0].MinSize()
	return fyne.NewSize(size.Width*MaxRows, size.Height*MaxRows)
}
func (t *triangleLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	parentSize := t.parent.Size()
	totalSize := t.MinSize(objects)
	leftX := parentSize.Width - t.MinSize(objects).Width
	topY := parentSize.Height - t.MinSize(objects).Height
	centerX := t.MinSize(objects).Width/2 + leftX
	size := objects[0].MinSize()
	c := 0
	fmt.Println(parentSize, totalSize)
	for i := 0; i < MaxRows; i++ {
		x := centerX - ((i+1)*size.Width)/2
		for j := 0; j <= i; j++ {
			pos := fyne.NewPos(x+j*size.Width, i*size.Height+topY)
			objects[c].Resize(size)
			objects[c].Move(pos)
			c++
		}
	}
}

type gameWindow struct {
	cells  map[int]*widget.Button
	window fyne.Window
	app    fyne.App
}

func newGameWindow(app fyne.App) *gameWindow {
	return &gameWindow{
		app:    app,
		window: app.NewWindow("Play"),
		cells:  make(map[int]*widget.Button),
	}
}

func (gw *gameWindow) CreateAndShow(onClose func()) {
	gw.window.SetOnClosed(onClose)

	click := func(i int) func() {
		return func() {
			if gw.cells[i].Text == ALIVE {
				gw.cells[i].SetText(DEAD)
				gw.cells[i].SetIcon(nil)
			} else {
				dialog.ShowInformation("Error", "Already Dead !!", gw.window)
			}
		}
	}

	var cells []fyne.CanvasObject
	for i := 0; i < MaxCell; i++ {
		btn := widget.NewButton(ALIVE, click(i))
		btn.SetIcon(theme.FyneLogo())
		gw.cells[i] = btn
		cells = append(cells, btn)
	}

	gw.window.Resize(fyne.Size{150, 200})

	container := fyne.NewContainer()
	container.AddObject(fyne.NewContainerWithLayout(&triangleLayout{parent: container}, cells...))

	gw.window.SetContent(widget.NewVBox(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			container,
			widget.NewButton("Logout", func() {
				gw.window.Close()
			},
			),
		)))
	gw.window.Show()
}
