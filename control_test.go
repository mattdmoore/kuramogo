package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestMakeBoundSliderWithLabel(t *testing.T) {
	test.NewApp()

	val := .5
	name := "Test"
	testSlider := makeBoundSliderWithLabel(&val, 0., 1., "%.2f", name)

	assert.NotNil(t, testSlider)
}

func TestMakeSliders(t *testing.T) {
	a := test.NewApp()
	window := a.NewWindow("Test")
	parameters := &parameters{
		nodeCount:   defaultNodeCount,
		speed:       defaultSpeed,
		coupling:    defaultCoupling,
		variability: defaultVariability,
	}

	testControlPanel := makeSliders(window, parameters)
	assert.NotNil(t, testControlPanel)
}
