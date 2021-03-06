package drawing

import (
	// "bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"os/exec"
)

func InitGraph(width, height int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

func DrawRow(rowNum int, row []float64, graph *image.RGBA, maxIntensity float64) {
	for j := 0; j < len(row); j++ {
		color := createColor(row[j], maxIntensity)
		graph.Set(rowNum, j, color)
	}
}

func SaveFrame(t int, frame [][]float64, maxIntensity float64, fileNameFormat string) {
	graph := InitGraph(len(frame), len(frame[0]))
	for x := 0; x < len(frame); x++ {
		DrawRow(x, frame[x], graph, maxIntensity)
	}
	save(graph, fmt.Sprintf(fileNameFormat, t))
}

func CreateGif(fileNameFormat, tempFolderName, gifFileName string) {
	// running the command from Go doesn't work
	// command := "sh -c /Users/kblumer/go/src/github.com/KatyBlumer/FluidDynamics/drawing/ffmpeg.sh"
	// command := "ffmpeg.sh"
	command := "sh"
	arg1 := "-c"
	arg2 := "ffmpeg.sh"

	out, err := exec.Command(command, arg1, arg2).Output()
	log.Printf("%s", out)
	if err != nil {
		log.Fatal(err)
	}
}

func createColor(pointIntensity float64, maxIntensity float64) (ans color.Color) {
	if pointIntensity > maxIntensity {
		ans = color.RGBA{255, 0, 0, 255} // red
		return
	}
	if pointIntensity <= 0.0 {
		ans = color.RGBA{0, 255, 0, 255} //green
		return
	}
	intensity := uint8(pointIntensity * (255 / maxIntensity))
	ans = color.RGBA{intensity, intensity, intensity, 255}
	return
}

func Show(im image.Image, filename string) {
	save(im, filename)
	ShowFile(filename)
}

func ShowFile(filename string) {
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
