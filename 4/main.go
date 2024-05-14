package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(ch <-chan int, id int) {
	for data := range ch {
		fmt.Printf("Workder %d: %d\n", id, data)
	}
}

func main() {
	var (
		workersCount int
		wg           sync.WaitGroup
		ch           chan int
	)
	ch = make(chan int)

	fmt.Println("Enter number of workers:")
	_, err := fmt.Scan(&workersCount)
	if err != nil {
		return
	}
	if workersCount == 0 {
		return
	}

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Worker(ch, i)
		}(i)
	}

	// Обработка сигнала Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; ; i++ {
		select {
		case <-sigChan:
			fmt.Println("laying off workers..")
			// Завершаем воркеров закрытием канала
			close(ch)
			wg.Wait()
			os.Exit(0)
		default:
			ch <- i
			time.Sleep(time.Second)
		}
	}

}
