package main

import "fmt"

func main() {
	fmt.Println("Start the fanOut-fanIn pattern")
	input := make(chan int)
	output := make(chan int)
	done := make(chan int)
	concurrents := 10

	// fanOut
	fanOut(input, concurrents, output)

	// start put data into workers
	for i := 0; i < 10; i++ {
		input <- i
	}

	// fanIN
	go func() {
		sum := 0
		for i := 0; i < concurrents; i++ {
			sum += <-output
		}
		done <- sum
		close(input)
		close(output)
		close(done)
	}()

	res := <-done
	fmt.Println("Final result is ", res)
}

func fanOut(in chan int, concurrentCount int, out chan int) {
	for i := 0; i < concurrentCount; i++ {
		go func(idx int) {
			fmt.Printf("Perform worker %d \n", idx)
			newInput := make(chan int)
			data := <-in
			// pipeline for doning some work
			res := pipe(newInput)
			newInput <- data * 2
			// Waiting for pipeline result
			for v := range res {
				out <- v
			}
			fmt.Printf("Done worker %d \n", idx)
		}(i)
	}
}

func pipe(in chan int) chan int {
	out := make(chan int)
	go func() {
		value := <-in
		out <- value * 2
		close(out)
	}()
	return out
}
