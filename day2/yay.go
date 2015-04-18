package main

import (
	"fmt"
	"image"
	"image/color"
)

// heavily borrowed from: http://nbviewer.ipython.org/github/barbagroup/CFDPython/tree/master/lessons/

type SimConstants struct {
	baseIntensity, c, tstep, xstep float64
	numSteps, numBoxes             int
}

type ViewConstants struct {
	maxIntensity float64
}

func main() {

	// t := 1.0
	// physicalWidth = float64(2.0)
	// runs := 20 //int(t / simConsts.tstep)

	simConsts := SimConstants{
		numSteps:      20,
		numBoxes:      41,
		baseIntensity: 1.0,
		c:             float64(1.0),
		tstep:         float64(0.025),
	}
	simConsts.xstep = float64(2.0 / float64(simConsts.numBoxes-1))

	viewConsts := ViewConstants{maxIntensity: 2.0}

	width := simConsts.numSteps
	height := simConsts.numBoxes

	graph := image.NewRGBA(image.Rect(0, 0, width, height))

	currRow := make([]float64, simConsts.numBoxes)
	nextRow := make([]float64, simConsts.numBoxes)

	for i := 0; i < simConsts.numBoxes; i++ {
		currRow[i] = simConsts.baseIntensity
	}
	for i := 15; i < 26; i++ {
		currRow[i] = viewConsts.maxIntensity
	}

	for i := 0; i < simConsts.numSteps; i++ {
		drawRow(i, currRow[:], graph, viewConsts)
		nextTimeStep(currRow[:], nextRow[:], simConsts)
		fmt.Println(sum(&currRow))
		currRow, nextRow = nextRow, currRow
	}

	save(graph, "new.png")
	show("new.png")

}

func nextTimeStep(prev []float64, next []float64, simConsts SimConstants) {
	ratio := simConsts.tstep / simConsts.xstep
	size := len(prev)
	// first box
	next[0] = prev[0] - ratio*prev[0]*(prev[0]-simConsts.baseIntensity)
	for i := 1; i < size; i++ {
		next[i] = prev[i] - ratio*prev[i]*(prev[i]-prev[i-1])
	}
	return
}

func drawRow(rowNum int, row []float64, graph *image.RGBA, viewConsts ViewConstants) {
	for j := 0; j < len(row); j++ {
		color := createColor(row[j], viewConsts.maxIntensity)
		graph.Set(rowNum, j, color)
	}
}

func createColor(pointIntensity float64, maxIntensity float64) (ans color.Color) {
	intensity := uint8(pointIntensity * (255 / maxIntensity))
	ans = color.RGBA{intensity, intensity, intensity, 255}
	return
}
