package concurrency

import "fmt"

func ShowBufferedLoop() {
	ch := make(chan string, 3) //channel of type string and size 3

	messages := []string{"Hello", "I", "Am", "Slim-Shady"}

	// sending through loop, but as you can see it can fit 3 at a time. only.
	for i := 0; i < 3; i++ {
		ch <- messages[i]
	}

	// reading from the buffer
	for i := 0; i < 3; i++ {
		fmt.Println("Received:", <-ch)
	}

	// thus this way Slim-Shady never gets printed.
	// since the channel is not big enough to hold 4 strings.

}
