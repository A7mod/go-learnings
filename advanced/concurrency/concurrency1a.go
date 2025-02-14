package concurrency

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Numbers:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func ShowGoroutines() {
	fmt.Println("Starting Goroutine....")

	go printNumbers() // calling function as a goroutine

	fmt.Println("Main function keeps on executing..")

	time.Sleep(3 * time.Second)
	fmt.Println("Main function Finished!")

	fmt.Println("----------------------------")

}
