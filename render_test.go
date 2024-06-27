package main

import (
	"math"
	"math/rand"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func setup() *renderer {
	test.NewApp()
	testRenderer := &renderer{}
	for i := range testRenderer.nodes {
		dOmega := rand.Float64() - .5
		testRenderer.nodes[i] = &node{
			circle: canvas.Circle{},
			size:   fyne.NewSize(nodeSize, nodeSize),
			dOmega: dOmega * defaultVariability,
			theta:  rand.Float64() * 2 * math.Pi,
			jitter: rand.Float32(),

			active: true,
		}
	}
	return testRenderer
}

func TestLayoutInitial(t *testing.T) {
	testRenderer := setup()

	initialNodeState := testRenderer.nodes
	testRenderer.Layout(nil, fyne.NewSize(width, height))
	for i, n := range testRenderer.nodes {
		assert.Equal(t, initialNodeState[i].circle, n.circle)
	}
}

func TestMeanNodePositionZero(t *testing.T) {
	testRenderer := setup()

	x, y := testRenderer.meanNodePosition(defaultNodeCount)

	assert.Zero(t, x)
	assert.Zero(t, y)
}

func TestMeanNodePositionNonZero(t *testing.T) {
	testRenderer := setup()

	for _, node := range testRenderer.nodes {
		node.x = float32(math.Cos(node.theta))
		node.y = float32(math.Sin(node.theta))
	}

	x, y := testRenderer.meanNodePosition(defaultNodeCount)
	assert.NotZero(t, x)
	assert.NotZero(t, y)
}

func TestMeanNodePositionBadNodeCount(t *testing.T) {
	testRenderer := setup()

	x, y := testRenderer.meanNodePosition(0)

	assert.Zero(t, x)
	assert.Zero(t, y)
}

func TestRenderCreatesContent(t *testing.T) {
	testRenderer := setup()

	content := testRenderer.render()
	assert.NotNil(t, content)
}
