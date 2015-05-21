package main

import (
	"os"
)

func make2DArray(x, y int) [][]float64 {
	arr := make([][]float64, x)
	for i := 0; i < y; i++ {
		arr[i] = make([]float64, y)
	}
	return arr
}

func makeArray(size int, value float64) []float64 {
	arr := make([]float64, size)
	for i := 0; i < size; i++ {
		arr[i] = value
	}
	return arr
}

func sum(a []float64) (s float64) {
	for _, v := range a {
		s += v
	}
	return
}

func sum2D(a [][]float64) (s float64) {
	for i := range a {
		s += sum(a[i])
	}
	return
}

func clearFolder(folderName string) {
	os.RemoveAll(folderName)
	os.Mkdir(folderName, 0755)
}
