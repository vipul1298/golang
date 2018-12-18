package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64{
  for z:=1.0;z<10 && z<=math.Sqrt(x);z++{
  z-=(z*z-x)/(2*z)
  }

  return z
  }

func main(){
  fmt.Println(Sqrt(4))
}
