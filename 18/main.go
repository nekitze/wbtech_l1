package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	counter := Counter{}

	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println(counter.Value())
}
