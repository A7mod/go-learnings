package con3

import (
	"fmt"
	"sync"
)

func MutexwithWaitGroup() {
	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	// add 5 groutuines
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()

			mutex.Lock()
			counter++
			fmt.Printf("Goroutine %d incremented counter to %d\n", id, counter)
			mutex.Unlock()
		}(i)
	}

	// wait until all goroutines finish

	wg.Wait()

	fmt.Println("Final Counter Value", counter)

}
