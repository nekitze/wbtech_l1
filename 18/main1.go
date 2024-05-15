package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterWithAtomic struct {
	value int64
}

func (c *CounterWithAtomic) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *CounterWithAtomic) Value() int64 {
	return c.value
}

func main() {
	counter := CounterWithAtomic{}

	wg := sync.WaitGroup{}

	// 100 * 5 = 500
	for i := 0; i < 100; i++ {
		wg.Add(1)
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
