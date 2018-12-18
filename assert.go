package main

import "fmt"

func main(){
   var i interface{}="Hello"

   s:=i.(string)
   fmt.Println(s)

   s,ok:=i.(string)  //assertion succeeded -check
   fmt.Println(s,ok)

   t,ok:=i.(float64)
   fmt.Println(t,ok)

   t=i.(float64)
   fmt.Println(t) //panic
}
