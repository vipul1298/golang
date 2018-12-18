package main

import (
  "fmt"
)

type student struct{
  name string
  age int
  roll_no int
  branch string
}

func (s student) String() string{
  return fmt.Sprintf("%v is in %v whose roll_no is %v and age is %v",s.name,s.branch,s.roll_no,s.age)
}
func main(){
  a:=student{"ayush",19,14,"CSE"}
  fmt.Println(a)
}
