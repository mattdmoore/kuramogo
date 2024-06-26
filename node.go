package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const (
	jitterPercent float32 = 0.15
)

type node struct {
	circle canvas.Circle
	size   fyne.Size
	jitter float32
	active bool

	omega  float64
	dOmega float64
	dTheta float64
	theta  float64

	dx float64

	x float32
	y float32
}

func (n *node) kuramoto(k float64, x float64, y float64) {
	if n.active {
		n.dx = (math.Cos(n.theta)*y - math.Sin(n.theta)*x) * k * maxCoupling
	} else {
		n.dx = 0
	}
}

func (n *node) updateNodeState(dt float64, speed float64, sigma float64) {
	n.omega = 1 + n.dOmega*sigma
	n.dTheta = dt * speed * maxSpeed * (n.omega + n.dx)
	n.theta += n.dTheta
	n.x = float32(math.Cos(n.theta))
	n.y = float32(math.Sin(n.theta))
}

func (n *node) updatePosition(radius float32, middle fyne.Position) {
	x := n.x*(radius-n.jitter*radius*jitterPercent) + middle.X
	y := n.y*(radius-n.jitter*radius*jitterPercent) + middle.Y
	n.circle.Move(fyne.NewPos(x, y))
}

func (n *node) redraw(sigma float64) {
	var yOffsetMultiplier float64 = 100

	if n.active {
		n.circle.Show()
		n.circle.FillColor = color.CMYK{
			uint8((n.dOmega + .5) * 256 * sigma),
			uint8((.5 - n.dOmega) * 256 * sigma),
			uint8(yOffsetMultiplier * sigma),
			uint8((.5 - n.dOmega) * 256 * (1 - sigma)),
		}
		n.circle.Resize(n.size)
		n.circle.Refresh()
	} else {
		n.circle.Hide()
	}
}
