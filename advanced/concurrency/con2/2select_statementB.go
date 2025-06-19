package con2

import (
	"fmt"
	"time"
)

func ShowSelectNotBasic() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Data from Go-routine"
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println("Received:", msg)
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout! No data yet...")
		default:
			fmt.Println("Doing something else...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
