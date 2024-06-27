package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const (
	minNodes int = 2
	maxNodes int = 500

	defaultNodeCount   float64 = 200
	defaultSpeed       float64 = .6
	defaultCoupling    float64 = 0
	defaultVariability float64 = 1

	width  float32 = 960
	height float32 = 720
)

type parameters struct {
	nodeCount   float64
	speed       float64
	coupling    float64
	variability float64
}

func main() {
	a := app.New()
	window := a.NewWindow("Kuramogo")
	parameters := &parameters{
		nodeCount:   defaultNodeCount,
		speed:       defaultSpeed,
		coupling:    defaultCoupling,
		variability: defaultVariability,
	}

	kuramotoRenderer := &renderer{}
	content := kuramotoRenderer.render()
	go kuramotoRenderer.animate(parameters)

	controlPanel := makeSliders(window, parameters)

	window.SetContent(container.NewVBox(content, controlPanel))
	window.Resize(fyne.NewSize(width, height))
	window.ShowAndRun()
}
