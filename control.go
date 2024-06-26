package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func makeSliders(_ fyne.Window, r *renderer) fyne.CanvasObject {
	speedSlider := makeBoundSliderWithLabel(binding.BindFloat(&r.speed), "Speed")
	couplingSlider := makeBoundSliderWithLabel(binding.BindFloat(&r.k), "Coupling")
	variabilitySlider := makeBoundSliderWithLabel(binding.BindFloat(&r.sigma), "Variability")

	return container.NewVBox(
		speedSlider,
		couplingSlider,
		variabilitySlider,
	)
}

func makeBoundSliderWithLabel(
	boundVariable binding.ExternalFloat,
	variableName string) fyne.CanvasObject {
	slider := widget.NewSliderWithData(0, 1, boundVariable)
	slider.Step = .01
	label := widget.NewLabelWithData(
		binding.FloatToStringWithFormat(boundVariable, variableName+": \t%.2f"),
	)
	return container.NewVBox(label, slider)
}
