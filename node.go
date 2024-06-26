package main

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const (
	jitter_scale float32 = 0.15
)

type node struct {
	circle canvas.Circle

	omega  float64
	dOmega float64
	dTheta float64
	theta  float64
	jitter float32
	dx     float64

	x float32
	y float32
}

func (n *node) kuramoto(r *renderer, K float64) {
	n.dx = (math.Cos(n.theta)*r.mean_y - math.Sin(n.theta)*r.mean_x) * K
	n.omega = 1 + n.dOmega
	n.dTheta = r.dTime * (n.omega + n.dx)
	n.theta += n.dTheta
	n.x = float32(math.Cos(n.theta))
	n.y = float32(math.Sin(n.theta))
}

func (n *node) updatePosition(radius float32, middle fyne.Position) {
	x = n.x*(radius-n.jitter*radius*jitter_scale) + middle.X
	y = n.y*(radius-n.jitter*radius*jitter_scale) + middle.Y
	n.circle.Move(fyne.NewPos(x, y))
	n.circle.Resize(fyne.NewSize(8, 8))
}
