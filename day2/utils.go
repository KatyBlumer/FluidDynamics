package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"os/exec"
)

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