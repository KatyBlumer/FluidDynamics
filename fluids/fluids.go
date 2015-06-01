package fluids

import (
	"fmt"
	"github.com/KatyBlumer/FluidDynamics/drawing"
	"math"
	"os"
)

// heavily borrowed from: http://nbviewer.ipython.org/github/barbagroup/CFDPython/tree/master/lessons/

type SimConstants struct {
	baseIntensity, maxIntensity, c, viscosity, tstep, xstep, ystep, sigma float64
	numTSteps, numXSteps, numYSteps                                       int
}

func initializeSimConstants(simConsts SimConstants) SimConstants {
	simConsts.xstep = float64(2.0 / float64(simConsts.numXSteps-1))
	simConsts.ystep = float64(2.0 / float64(simConsts.numYSteps-1))
	simConsts.sigma = 0.25
	simConsts.tstep = simConsts.sigma * math.Pow(simConsts.xstep, 2) / simConsts.viscosity

	return simConsts
}

func main() {
	simConsts := initializeSimConstants(SimConstants{
		numTSteps:     200,
		numXSteps:     400,
		numYSteps:     400,
		baseIntensity: 1.0,
		maxIntensity:  2.0,
		c:             1.0,
		viscosity:     0.05,
	})

	makeGif(simConsts, nextTimeStep2DDiffusion)
}

func makeGif(simConsts SimConstants, timeStepFn func([][]float64, [][]float64, SimConstants)) {

	// t := 1.0
	// physicalWidth = float64(2.0)
	// runs := 20 //int(t / simConsts.tstep)

	fileNameFormat := "graph%d.png"
	tempFolderName := "../temp/"
	gifFileName := "../graph.gif"
	clearFolder(tempFolderName)
	os.Remove(gifFileName)

	currFrame := make2DArray(simConsts.numXSteps, simConsts.numYSteps)
	nextFrame := make2DArray(simConsts.numXSteps, simConsts.numYSteps)

	for x := 0; x < simConsts.numXSteps; x++ {
		for y := 0; y < simConsts.numYSteps; y++ {
			if x >= 149 && x <= 250 && y >= 149 && y <= 250 {
				currFrame[x][y] = simConsts.maxIntensity
			} else {
				currFrame[x][y] = simConsts.baseIntensity
			}
		}
	}

	for t := 0; t < simConsts.numTSteps; t++ {
		drawing.SaveFrame(t, currFrame[:], simConsts.maxIntensity, tempFolderName+fileNameFormat)
		timeStepFn(currFrame[:], nextFrame[:], simConsts)
		fmt.Println(sum2D(currFrame))

		currFrame, nextFrame = nextFrame, currFrame
	}

	drawing.CreateGif(fileNameFormat, tempFolderName, gifFileName)
	drawing.ShowFile(gifFileName)
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

func nextTimeStep2DLinearConvection(curr [][]float64, next [][]float64, simConsts SimConstants) {
	xCoeff := simConsts.c * simConsts.tstep / simConsts.xstep
	yCoeff := simConsts.c * simConsts.tstep / simConsts.ystep

	xSize := len(curr)
	ySize := len(curr[0])

	for x := 0; x < xSize; x++ {
		left, right := x-1, x+1
		if left < 0 {
			left = xSize - 1
		}
		if right >= xSize {
			right = 0
		}
		for y := 0; y < ySize; y++ {
			top, bottom := y-1, y+1
			if top < 0 {
				top = ySize - 1
			}
			if bottom >= ySize {
				bottom = 0
			}
			next[x][y] = curr[x][y] - xCoeff*(curr[x][y]-curr[left][y]) - yCoeff*(curr[x][y]-curr[x][top])
		}
	}
	return
}

func nextTimeStep2DNonLinearConvection(curr, next, currV, nextV [][]float64, simConsts SimConstants) {
	xCoeff := simConsts.tstep / simConsts.xstep
	yCoeff := simConsts.tstep / simConsts.ystep

	xSize := len(curr)
	ySize := len(curr[0])

	for x := 0; x < xSize; x++ {
		left, right := x-1, x+1
		if left < 0 {
			left = xSize - 1
		}
		if right >= xSize {
			right = 0
		}
		for y := 0; y < ySize; y++ {
			top, bottom := y-1, y+1
			if top < 0 {
				top = ySize - 1
			}
			if bottom >= ySize {
				bottom = 0
			}
			next[x][y] = curr[x][y] -
				xCoeff*curr[x][y]*(curr[x][y]-curr[left][y]) -
				yCoeff*currV[x][y]*(curr[x][y]-curr[x][top])
			nextV[x][y] = currV[x][y] -
				xCoeff*curr[x][y]*(currV[x][y]-currV[left][y]) -
				yCoeff*currV[x][y]*(currV[x][y]-currV[x][top])
		}
	}
	return
}

func nextTimeStep2DDiffusion(curr, next [][]float64, simConsts SimConstants) {
	xCoeff := simConsts.viscosity * simConsts.tstep / (simConsts.xstep * simConsts.xstep)
	yCoeff := simConsts.viscosity * simConsts.tstep / (simConsts.ystep * simConsts.ystep)

	xSize := len(curr)
	ySize := len(curr[0])

	for x := 0; x < xSize; x++ {
		left, right := x-1, x+1
		if left < 0 {
			left = xSize - 1
		}
		if right >= xSize {
			right = 0
		}
		for y := 0; y < ySize; y++ {
			top, bottom := y-1, y+1
			if top < 0 {
				top = ySize - 1
			}
			if bottom >= ySize {
				bottom = 0
			}
			next[x][y] = curr[x][y] +
				xCoeff*(curr[left][y]-2*curr[x][y]+curr[right][y]) +
				yCoeff*(curr[x][top]-2*curr[x][y]+curr[x][bottom])
		}
	}
	addBoundaryConditions(next, simConsts.baseIntensity)
	return
}

func addBoundaryConditions(arr [][]float64, value float64) {
	arr[0] = makeArray(len(arr), value)
	arr[len(arr)-1] = makeArray(len(arr), value)
	for i := 0; i < len(arr); i++ {
		arr[i][0] = value
		arr[i][len(arr[0])-1] = value
	}
}
