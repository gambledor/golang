// Fan in fan out example
package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2,3,4,5)
	// Distribute the sq work across two goroutines that both
	// read from in channel.
	c1 := square(in)
	c2 := square(in)
	// consumes the merged output from c1 and c2
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}
// First stage
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}
// Second stage
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
// Third stage
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	// Start an output goroutine for each input channel in cs.
	// output copies values from c to out until c is closed
	// then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	// Start a goroutine to close out once all the output goroutines are done.
	// This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
