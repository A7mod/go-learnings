package con2

import (
	"fmt"
	"time"
)

func NewTicker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(6 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at:", t)
		case <-done:
			fmt.Println("Stopped ticker after 6 seconds.")
			return

		}

	}
}
