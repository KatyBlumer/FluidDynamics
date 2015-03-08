package main

import (
	"image"
	"image/color"
	"fmt"
)

type SimConstants struct {
	c, tstep, xstep float64
}

func main() {
	
	size := 41
	t := 1.0
	
	maxIntensity := float64(10.0)

	const width = float64(2)
	consts := SimConstants {
		c: float64(1.0), 
		tstep: float64(0.005), 
		xstep: float64(width / float64(size)),
	}
	runs := int(t / consts.tstep)

	graph := image.NewRGBA(image.Rect(0, 0, runs, size))
	a := make([]float64, size)
	next := make([]float64, size)

	for i := 15; i < 26; i++ {
		a[i] = maxIntensity
	}
	
	for i := 0; i < runs; i++ {		
		
		// draw row
		for j := 0; j < len(a); j++ {
			color := createColor(a[j], maxIntensity)
			graph.Set(i, j, color)
		}
		fmt.Println(a)
		//make next row
		update(a[:], next[:], consts)
		a, next = next, a
	}

	save(graph, "new.png")
	show("new.png")
	
}

func update(prev []float64, next []float64, consts SimConstants) {
	ratio := (consts.c * consts.tstep / consts.xstep)
	size := len(prev)
	// first box
	next[0] = prev[0] - ratio * (prev[0])
	for i := 1; i < size; i++ {
		next[i] = prev[i] - ratio * (prev[i] - prev[i-1])
	}
	return
}

func createColor(point float64, maxIntensity float64) (ans color.Color) {
	intensity := uint8(point  * (255/maxIntensity))
	ans = color.RGBA{0, intensity, 0, 255}
	return
}