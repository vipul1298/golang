package main

import (
"fmt"
)

func main(){
  fmt.Println("square of a number in reverse order")
  for i:=1;i<10;i++{
    defer fmt.Println(i*i)
  }

}
