package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const sliderSteps float64 = 500

func makeSliders(_ fyne.Window, p *parameters) fyne.CanvasObject {
	nSlider := makeBoundSliderWithLabel(&p.nodeCount, float64(minNodes), float64(maxNodes), "%.0f", "N")
	speedSlider := makeBoundSliderWithLabel(&p.speed, 0, 1, "%.3f", "Speed")
	couplingSlider := makeBoundSliderWithLabel(&p.coupling, 0, 1, "%.3f", "Coupling")
	variabilitySlider := makeBoundSliderWithLabel(&p.variability, 0, 1, "%.3f", "Variability")

	return container.NewVBox(
		nSlider,
		speedSlider,
		couplingSlider,
		variabilitySlider,
	)
}

func makeBoundSliderWithLabel(
	variable *float64,
	min float64,
	max float64,
	format string,
	variableName string) *fyne.Container {
	boundVariable := binding.BindFloat(variable)
	slider := widget.NewSliderWithData(min, max, boundVariable)
	slider.Step = max / sliderSteps
	label := widget.NewLabelWithData(
		binding.FloatToStringWithFormat(boundVariable, variableName+": \t"+format),
	)
	return container.NewVBox(label, slider)
}
