//coverage:ignore file
package main

import (
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type renderer struct {
	canvas fyne.CanvasObject
	nodes  [n]*node

	k     float64
	speed float64
	sigma float64
}

func (r *renderer) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	middle := fyne.NewPos(size.Width*.5, size.Height*.5)
	radius := fyne.Min(size.Width, size.Height) * .5

	x, y := r.getMeanNodePosition()
	for _, n := range r.nodes {
		n.kuramoto(r.k, x, y)
		n.updateNodeState(dTime, r.speed, r.sigma)
		n.updatePosition(radius, middle)
		n.redraw(r.sigma)
	}
}

func (r *renderer) getMeanNodePosition() (float64, float64) {
	var sumX, sumY float32 = 0, 0
	for _, node := range r.nodes {
		sumX += node.x
		sumY += node.y
	}
	return float64(sumX) / float64(n), float64(sumY) / float64(n)
}

func (r *renderer) MinSize(_ []fyne.CanvasObject) fyne.Size {
	var radius float32 = 600
	return fyne.NewSize(radius, radius)
}

func (r *renderer) render() *fyne.Container {
	container := container.NewWithoutLayout()

	var nodeSize float32 = 8
	for i := range r.nodes {
		dOmega := rand.Float64() - .5
		r.nodes[i] = &node{
			circle: canvas.Circle{},
			size:   fyne.NewSize(nodeSize, nodeSize),
			dOmega: dOmega * r.sigma,
			theta:  rand.Float64() * 2 * math.Pi,
			jitter: rand.Float32(),
		}
		container.Add(&r.nodes[i].circle)
	}

	container.Layout = r
	r.canvas = container
	return container
}

func (r *renderer) animate(co fyne.CanvasObject) {
	var refreshRate time.Duration = 60
	tick := time.NewTicker(time.Second / refreshRate)
	go func() {
		for {
			r.Layout(nil, co.Size())
			<-tick.C
		}
	}()
}
