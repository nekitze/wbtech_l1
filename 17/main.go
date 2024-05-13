package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) (int, bool) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if target < arr[mid] {
			right = mid - 1
		} else if target > arr[mid] {
			left = mid + 1
		} else {
			return mid, true
		}
	}

	return 0, false
}

func main() {
	arr := []int{1, 1, 2, 3, 4, 5, 7, 10, 11}
	target := 7

	index, found := binarySearch(arr, target)

	if found {
		fmt.Printf("Index of element %d is %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found!\n", target)
	}
}
