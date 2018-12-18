package main

import (
 "fmt"
 "time"
)

func main(){
  fmt.Println("When's Monday?")
  today:=time.Now().Weekday()
  switch time.Monday {
  case today + 0:
  fmt.Println("today")
  case today + 1:
  fmt.Println("Tomorrow")
  case today + 2:
  fmt.Println("In two days")
  default:
  fmt.Println("too far away")
  }

}
