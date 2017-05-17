package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	if connection, err := net.Dial("tcp", ":1024"); err == nil {
		defer connection.Close()
		if text, err := bufio.NewReader(connection).ReadString('\n'); err == nil {
			fmt.Printf(text)
		}
	}
}
