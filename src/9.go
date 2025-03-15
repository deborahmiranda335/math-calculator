package main

import (
	"fmt"
	"math"
)

func main() {
	// Calculate the hypotenuse of a right triangle
	a := 5.0
	b := 6.0

	hypotenuse := math.Sqrt(a*a + b*b)

	fmt.Println("The hypotenuse is", hypotenuse)
}
