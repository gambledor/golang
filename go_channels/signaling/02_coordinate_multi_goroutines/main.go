// In this example, a hundred goroutines are started, waiting for communication of data on
// the start channel (or for it to be closed). In this case, once closed, all goroutines start.
package main

import (
	"fmt"
	"time"
)

func worker(start chan bool) {
	<-start
	// do staff
	fmt.Print("ciao ")
}

func main() {
	start := make(chan bool)

	for i := 0; i < 100; i++ {
		go worker(start)
		// start <- true
	}

	close(start)
	// ... all workers running now
	time.Sleep(time.Second * 5)
}
