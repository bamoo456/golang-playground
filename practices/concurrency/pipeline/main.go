package main

import (
	"fmt"
)

func main() {
	fmt.Println("start the pipeline demo")
	for i := 0; i < 5; i++ {
		initCh := make(chan int)
		go func(input chan int) {
			pipeOut(pipeIn(pipeIn(pipeIn(input))))
		}(initCh)
		initCh <- i
		// close initCh channel
		close(initCh)
	}
}

func pipeIn(in chan int) chan int {
	out := make(chan int)
	go func() {
		idx := <-in
		fmt.Println("pipeline in with idx ", idx)
		idx++
		out <- idx
		// close out channel
		close(out)
	}()
	return out
}

func pipeOut(in chan int) {
	idx := <-in
	fmt.Println("pipeline out with idx ", idx)
}
