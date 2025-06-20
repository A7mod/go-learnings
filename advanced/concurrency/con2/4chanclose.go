package con2

import "fmt"

func ClosingChannel() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	close(ch) // channel closed, no more sending

	// trying to read after channel is closed!!
	val1 := <-ch
	fmt.Println("Received:", val1)

	val2 := <-ch
	fmt.Println("Received:", val2)

	val3 := <-ch
	fmt.Println("Received:", val3)

}
