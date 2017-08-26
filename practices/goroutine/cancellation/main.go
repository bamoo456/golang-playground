package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	quit := make(chan struct{})
	semaphore := make(chan struct{}, 3)
	resCh := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		jId := i
		go func() {
			defer wg.Done()

			select {
			// if goroutine receive the quit signal, it will cancel its job
			case <-quit:
				fmt.Printf("job #%d is canceled\n", jId)
				return

			// if goroutine acquire an semaphore, it will complete its job
			case semaphore <- struct{}{}:
				defer func() { <-semaphore }()
				worker()
				fmt.Printf("job #%d is completed\n", jId)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	// this goroutine is trying to simulate the "quit" sender
	go func() {
		<-time.After(500 * time.Millisecond)
		fmt.Println("==== trigger the quit ====")
		close(quit)
	}()

loop:
	for {
		select {
		// if it get the quit signal, it will drain the all result from current queue, and break the loop
		case <-quit:
			// must drain of the resCh to ensure not active worker goroutine is there
			for range resCh {
				// drain the result from resCh
			}
			break loop

		// consume the result normally
		case _, ok := <-resCh:
			if !ok {
				fmt.Println("All jobs are completed")
				break loop
			}
		}
	}

	fmt.Println("Completed the main function")
}

func worker() {
	d := 300 * time.Millisecond
	<-time.After(d)
}
