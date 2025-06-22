package con3

import (
	"fmt"
	"sync"
)

func RaceCondition() {
	var counter int
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			counter++
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Final Counter (with Race condition):", counter)

	// this is giving random numbers, but it should give 1000 as output.!
}
