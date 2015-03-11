package main

import (
	"fmt"
	"image"
	"image/color"
)

// heavily borrowed from: http://nbviewer.ipython.org/github/barbagroup/CFDPython/tree/master/lessons/

type SimConstants struct {
	baseIntensity, c, tstep, xstep float64
}

func main() {

	size := 41
	// t := 1.0

	maxIntensity := float64(2.0)

	const width = float64(2)
	consts := SimConstants{
		baseIntensity: 1.0,
		c:             float64(1.0),
		tstep:         float64(0.025),
		xstep:         float64(width / float64(size-1)),
	}
	runs := 20 //int(t / consts.tstep)

	graph := image.NewRGBA(image.Rect(0, 0, runs, size))
	a := make([]float64, size)
	next := make([]float64, size)

	for i := 0; i < size; i++ {
		a[i] = consts.baseIntensity
	}
	for i := 15; i < 26; i++ {
		a[i] = maxIntensity
	}

	for i := 0; i < runs; i++ {

		// draw row
		for j := 0; j < len(a); j++ {
			color := createColor(a[j], maxIntensity)
			graph.Set(i, j, color)
		}
		//make next row
		update(a[:], next[:], consts)
		a, next = next, a
		fmt.Println(sum(&a))
	}

	save(graph, "new.png")
	show("new.png")

}

func update(prev []float64, next []float64, consts SimConstants) {
	ratio := consts.tstep / consts.xstep
	size := len(prev)
	// first box
	next[0] = prev[0] - ratio*prev[0]*(prev[0]-consts.baseIntensity)
	for i := 1; i < size; i++ {
		next[i] = prev[i] - ratio*prev[i]*(prev[i]-prev[i-1])
	}
	return
}

func createColor(point float64, maxIntensity float64) (ans color.Color) {
	intensity := uint8(point * (255 / maxIntensity))
	ans = color.RGBA{intensity, intensity, intensity, 255}
	return
}
