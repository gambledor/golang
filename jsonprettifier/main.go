// Package main
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	text := make([]byte, 2)
	for scanner.Scan() {
		var bytes []byte
		bytes = scanner.Bytes()
		text = append(text, bytes...)
	}

	fmt.Println(text)
	var result []interface{}
	if err := json.Unmarshal(text, &result); err != nil {
		panic(err)
	}
	fmt.Println("%+v", result)
}
