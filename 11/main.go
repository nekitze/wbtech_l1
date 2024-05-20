package main

import "fmt"

func intersect(a, b []int) (intersection []int) {
	m := make(map[int]int)

	for _, v := range a {
		m[v]++
	}

	for _, v := range b {
		m[v]++
	}

	for k, v := range m {
		if v > 1 {
			intersection = append(intersection, k)
		}
	}

	return intersection
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6}

	intersection := intersect(a, b)
	fmt.Println(intersection)
}
