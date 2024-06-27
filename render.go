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

const (
	refreshRate time.Duration = 60
	nodeSize    float32       = 8
	minRadius   float32       = 600
)

type renderer struct {
	content fyne.CanvasObject
	nodes   [maxNodes]*node
}

func (r *renderer) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	middle := fyne.NewPos(size.Width*.5, size.Height*.5)
	radius := fyne.Min(size.Width, size.Height) * .5

	for _, node := range r.nodes {
		node.redraw(node.position(radius, middle))
	}
}

func (r *renderer) meanNodePosition(n float64) (float64, float64) {
	if n == 0 {
		return 0, 0
	}
	var sumX, sumY float32
	for _, node := range r.nodes {
		if node.active {
			sumX += node.x
			sumY += node.y
		}
	}
	return float64(sumX) / n, float64(sumY) / n
}

func (r *renderer) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(minRadius, minRadius)
}

func (r *renderer) render() *fyne.Container {
	container := container.NewWithoutLayout()

	for i := range r.nodes {
		dOmega := rand.Float64() - .5
		r.nodes[i] = &node{
			circle: canvas.Circle{},
			size:   fyne.NewSize(nodeSize, nodeSize),
			dOmega: dOmega * defaultVariability,
			theta:  rand.Float64() * 2 * math.Pi,
			jitter: rand.Float32(),
		}
		container.Add(&r.nodes[i].circle)
	}

	container.Layout = r
	r.content = container
	return container
}

func (r *renderer) animate(p *parameters) {
	tick := time.NewTicker(time.Second / refreshRate)
	go func() {
		for {
			x, y := r.meanNodePosition(p.nodeCount)
			for i, node := range r.nodes {
				node.active = i < int(p.nodeCount)
				node.kuramoto(*p, x, y)
				node.setColor(p.variability)
			}
			r.Layout(nil, r.content.Size())
			<-tick.C
		}
	}()
}
