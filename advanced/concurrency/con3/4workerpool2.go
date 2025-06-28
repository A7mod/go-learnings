package con3

import (
	"fmt"
	"sync"
	"time"
)

func doWork(id, job int) string {
	time.Sleep(500 * time.Millisecond)

	return fmt.Sprintf("Worker %d finished job %d.", id, job)
}

func WorkerPoolResults() {
	const numJobs = 6
	const numWorkers = 2

	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				output := doWork(workerID, job)
				results <- output
			}
		}(w)
	}

	// send jobs

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// wait for workers to finish, then close channels

	go func() {
		wg.Wait()
		close(results)
	}()

	// read results

	for res := range results {
		fmt.Println(res)
	}

}
