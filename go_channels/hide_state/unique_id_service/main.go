// Unique ID service
// A goroutine is started to generate unique hexadecimal id’s.
// Each id is sent via the id channel and the goroutine halts until
// the channel is read. Each time the channel is read, the goroutine is free
// to increment the value and send another.
package main

import (
	"fmt"
)

func main() {
	id := make(chan string)

	go func() {
		var counter int64 = 1
		for i := 0; i < 10; i++ {
			id <- fmt.Sprintf("%x", counter)
			counter++
		}
		close(id)
	}()

	for c := range id {
		fmt.Printf("%x\n", c)
	}
}
