package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const (
	Nmax  int     = 600
	sigma float64 = 2

	minSpeed float64 = .001
	maxSpeed float64 = .01

	maxCoupling float64 = 3
)

var (
	x, y float32
	K    float64
)

type renderer struct {
	dTime float64

	canvas fyne.CanvasObject

	nodes  [Nmax]*node
	mean_x float64
	mean_y float64
}

func (r *renderer) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	middle := fyne.NewPos(size.Width/2, size.Height/2)
	radius := fyne.Min(size.Width, size.Height) / 2

	r.updateMeanNodePosition()
	for _, n := range r.nodes {
		go n.kuramoto(r, K)
		go n.updatePosition(radius, middle)
	}
}

func (r *renderer) updateMeanNodePosition() {
	sumX, sumY := float32(0), float32(0)
	N := float64(Nmax)
	for _, n := range r.nodes {
		sumX += n.x
		sumY += n.y
	}
	r.mean_x = float64(sumX) / N
	r.mean_y = float64(sumY) / N
}

func (r *renderer) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(600, 600)
}

func (r *renderer) render() *fyne.Container {
	container := container.NewWithoutLayout()
	for i := range r.nodes {
		dOmega := rand.Float64()
		nodeColor := color.Gray16{uint16((dOmega + .5) * math.Pow(2, 15))}
		n := &canvas.Circle{FillColor: nodeColor}
		r.nodes[i] = &node{
			circle: *n,
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
	tick := time.NewTicker(time.Second / 144)
	go func() {
		for {
			<-tick.C
			r.Layout(nil, co.Size())
		}
	}()
}

func main() {
	a := app.New()
	w := a.NewWindow("Kuramogo")

	model := &renderer{dTime: .003}
	content := model.render()
	go model.animate(content)

	coupling := binding.BindFloat(&K)
	couplingSlider := widget.NewSliderWithData(0, maxCoupling, coupling)
	couplingSlider.Step = .01

	speed := binding.BindFloat(&model.dTime)
	speedSlider := widget.NewSliderWithData(minSpeed, maxSpeed, speed)
	speedSlider.Step = .0001

	w.SetContent(container.NewVBox(content, couplingSlider, speedSlider))
	w.Resize(fyne.NewSize(960, 720))
	w.ShowAndRun()
}
