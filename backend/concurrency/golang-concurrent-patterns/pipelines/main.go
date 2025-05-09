package pipeline

import (
	"fmt"
)

func AddOnPipe[X, Y any](q <-chan int, f func(X) Y, in <-chan X) chan Y {
	output := make(chan Y)
	go func() {
		defer close(output)
		for {
			select {
			case <-q:
				return
			case input := <-in:
				output <- f(input)
			}
		}
	}()
	return output
}

func main() {
	input := make(chan int)
	quit := make(chan int)
	output := AddOnPipe(quit, Box,
		AddOnPipe(quit, AddToppings,
			AddOnPipe(quit, Bake,
				AddOnPipe(quit, Mixture,
					AddOnPipe(quit, PrepareTray, input)))))
	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-output, "received")
	}
}
