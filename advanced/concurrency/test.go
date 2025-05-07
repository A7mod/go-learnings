package concurrency

import "fmt"

func ShowKaise() {
	ch := make(chan string)
	ch <- "Hello"     // Deadlock! No receiver is there.
	fmt.Println(<-ch) // This line will never run
}
