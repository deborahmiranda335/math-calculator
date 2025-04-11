package main

import (
	"fmt"
)

func main() {
	var s1, s2, result int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&s1, &s2)
	result = s1 + s2
	fmt.Println("Sum of", s1, "and", s2, "is", result)
}
