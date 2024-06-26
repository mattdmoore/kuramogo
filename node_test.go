package main

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/stretchr/testify/assert"
)

const float64EqualityThreshold = 1e-9

var defaultNode node = node{
	circle: canvas.Circle{},
	dOmega: .7,
}

func TestKuramoto(t *testing.T) {
	n := defaultNode
	K := 1.5
	x, y := .1, .2
	n.kuramoto(K, x, y)
	assert.InDelta(t, .3, n.dx, float64EqualityThreshold)
}

func TestUpdateNodeState(t *testing.T) {
	n := defaultNode
	n.dx = .3
	dt := .003

	speed = 1
	sigma = 1

	n.updateNodeState(dt)
	assert.InDelta(t, .006, n.theta, float64EqualityThreshold)
}

func TestUpdatePosition(t *testing.T) {
	n := defaultNode
	n.x, n.y = 0, 1

	radius := float32(100)
	middle := fyne.NewPos(100, 100)

	n.updatePosition(radius, middle)
	assert.Equal(t, fyne.NewPos(100, 200), n.circle.Position1)
}
