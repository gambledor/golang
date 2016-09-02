// Package main provides how can we know the specific type stored in the interface.
// 1. Assertion of Comma-ok pattern
package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{"David", 65}

	for idx, elem := range list {
		switch value := elem.(type) { // elem.(type) cannot be used outside switch body
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", idx, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", idx, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", idx, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", idx)
		}
	}
}
