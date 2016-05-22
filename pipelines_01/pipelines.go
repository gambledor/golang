// Pepeline example form golang.org/blog/pipelines
package main

import "fmt"

func main(){
	// set up the pipeline
	printer(square(gen(2, 3, 4, 5)))
	printer(square(square(gen(2, 3, 4, 5))))
}

// First pipeline stage
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
// Second pipeline stage
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
// Consumes and prints the channel itmes.
func printer(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}
