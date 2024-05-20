package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	ch := make(chan int)

	for _, num := range nums {
		go func(n int) {
			ch <- n * n
		}(num)
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(<-ch)
	}
}
