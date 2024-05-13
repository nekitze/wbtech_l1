package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -0.1, 0, 0.33}
	step := 10

	set := make(map[int][]float64)
	for _, t := range temperatures {
		set[int(t)/step*step] = append(set[int(t)/step*step], t)
	}

	for k, subset := range set {
		fmt.Printf("%d:%v\n", k, subset)
	}
}
