package main

import (
	"fmt"
	"sync"
)

type MyConcurrentMap struct {
	Map map[int]int
	mu  sync.Mutex
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{Map: make(map[int]int)}
}

func (m *MyConcurrentMap) Add(key, value int) {
	m.mu.Lock()
	m.Map[key] = value
	m.mu.Unlock()
}

func main() {
	myMap := NewMyConcurrentMap()

	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			myMap.Add(i, i*10)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(myMap.Map)
}
