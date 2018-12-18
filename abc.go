package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main(){
  fmt.Println("welcome")
  fmt.Println("The time is",time.Now())
  fmt.Println("number",rand.Intn(11))
}
