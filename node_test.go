package main

import (
	"image/color"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/stretchr/testify/assert"
)

const float64EqualityThreshold = 1e-9

func TestKuramotoCoupled(t *testing.T) {
	n := node{
		dOmega: .7,
		active: true,
	}

	parameters := &parameters{
		nodeCount:   defaultNodeCount,
		coupling:    1.5 / maxCoupling,
		speed:       defaultSpeed,
		variability: defaultVariability,
	}

	x, y := .1, .2
	n.kuramoto(*parameters, x, y)
	assert.InDelta(t, .3, n.dx, float64EqualityThreshold)
}

func TestKuramotoUncoupled(t *testing.T) {
	n := node{
		dOmega: .7,
		active: true,
	}

	parameters := &parameters{
		nodeCount:   defaultNodeCount,
		coupling:    0,
		speed:       defaultSpeed,
		variability: defaultVariability,
	}

	x, y := .1, .2
	n.kuramoto(*parameters, x, y)
	assert.InDelta(t, 0, n.dx, float64EqualityThreshold)
}

func TestPosition(t *testing.T) {
	n := node{
		circle: canvas.Circle{},
		active: true,
		x:      0,
		y:      1,
	}

	radius := float32(100)
	middle := fyne.NewPos(100, 100)

	position := n.position(radius, middle)
	assert.Equal(t, fyne.NewPos(100, 200), position)
}

func TestSetColorZeroVariability(t *testing.T) {
	n := node{
		circle: canvas.Circle{},
		active: true,
		dOmega: .75,
	}

	varibility := 0.
	n.setColor(varibility)
	expectedColor := color.CMYK{0x0, 0x0, 0x0, 0xc0}

	assert.Equal(t, expectedColor, n.circle.FillColor)
}

func TestSetColorMaxVariability(t *testing.T) {
	n := node{
		circle: canvas.Circle{},
		active: true,
		dOmega: .75,
	}

	varibility := 1.
	n.setColor(varibility)
	expectedColor := color.CMYK{0x40, 0xc0, 0x64, 0x0}

	assert.Equal(t, expectedColor, n.circle.FillColor)
}

func TestRedrawActive(t *testing.T) {
	n := node{
		circle: canvas.Circle{},
		active: true,
	}

	beforeState := n.circle
	position := fyne.NewPos(100, 200)
	n.redraw(position)

	assert.NotEqual(t, beforeState, n.circle)
}

func TestRedrawInactive(t *testing.T) {
	n := node{
		circle: canvas.Circle{Hidden: true},
		active: false,
	}

	beforeState := n.circle
	position := fyne.NewPos(100, 200)
	n.redraw(position)

	assert.Equal(t, beforeState, n.circle)
}
