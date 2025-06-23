package con3

import "fmt"

func generatorFanIn(out chan<- int, values ...int) {
	for _, v := range values {
		out <- v
	}
}

func FanIn() {
	out := make(chan int, 10)

	go generatorFanIn(out, 1, 2, 3)
	go generatorFanIn(out, 4, 5, 6)

	// Read 6 values from a combined output
	for i := 1; i <= 6; i++ {
		fmt.Println("Received:", <-out)
	}

}
