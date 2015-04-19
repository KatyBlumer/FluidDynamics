package main

import (
	"fmt"
	"github.com/KatyBlumer/FluidDynamics/drawing"
	"math"
)

// heavily borrowed from: http://nbviewer.ipython.org/github/barbagroup/CFDPython/tree/master/lessons/

type SimConstants struct {
	baseIntensity, maxIntensity, c, viscosity, tstep, xstep, sigma float64
	numSteps, numBoxes                                             int
}

func main() {

	// t := 1.0
	// physicalWidth = float64(2.0)
	// runs := 20 //int(t / simConsts.tstep)

	simConsts := SimConstants{
		numSteps:      1000,
		numBoxes:      400,
		baseIntensity: 1.0,
		maxIntensity:  2.0,
		c:             float64(1.0),
		viscosity:     0.2,
	}
	simConsts.xstep = float64(2.0 / float64(simConsts.numBoxes-1))
	simConsts.sigma = 1.0 / simConsts.maxIntensity
	simConsts.tstep = simConsts.sigma * math.Pow(simConsts.xstep, 2) / simConsts.viscosity

	width := simConsts.numSteps
	height := simConsts.numBoxes

	graph := drawing.InitGraph(width, height)

	currRow := make([]float64, simConsts.numBoxes)
	nextRow := make([]float64, simConsts.numBoxes)

	for i := 0; i < simConsts.numBoxes; i++ {
		currRow[i] = simConsts.baseIntensity
	}
	for i := 160; i < 240; i++ {
		currRow[i] = simConsts.maxIntensity
	}

	for i := 0; i < simConsts.numSteps; i++ {
		drawing.DrawRow(i, currRow[:], graph, simConsts.maxIntensity)
		nextTimeStepBurgers(currRow[:], nextRow[:], simConsts)
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

func nextTimeStepDiffusion(curr []float64, next []float64, simConsts SimConstants) {
	ratio := simConsts.viscosity * simConsts.tstep / math.Pow(simConsts.xstep, 2)
	size := len(curr)
	// first box
	// next[0] = curr[0] - ratio*(2*curr[1]-2*curr[0])
	for i := 1; i < size-1; i++ {
		next[i] = curr[i] + ratio*(curr[i+1]-2*curr[i]+curr[i-1])
	}
	// next[size-1] = curr[size-1] - ratio*(2*curr[size-2]-2*curr[size-1])
	return
}

func nextTimeStepBurgers(curr []float64, next []float64, simConsts SimConstants) {
	coeff1 := simConsts.tstep / simConsts.xstep
	coeff2 := simConsts.viscosity * simConsts.tstep / math.Pow(simConsts.xstep, 2)

	size := len(curr)

	for i := 0; i < size; i++ {
		left, right := i-1, i+1
		if left < 0 {
			left = size - 1
		}
		if right >= size {
			right = 0
		}
		next[i] = curr[i] - coeff1*curr[i]*(curr[i]-curr[left]) + coeff2*(curr[right]-2*curr[i]+curr[left])
	}
	return
}
