//coverage:ignore file
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const (
	n int = 500

	dTime float64 = .002

	defaultSpeed float64 = .6
	maxSpeed     float64 = 10

	defaultCoupling float64 = 0
	maxCoupling     float64 = 3

	defaultVariability float64 = 1
)

func main() {
	a := app.New()
	window := a.NewWindow("Kuramogo")

	kuramotoRenderer := &renderer{
		k:     defaultCoupling,
		speed: defaultSpeed,
		sigma: defaultVariability,
	}
	content := kuramotoRenderer.render()
	go kuramotoRenderer.animate(content)

	controlPanel := makeSliders(window, kuramotoRenderer)

	window.SetContent(container.NewVBox(content, controlPanel))

	var width float32 = 960
	var height float32 = 720
	windowSize := fyne.NewSize(width, height)

	window.Resize(windowSize)
	window.ShowAndRun()
}
