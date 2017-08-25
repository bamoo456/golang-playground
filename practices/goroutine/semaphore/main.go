package main

import (
	"fmt"
	"sync"
	"time"
)

// using channel and goroutine to demo the semaphore pattern

func main() {
	maxWorkers := 10
	sema := make(chan struct{}, maxWorkers)
	out := make(chan int)

	worker := func(jobId int) {
		// acquire a semaphore
		sema <- struct{}{}
		defer func() { <-sema }()

		// pretend doing some task
		<-time.After(time.Duration(100) * time.Millisecond)

		out <- jobId
	}

	var wg sync.WaitGroup

	go func() {
		wg.Wait()
		close(out)
	}()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i + 1)
	}

	for jobId := range out {
		fmt.Printf("Job id %d is completed\n", jobId)
	}

}
