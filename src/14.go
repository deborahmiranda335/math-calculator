  package main
  
  import (
      "math/rand"
      "time"
  )
  
  func GenerateRandom() int {
      rand.Seed(time.Now().UnixNano())
      return rand.Intn(10)
  }