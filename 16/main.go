package main

import (
	"fmt"
)

func quickSort(arr []int) {
	sortPartition(arr, 0, len(arr)-1)
}

func sortPartition(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivot := arr[(low+high)/2]
	lowI := low
	highI := high

	for lowI <= highI {
		for arr[lowI] < pivot {
			lowI++
		}
		for arr[highI] > pivot {
			highI--
		}
		if lowI <= highI {
			arr[lowI], arr[highI] = arr[highI], arr[lowI]
			lowI++
			highI--
		}
	}

	sortPartition(arr, low, highI)
	sortPartition(arr, lowI, high)
}

func main() {
	arr := []int{5, 1, 2, 6, 4, 7, 7, 10, 8}
	quickSort(arr)
	fmt.Println(arr)
}
