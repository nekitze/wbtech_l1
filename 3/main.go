package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	var (
		sum int
		mu  sync.Mutex
		wg  sync.WaitGroup
	)

	wg.Add(len(nums))

	for _, num := range nums {
		go func(n int) {
			defer wg.Done()
			square := n * n

			mu.Lock()
			sum += square
			mu.Unlock()
		}(num)
	}

	wg.Wait()
	fmt.Println(sum)
}
