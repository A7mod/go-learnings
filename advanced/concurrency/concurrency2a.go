package concurrency

import "fmt"

func ShowBasicChannel() {
	// creating a channel of type int

	ch := make(chan int)
	vu := make(chan int)

	// goroutine to send data into channel
	go func() {
		ch <- 10 // sending 10 into channel
		go func() {
			vu <- 11
		}()

		ch <- 12
		vu <- 13

	}()

	// Receiving data from channel

	value := <-ch // Blocking calls, waiting for data
	noice := <-vu
	bhel := <-ch
	nice := <-vu

	fmt.Println("Received: value1", value)
	fmt.Println("Received: value2", value)
	fmt.Println("Received: value3", nice)
	fmt.Println("Received: noice", noice)
	fmt.Println("Received: bhel", bhel)
	fmt.Println("Received: value4", value)
	fmt.Println("Received: bhel", bhel)
	fmt.Println("----------------------------")

}
