package main

import (
  "fmt"
  "math"
)
 type vertex struct{
   X,Y float64
 }

 func(v vertex) abs() float64{
   return math.Sqrt(v.X*v.X +v.Y*v.Y)
 }
 func main(){
   fmt.Println(vertex{3,4}.abs())
 }
