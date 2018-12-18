package main

import (
  "fmt"
)

func do(i interface{}){
  switch v:=i.(type){
  case int:
      fmt.Printf("i m an integer value: %d\n",v)
  case string:
      fmt.Printf("i m a string: %q\n",v)
  case float64:
      fmt.Printf("i m float value: %v\n",v)
  default:
      fmt.Printf("dont know about %T",v)
  }
}

func main(){
  do(21)
  do("hola madrid")
  do(3.14)
  do(true)
}
