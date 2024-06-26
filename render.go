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
	nodes  [Nmax]*node
}

func (r *renderer) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	middle := fyne.NewPos(size.Width/2, size.Height/2)
	radius := fyne.Min(size.Width, size.Height) / 2

	x, y := r.getMeanNodePosition()
	for _, n := range r.nodes {
		n.kuramoto(K, x, y)
		n.updateNodeState(dTime)
		n.updatePosition(radius, middle)
		n.redraw()
	}
}

func (r *renderer) getMeanNodePosition() (float64, float64) {
	sumX, sumY := float32(0), float32(0)
	N := float64(Nmax)
	for _, n := range r.nodes {
		sumX += n.x
		sumY += n.y
	}
	return float64(sumX) / N, float64(sumY) / N
}

func (r *renderer) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(600, 600)
}

func (r *renderer) render() *fyne.Container {
	container := container.NewWithoutLayout()
	for i := range r.nodes {
		dOmega := rand.Float64() - .5
		r.nodes[i] = &node{
			circle: canvas.Circle{},
			dOmega: dOmega * sigma,
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
	tick := time.NewTicker(time.Second / 60)
	go func() {
		for {
			<-tick.C
			r.Layout(nil, co.Size())
		}
	}()
}
