package main

import (
  "fmt"
)

func main(){
  pow:=make([]int,12)
//using of index and value in for loop  
  for i:=range pow{
     pow[i]=1<<uint(i)
  }
  for _,value:=range pow{
     fmt.Printf("%d\n",value)
  }
}
