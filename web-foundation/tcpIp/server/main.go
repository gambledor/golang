// Package main provides the entry point
package main

import (
	"fmt"
	"net"
)

func main() {
	if listener, error := net.Listen("tcp", ":1024"); error == nil {
		for {
			if connection, error := listener.Accept(); error == nil {
				go func(c net.Conn) {
					defer c.Close()
					fmt.Fprintln(c, "hello world")
				}(connection)
			}
		}
	}
}
