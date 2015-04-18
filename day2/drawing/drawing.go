package drawing

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"os/exec"
)

type ViewConstants struct {
	MaxIntensity float64
}

func InitGraph(width, height int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

func DrawRow(rowNum int, row []float64, graph *image.RGBA, viewConsts ViewConstants) {
	for j := 0; j < len(row); j++ {
		color := createColor(row[j], viewConsts.MaxIntensity)
		graph.Set(rowNum, j, color)
	}
}

func createColor(pointIntensity float64, maxIntensity float64) (ans color.Color) {
	intensity := uint8(pointIntensity * (255 / maxIntensity))
	ans = color.RGBA{intensity, intensity, intensity, 255}
	return
}

func Show(im image.Image) {
	filename := "graph.png"
	save(im, filename)

	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, filename)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// save and show  a specified file by Preview.app for OS X(darwin)
func save(im image.Image, filename string) {
	w, _ := os.Create(filename)
	defer w.Close()
	png.Encode(w, im) //Encode writes the Image m to w in PNG format.
}
