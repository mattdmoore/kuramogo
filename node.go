package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const (
	timeIncrement float64 = .002
	jitterPercent float32 = 0.15

	maxSpeed    float64 = 10
	maxCoupling float64 = 3
)

type node struct {
	circle canvas.Circle
	size   fyne.Size
	jitter float32
	active bool

	omega  float64
	dOmega float64
	theta  float64
	dTheta float64

	dx float64

	x float32
	y float32
}

func (n *node) kuramoto(p parameters, x float64, y float64) {
	if n.active {
		// This is the Kuramoto model
		n.dx = (math.Cos(n.theta)*y - math.Sin(n.theta)*x) * p.coupling * maxCoupling
	} else {
		n.dx = 0
	}
	// Update node state
	n.omega = 1 + n.dOmega*p.variability
	n.dTheta = timeIncrement * p.speed * maxSpeed * (n.omega + n.dx)
	n.theta += n.dTheta
	n.x = float32(math.Cos(n.theta))
	n.y = float32(math.Sin(n.theta))
}

func (n *node) position(radius float32, middle fyne.Position) fyne.Position {
	// Return actual node position on screen
	x := n.x*(radius-n.jitter*radius*jitterPercent) + middle.X
	y := n.y*(radius-n.jitter*radius*jitterPercent) + middle.Y
	return fyne.NewPos(x, y)
}

func (n *node) setColor(variability float64) {
	// Set color saturation using variability modifier
	if n.active {
		var yOffsetMultiplier float64 = 100
		n.circle.FillColor = color.CMYK{
			uint8((n.dOmega + .5) * 256 * variability),
			uint8((.5 - n.dOmega) * 256 * variability),
			uint8(yOffsetMultiplier * variability),
			uint8((.5 - n.dOmega) * 256 * (1 - variability)),
		}
	}
}

func (n *node) redraw(position fyne.Position) {
	if n.active {
		n.circle.Show()
		n.circle.Resize(n.size)
		n.circle.Move(position)
		n.circle.Refresh()
	} else {
		n.circle.Hide()
	}
}
