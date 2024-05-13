package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 3

	s = RemoveElement(i, s)

	fmt.Println(s)
}

func RemoveElement(i int, s []int) []int {
	if i >= len(s) {
		return s
	}
	// slice before element + slice after element
	return append(s[:i], s[i+1:]...)
}
