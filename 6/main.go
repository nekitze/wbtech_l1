package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start worker which will be stopped by context cancel")
	ctx, cancel := context.WithCancel(context.Background())
	go workerContext(ctx, "With cancel")
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

	fmt.Println("-------------------")

	fmt.Println("Start worker which will be stopped by context timeout")
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go workerContext(ctx, "With timeout")
	time.Sleep(4 * time.Second)

	fmt.Println("-------------------")

	fmt.Println("Start worker which will be stopped by context deadline")
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	go workerContext(ctx, "With deadline")
	time.Sleep(4 * time.Second)

	fmt.Println("-------------------")

	fmt.Println("Start worker which will be stopped by channel signal")
	stop := make(chan bool)
	go workerChan(stop, "Chan1")
	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(2 * time.Second)

	fmt.Println("-------------------")

	fmt.Println("Start worker which will be stopped by closing channel")
	ch := make(chan bool)
	go workerChan(ch, "Chan2")
	time.Sleep(3 * time.Second)
	close(ch)
	time.Sleep(2 * time.Second)
}

func workerChan(stop chan bool, name string) {
	for {
		select {
		case <-stop:
			fmt.Printf("worker %s stopped\n", name)
			return
		default:
			fmt.Println("worker working")
			time.Sleep(time.Second)
		}
	}
}

func workerContext(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %s stopped\n", name)
			return
		default:
			fmt.Println("worker working")
			time.Sleep(time.Second)
		}
	}
}
