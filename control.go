package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func makeSliders(_ fyne.Window, r *renderer) fyne.CanvasObject {
	nSlider := makeBoundSliderWithLabel(binding.BindFloat(&r.n), float64(nMin), float64(nMax), "%.0f", "N")
	speedSlider := makeBoundSliderWithLabel(binding.BindFloat(&r.speed), 0, 1, "%.3f", "Speed")
	couplingSlider := makeBoundSliderWithLabel(binding.BindFloat(&r.k), 0, 1, "%.3f", "Coupling")
	variabilitySlider := makeBoundSliderWithLabel(binding.BindFloat(&r.sigma), 0, 1, "%.3f", "Variability")

	return container.NewVBox(
		nSlider,
		speedSlider,
		couplingSlider,
		variabilitySlider,
	)
}

func makeBoundSliderWithLabel(
	boundVariable binding.ExternalFloat,
	min float64,
	max float64,
	format string,
	variableName string) fyne.CanvasObject {
	slider := widget.NewSliderWithData(min, max, boundVariable)
	slider.Step = max / 1000
	label := widget.NewLabelWithData(
		binding.FloatToStringWithFormat(boundVariable, variableName+": \t"+format),
	)
	return container.NewVBox(label, slider)
}
