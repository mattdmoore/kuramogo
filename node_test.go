package main

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/stretchr/testify/assert"
)

const float64EqualityThreshold = 1e-9

func TestKuramoto(t *testing.T) {
	n := node{
		circle: canvas.Circle{},
		dOmega: .7,
	}
	k := 1.5 / maxCoupling

	x, y := .1, .2
	n.kuramoto(k, x, y)
	assert.InDelta(t, .3, n.dx, float64EqualityThreshold)
}

func TestUpdateNodeState(t *testing.T) {
	n := node{
		circle: canvas.Circle{},
		dOmega: .7,
	}
	n.dx = .3

	speed := 1 / maxSpeed
	sigma := 1.

	n.updateNodeState(dTime, speed, sigma)
	assert.InDelta(t, .002, n.theta, float64EqualityThreshold)
}

func TestUpdatePosition(t *testing.T) {
	n := node{
		circle: canvas.Circle{},
		dOmega: .7,
	}
	n.x, n.y = 0, 1

	radius := float32(100)
	middle := fyne.NewPos(100, 100)

	n.updatePosition(radius, middle)
	assert.Equal(t, fyne.NewPos(100, 200), n.circle.Position1)
}
