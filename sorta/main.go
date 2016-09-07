// Package main provides
// Encapsulation is at package level
package main

import (
	"fmt"
	"github.com/gambledor/golang/sorta/composition"
	"github.com/gambledor/golang/sorta/polimorphism"
)

// a struct is a value
type SomeStruct struct {
	Field string
}

// types are associated with method set
func (s *SomeStruct) foo(field string) {
	s.Field = field
}

func main() {
	value := SomeStruct{Field: "Structs are values"}
	copy := value
	copy.Field = "This is a Copy & does not change the variable value"

	fmt.Println(value.Field)
	fmt.Println(copy.Field)
	//------------------------------------------------------------
	someStruct := new(SomeStruct)
	someStruct.foo("a")
	fmt.Println(someStruct.Field) // "a"

	someStruct.foo("b")
	fmt.Println(someStruct.Field) // "b"
	//------------------------------------------------------------
	hillary := new(polimorphism.Hillary)
	hillary.Slogan() // "Stronger together."
	// polymorphism.SaySlogan(hillary) // "Stronger together."

	trump := new(polimorphism.Trump)
	polimorphism.SaySlogan(trump) // "Make America great again."
	//------------------------------------------------------------
	// amy is composed of type Human
	amy := composition.Amy{
		Human: composition.Human{
			FirstName: "Amy",
			LastName:  "Chen",
			CanSwim:   true,
		},
	}
	// alan is composed of type Human
	alan := composition.Alan{
		Human: composition.Human{
			FirstName: "Alan",
			LastName:  "Chen",
			CanSwim:   false,
		},
	}

	// Human's method set are forwarded to type Amy
	amy.Name()
	amy.Swim()

	alan.Name()
	alan.Swim()
}
