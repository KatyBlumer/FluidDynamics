package main

import (
	"fmt"
	"github.com/KatyBlumer/FluidDynamics/day2/drawing"
)

// heavily borrowed from: http://nbviewer.ipython.org/github/barbagroup/CFDPython/tree/master/lessons/

type SimConstants struct {
	baseIntensity, c, tstep, xstep float64
	numSteps, numBoxes             int
}

func main() {

	// t := 1.0
	// physicalWidth = float64(2.0)
	// runs := 20 //int(t / simConsts.tstep)

	simConsts := SimConstants{
		numSteps:      4000,
		numBoxes:      4000,
		baseIntensity: 1.0,
		c:             float64(1.0),
		tstep:         float64(0.00025),
	}
	simConsts.xstep = float64(2.0 / float64(simConsts.numBoxes-1))

	viewConsts := drawing.ViewConstants{MaxIntensity: 2.0}

	width := simConsts.numSteps
	height := simConsts.numBoxes

	graph := drawing.InitGraph(width, height)

	currRow := make([]float64, simConsts.numBoxes)
	nextRow := make([]float64, simConsts.numBoxes)

	for i := 0; i < simConsts.numBoxes; i++ {
		currRow[i] = simConsts.baseIntensity
	}
	for i := 200; i < 250; i++ {
		currRow[i] = viewConsts.MaxIntensity
	}

	for i := 0; i < simConsts.numSteps; i++ {
		drawing.DrawRow(i, currRow[:], graph, viewConsts)
		nextTimeStepNonLinearConvectionTrapezoidalX(currRow[:], nextRow[:], simConsts)
		fmt.Println(sum(&currRow))
		currRow, nextRow = nextRow, currRow
	}

	drawing.Show(graph)

}

func nextTimeStepAverage(curr []float64, next []float64, simConsts SimConstants) {
	size := len(curr)
	next[0] = curr[0]
	for i := 1; i < size-1; i++ {
		next[i] = (curr[i-1] + curr[i] + curr[i+1]) / 3
	}
	next[size-1] = curr[size-1]
}

func nextTimeStepNonLinearConvectionForwardDifferenceX(curr []float64, next []float64, simConsts SimConstants) {
	ratio := simConsts.tstep / simConsts.xstep
	size := len(curr)
	// first box
	next[0] = curr[0] - ratio*curr[0]*(curr[0]-simConsts.baseIntensity)
	for i := 1; i < size; i++ {
		next[i] = curr[i] - ratio*curr[i]*(curr[i]-curr[i-1])
	}
	return
}

func nextTimeStepNonLinearConvectionTrapezoidalX(curr []float64, next []float64, simConsts SimConstants) {
	ratio := simConsts.tstep / simConsts.xstep
	size := len(curr)
	// first box
	next[0] = curr[0] - ratio*curr[0]*(curr[0]-simConsts.baseIntensity)
	for i := 1; i < size-1; i++ {
		next[i] = curr[i] - (ratio/2)*curr[i]*(curr[i+1]-curr[i-1])
	}
	next[size-1] = curr[size-1] - ratio*curr[size-1]*(simConsts.baseIntensity-curr[size-1])
	return
}
