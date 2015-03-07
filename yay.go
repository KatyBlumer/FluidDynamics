package main

import (
	"fmt"
	"image"
	"image/color"
	// "image/draw"
	"image/png"
	"math"
	"log"
	"os"
	"os/exec"
)

var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	black color.Color = color.RGBA{0, 0, 0, 255}
	blue  color.Color = color.RGBA{0, 0, 255, 255}
)



func main() {
	
	size := 100
	runs := 100
	start := float64(10.0)
	scale := 1.0

	graph := image.NewRGBA(image.Rect(0, 0, runs, size))
	a := make([]float64, size)

	a[size/5] = start
	a[size/2] = start
	
	for i := 0; i < runs; i++ {		
		
		// draw row
		for j := 0; j < len(a); j++ {
			color := createColor(a[j], start, scale)
			graph.Set(i, j, color)
		}
		fmt.Println(sum(&a))

		//make next row
		a = update(&a)
	}

	save(graph, "new.png")
	show("new.png")
	
}

func update(prev *[]float64) (next []float64) {
	size := len(*prev)
	fmt.Println(size)
	fmt.Println(prev)
	fmt.Println(*prev)
	// first box
	next[0] = ((2 * (*prev)[0]) + (*prev)[1]) / 3
	for j := 1; j < size-1; j++ {
		next[j] = ((*prev)[j-1] + (*prev)[j] + (*prev)[j+1]) / 3
	}
	// last box
	next[size-1] = ((*prev)[size-2] + (2 * (*prev)[size-1])) / 3
	return
}

func createColor(point float64, maxIntensity float64, scale float64) (ans color.Color) {
	intensity := 255 - uint8(math.Log(point * (255/maxIntensity)))
	ans = color.RGBA{0, intensity, 0, 255}
	return
}

func sum(a *[]float64) (s float64) {
	for _, v := range *a {
		s += v
	}
	return
}

// save and show  a specified file by Preview.app for OS X(darwin)
func save(im image.Image, filename string) {
	w, _ := os.Create(filename)
	defer w.Close()
	png.Encode(w, im) //Encode writes the Image m to w in PNG format.
}

func show(filename string) {
	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, filename)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}