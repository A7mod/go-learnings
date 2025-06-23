package con3

import (
	"fmt"
	"sync"
)

func workerFaninFanout(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- fmt.Sprintf("Worker %d finished job %d\n", id, job)
	}
}

func FanOutFanIn() {
	jobs := make(chan int, 5)
	results := make(chan string, 5)
	var wg sync.WaitGroup

	// start  2 workers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go workerFaninFanout(i, jobs, results, &wg)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results (fan-in)
	for r := range results {
		fmt.Println(r)
	}

}
