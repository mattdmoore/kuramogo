package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const (
	jitter_scale float32 = 0.15
)

type node struct {
	circle canvas.Circle
	jitter float32

	omega  float64
	dOmega float64
	dTheta float64
	theta  float64

	dx float64

	x float32
	y float32
}

func (n *node) kuramoto(K float64, x float64, y float64) {
	n.dx = (math.Cos(n.theta)*y - math.Sin(n.theta)*x) * K * maxCoupling
}

func (n *node) updateNodeState(dt float64) {
	n.omega = 1 + n.dOmega*sigma
	n.dTheta = dt * speed * maxSpeed * (n.omega + n.dx)
	n.theta += n.dTheta
	n.x = float32(math.Cos(n.theta))
	n.y = float32(math.Sin(n.theta))
}

func (n *node) updatePosition(radius float32, middle fyne.Position) {
	x = n.x*(radius-n.jitter*radius*jitter_scale) + middle.X
	y = n.y*(radius-n.jitter*radius*jitter_scale) + middle.Y
	n.circle.Move(fyne.NewPos(x, y))
}

func (n *node) redraw() {
	n.circle.FillColor = color.CMYK{
		uint8((n.dOmega + .5) * math.Pow(2, 8) * sigma),
		uint8((.5 - n.dOmega) * math.Pow(2, 8) * sigma),
		uint8(100 * sigma),
		uint8((.5 - n.dOmega) * math.Pow(2, 8) * (1 - sigma)),
	}
	n.circle.Resize(fyne.NewSize(8, 8))
	n.circle.Refresh()
}
