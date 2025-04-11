// Go Code 101: Simple Calculator

package main

import (
    "fmt"
)

func main() {
    var num1 int
    var num2 int
    var operation string = ""

    fmt.Println("Enter first number:")
    _, err := fmt.Scanln(&num1)
    if err != nil {
        panic(err)
    }

    fmt.Println("Enter second number:")
    _, err = fmt.Scanln(&num2)
    if err != nil {
        panic(err)
    }

    fmt.Println("Enter operation ( + - * / ):")
    _, err = fmt.Scanln(&operation)
    if err != nil {
        panic(err)
    }

    switch operation {
    case "+":
        result := num1 + num2
        fmt.Printf("%d + %d is: %d\n", num1, num2, result)
    case "-":
        result := num1 - num2
        fmt.Printf("%d - %d is: %d\n", num1, num2, result)
    case "*":
        result := num1 * num2
        fmt.Printf("%d * %d is: %d\n", num1, num2, result)
    case "/":
        if num2 != 0 {
            result := num1 / num2
            fmt.Printf("%d / %d is: %.2f\n", num1, num2, result)
        } else {
            fmt.Println("Error: Division by zero")
        }
    default:
        fmt.Println("Invalid operation")
}

