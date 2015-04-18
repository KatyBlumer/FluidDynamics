package main

func sum(a *[]float64) (s float64) {
	for _, v := range *a {
		s += v
	}
	return
}
