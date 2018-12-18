package main

import (
"fmt"
)

func main(){
q:=[]int{2,3,5,7,11,13}
r:=[]bool{true,true,false,false,true,false}
s:=[]struct{
  i int
  b bool
}{
  {2,true},
  {3,true},
  {5,false},
  {7,false},
  {11,true},
  {13,false},
}
fmt.Println(q)
fmt.Println(r)
fmt.Println(s)
  q=q[:0]
  printslice(q)
  q=q[:4]
  printslice(q)
  q=q[2:]
  printslice(q)
}
func printslice(q []int){
  fmt.Printf("len=%d cap=%d %v\n",len(q),cap(q),q)
}
