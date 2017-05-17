// Package main provides entry point
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	address       string
	secureAddress string
	certificate   string
	key           string
)

var servers sync.WaitGroup

func init() {
	if address = os.Getenv("SERVE_HTTP"); address == "" {
		address = ":1024"
	}
	if secureAddress = os.Getenv("SERVE_HTTPS"); secureAddress == "" {
		secureAddress = ":1025"
	}
	if certificate = os.Getenv("SERVE_CERT"); certificate == "" {
		certificate = "cert.pem"
	}
	if key = os.Getenv("SERVE_KEY"); key == "" {
		key = "key.pem"
	}

	go SingnalHandler(make(chan os.Signal, 1))
}

func main() {
	message := "hello world"

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, message)
	})

	log.Printf("About to listen on http %s and https %s. Go to https://globuntu:%s", address, secureAddress, secureAddress)

	Launch("HTTP", func() error {
		return http.ListenAndServe(address, nil)
	})

	Launch("HTTPS", func() error {
		return http.ListenAndServeTLS(secureAddress, certificate, key, nil)
	})

	servers.Wait()
}

// Launch launches the function given as parameter
func Launch(name string, f func() error) {
	servers.Add(1)
	go func() {
		defer servers.Done()
		if e := f(); e != nil {
			fmt.Println(name, "->", e)
			syscall.Kill(syscall.Getpid(), syscall.SIGABRT)
		}
	}()
}

// SingnalHandler handles Interrupt signal
func SingnalHandler(c chan os.Signal) {
	signal.Notify(c, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := <-c; ; s = <-c {
		log.Printf("Got signal:", s)
		switch s {
		case os.Interrupt:
			fmt.Println("^C received")
			fmt.Println("interrupt - continue running")
		case syscall.SIGABRT:
			fmt.Println("abnormal exit")
			os.Exit(1)
		case syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("clean shutdown")
			os.Exit(0)
		}
	}
}
