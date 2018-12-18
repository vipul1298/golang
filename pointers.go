package main

import (
 "fmt"
)

func main(){
  i,j:=3,4
  p:= &i
  fmt.Println(*p)
  *p=32
  fmt.Println(i)
  p=&j
  *p=(*p)*2
  fmt.Println(*p)


}
