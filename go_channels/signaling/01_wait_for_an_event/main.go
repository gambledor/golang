package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	fmt.Println("Started...")
	go func() {
		// ... do some stuff
		time.Sleep(time.Second * 5)
		close(c)
	}()

	// halt for communcation of data via the channel or for it to be closed.
	<-c

	fmt.Println("Done")
}
