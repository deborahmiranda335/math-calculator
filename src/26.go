package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Generate random numbers between 0 and 100
    const MAX = 100
    const N = 5

    var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
    var result []int

    for i := range result {
        result[i] = int(rng.Intn(MAX))
    }

    fmt.Println("Random numbers:", result)
}
