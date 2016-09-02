// Package main provides that any variable can be used as interface
// So how can we use this feature to pass any type of variable to function?
// fmt.Println can accept any type as argument if it implements the Stringer interface
package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

// Human implements fmt.Striger
func (h Human) String() string {
	return "Name:" + h.name + ", Age: " + strconv.Itoa(h.age) + " years, Contact:" + h.phone
}

func main() {
	Alice := Human{"Alice", 39, "000-777-XXX"}
	fmt.Println("This human is: ", Alice)
}
