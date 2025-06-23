package con3

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func Fanout() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// Start 3 workers fan-out
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs) // close input channel
	wg.Wait()

}
