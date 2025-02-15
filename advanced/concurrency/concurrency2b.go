package concurrency

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Hello from Goroutine!"
}

func ShowChannelFunction() {
	ch := make(chan string)

	go sendData(ch)

	fmt.Println("Waiting for data...")
	message := <-ch
	fmt.Println("Received:", message)
}
