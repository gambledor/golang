// Package main provides the command to create
// a path directory
package main

import (
	"flag"
	"fmt"
	"os"
)

func mkdirp(path string) {
	err := os.MkdirAll(path, 0766)
	if err != nil {
		panic(err)
	}
}

// This does not change the shell directory
func changeDir(path string) {
	// fmt.Println("Hello changeDir with param ", path)
	err := os.Chdir(path)
	if err != nil {
		panic(err)
	}
}

func main() {

	cdFlag := flag.Bool("cd", false, "To change directory after it has been created")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Directory name missing!")
		os.Exit(2)
	}

	path := flag.Arg(0)
	mkdirp(path)

	// fmt.Printf("cd flag %t", *cdFlag)
	if *cdFlag {
		changeDir(path)
	}
}
