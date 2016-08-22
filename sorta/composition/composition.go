// Package compsition
package composition

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	CanSwim   bool
}

// Amy is embedded with the Human type
// and cah thus invoke methods in Human's method set
type Amy struct {
	Human
}

type Alan struct {
	Human
}

func (h *Human) Name() {
	fmt.Printf("Hello! My name is %v %v ", h.FirstName, h.LastName)
}

func (h *Human) Swim() {
	if h.CanSwim {
		fmt.Println("I can swim!")
	} else {
		fmt.Println("I cannot swim!")
	}
}
