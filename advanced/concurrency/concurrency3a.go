package concurrency

import (
	"fmt"
)

func ShowChannelBak() {
	ch := make(chan string, 2) //buffered channel with size 2

	ch <- "hello"
	ch <- "world"

	fmt.Println(<-ch) // Now it receives
	fmt.Println(<-ch) // Receives second value
}
