package main

func make2DArray(x, y int) [][]float64 {
	arr := make([][]float64, x)
	for i := 0; i < y; i++ {
		arr[i] = make([]float64, y)
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
