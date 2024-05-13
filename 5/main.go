package main

import (
	"fmt"
	"time"
)

func main() {
	n := 0
	fmt.Println("Enter counter:")
	fmt.Scan(&n)

	start := time.Now()
	ch := make(chan int)

	go func() {
		for data := range ch {
			fmt.Println(data)
		}
	}()

	for i := 0; int(time.Since(start).Seconds()) < n; i++ {
		ch <- i
	}
	close(ch)
}
