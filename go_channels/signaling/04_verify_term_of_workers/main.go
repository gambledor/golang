// Verify termination of workers.
// a goroutine is started, waiting for communication of data on the die channel
// (or for it to be closed). In this case, once closed, the goroutine performs
// termination tasks, then signals to the main goroutine (via the same diechannel)
// that it’s finished.
package main

import (
	"fmt"
)

func worker(die chan bool) {
	for {
		select {
		// ... do stuff cases
		case <-die:
			// ... do termination tasks
			fmt.Print("worker terminated.")
			die <- true
			return
		}
	}
}

func main() {
	die := make(chan bool)
	go worker(die)
	die <- true

	// wait until the goroutine has terminated
	<-die
}
