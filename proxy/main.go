package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func main() {
	// 1. Listen for connection forever.
	if ln, err := net.Listen("tcp", ":8080"); err == nil {
		for {
			// 2. Accept connection.
			if conn, err := ln.Accept(); err == nil {
				reader := bufio.NewReader(conn)
				// 3. Read requests from the client
				if req, err := http.ReadRequest(reader); err == nil {
					fmt.Println(req)
				}
			}
		}
	}
}
