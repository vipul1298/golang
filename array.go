package main

import (
"fmt"
)

func main(){
 primes:=[6]int{2,3,5,7,11,13}
 var s []int=primes[1:4]
 fmt.Println(primes)
 fmt.Println(s)
 x:=12
 for i:=0;i<6;i++{
   if primes[i]==x{
   fmt.Println(i)
   break
   }else{
     fmt.Println("error")
     break
   }
 }

}
