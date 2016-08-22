// Package polimorphism
package polimorphism

import "fmt"

type SloganSayer interface {
	Slogan()
}

// SaySlogan takes in a parameter of type SloganSayer
func SaySlogan(sloganSayer SloganSayer) {
	sloganSayer.Slogan()
}

// Hillary implicitly satisfies the SloganSayer interface
// by implementing the function Slogan.
// Thus, Hillary is also type SloganSayer.
type Hillary struct {}

func (h *Hillary) Slogan() {
	fmt.Println("Stronger together.")
}

// same idea for Trump
type Trump struct { }

func (t *Trump) Slogan() {
	fmt.Println("Make America great again.")
}
