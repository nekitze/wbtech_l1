package main

import (
	"fmt"
	"time"
)

func main() {
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
