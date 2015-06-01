package web

import (
	"fmt"
	"github.com/KatyBlumer/FluidDynamics/fluids"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/lastFrame", showLastFrame)
	http.HandleFunc("/gif", showGif)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

const image = `
<html>
  <body>
    <img src="%v">
  </body>
</html>
`

func showLastFrame(w http.ResponseWriter, r *http.Request) {
	fluids.MakeLastFrame()
	fmt.Fprintf(w, image, "../lastFrame.png")
}

// DOESN'T WORK - need to use Google Compute Engine for ffmpeg
func showGif(w http.ResponseWriter, r *http.Request) {
	fluids.MakeGif()
	fmt.Fprintf(w, image, "../graph.gif")
}
