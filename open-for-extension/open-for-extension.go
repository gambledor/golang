//dave.cheney.net
package main

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Legs() int { return 4 }

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs\n", c.Legs())
}

type OctoCat struct {
	Cat
}

func (o OctoCat) Legs() int { return 5 }

func main() {
	var octo OctoCat
	fmt.Println(octo.Legs()) // 5
	octo.PrintLegs() // I have 5 legs
}

// This is because PrintLegs is defined on Cat type
// It takes a Cat as its receiver, and so it dispatches
// to Cat's Legs method. Cat has no knowledge of the type
// it has been embedded into, so its methos set cannot be altered by embedding.
// Thus we say Go's types are open for extension and closed for modification.

