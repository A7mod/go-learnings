package con3

import (
	"fmt"
	"sync"
	"time"
)

func workers(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond) // simulate work
	}
}

func WorkerPool() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// start the worker pool
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done() // only once per goroutine
			for job := range jobs {
				fmt.Printf("Worker %d started job %d\n", workerID, job)
				time.Sleep(1 * time.Second)
				fmt.Printf("Worker %d finished job %d\n", workerID, job)
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // no more jobs

	wg.Wait()

}
