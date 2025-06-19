package con2

import (
	"fmt"
	"time"
)

func SendData(ch chan<- string) {
	ch <- "Sent from Sender.\n"
}

func ReceiveData(ch <-chan string) {
	msg := <-ch
	fmt.Println("Received in Receiver:\n", msg)
}

func ShowChannelDirection() {
	ch := make(chan string)

	go SendData(ch)

	go ReceiveData(ch)

	fmt.Println("BKL kuch print ni hoga isme?")

	// Let the goroutines finish
	time.Sleep(1 * time.Second)

}
