package con2

import "fmt"

func RangeChannel() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // important as otherwise loop below will hang.??
	}()

	for val := range ch {
		fmt.Println("Received:", val)
	}

}
