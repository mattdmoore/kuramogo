//coverage:ignore file
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const (
	Nmax int = 1000

	dTime       float64 = .001
	maxSpeed    float64 = 10
	maxCoupling float64 = 3
)

var (
	x, y float32
	K    float64

	speed float64 = .5
	sigma float64 = .75
)

func main() {
	a := app.New()
	w := a.NewWindow("Kuramogo")

	kuramotoRenderer := &renderer{}
	content := kuramotoRenderer.render()
	go kuramotoRenderer.animate(content)

	controlPanel := makeSliders(w)

	w.SetContent(container.NewVBox(content, controlPanel))
	w.Resize(fyne.NewSize(960, 720))
	w.ShowAndRun()
}
