/**
A goroutine is started to recycle memory buffer.
The 'give' channel receives old memory buffers and stores them in a list.
While the 'get' channel dispenses these buffers for use.
If no buffers are available in the list, a new one is created.
*/
package main

import (
	"container/list"
)

func main() {
	give := make(chan []byte)
	get := make(chan []byte)

	go func() {
		q := new(list.List)

		for {
			if q.Len() == 0 {
				q.PushFront(make([]byte, 100))
			}

			e := q.Front()

			select {
			case s := <-give:
				q.PushFront(s)
			case get <- e.Value.([]byte):
				q.Remove(e)
			}
		}
	}()

	// gets a new buffer from the recycler.
	buffer := <-get
	buffer22 := <-get

	// give it back to the recycler.
	give <- buffer

	// get the recycled buffer again.
	buffer = <-get
	give <- buffer22
}
