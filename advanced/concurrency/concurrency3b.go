//buffered channels blocking when full..

package concurrency

import "fmt"

func ShowBufferedChannelBlocking() {
	ch := make(chan int, 2) //channel of type int and size 2

	fmt.Println("Sending 1:")
	ch <- 1

	fmt.Println("Sending 2:")
	ch <- 2

	fmt.Println("trying to send 3, (will block if there is no receiver)")
	// This will block unless someone receives from the channel, abhi 2 jaa chuke h na bhai somoene has to read that first, tab jaega 3.
	// ch <- 3  comment out krna pada kyuki mast error aarhi h - ["fatal error: all goroutines are asleep - deadlock!"]

	fmt.Println("Reading", <-ch)

	fmt.Println("Now sending 3 after reading:")
	ch <- 3

	fmt.Println("Reading", <-ch)
	fmt.Println("Reading", <-ch)

}
