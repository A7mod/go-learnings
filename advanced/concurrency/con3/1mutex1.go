package con3

import (
	"fmt"
	"sync"
)

func MutexBasic() {
	var counter int
	var mutex sync.Mutex

	// goroutine 1
	go func() {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}()

	//goroutine 2
	go func() {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}()

	// abhi wait groups use nahi kiye h na
	fmt.Println("This may print before go-routine finishes:", counter)

}
