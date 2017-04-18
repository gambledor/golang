package main

import "fmt"

// define a function type
type testInt func(int) bool

func isOdd(integer int) bool {
	return integer%2 == 0
}

func isEven(integer int) bool {
	return !isOdd(integer)
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("slice = ", slice)

	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice: ", odd)

	even := filter(slice, isEven)
	fmt.Println("Even elements of slice: ", even)
}
