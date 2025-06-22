package con3

import (
	"fmt"
	"sync"
)

func RaceConditionFixed() {
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Race condition fixed with mutex:", counter)

}
