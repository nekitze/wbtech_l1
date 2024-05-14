package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	values := [5]int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := sync.WaitGroup{}

	// Пишем в первый канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, value := range values {
			ch1 <- value
			time.Sleep(time.Second)
		}
	}()

	// Отправляем значение * 2 из первого канала во второй
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(values); i++ {
			ch2 <- 2 * <-ch1
		}
	}()

	// Выводим значения из второго канала в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(values); i++ {
			fmt.Println(<-ch2)
		}
	}()

	wg.Wait()
}
