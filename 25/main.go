package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)

	//start := time.Now()
	//for time.Since(start) < duration {}
}

func main() {
	start := time.Now()
	sleep(2 * time.Second)
	fmt.Printf("Отдохнули %f сек\n", time.Since(start).Seconds())
}
