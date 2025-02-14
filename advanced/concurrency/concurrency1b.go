package concurrency

import (
	"fmt"
	"time"
)

// this is an example of multiple goroutines, how they run "concurrently"

func printMessage(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond) // to cause simulation of delay
	}
}

func ShowGoroutines2() {
	fmt.Println("Starting Multiple Goroutines....")

	go printMessage("Goroutine 1:")
	go printMessage("Goroutine 2:")
	go printMessage("Goroutine 3:")

	fmt.Println("Main function execution continues...")

	time.Sleep(3 * time.Second)
	fmt.Println("Main function finished!")

	// this is interesting!

	fmt.Println("----------------------------")
}
