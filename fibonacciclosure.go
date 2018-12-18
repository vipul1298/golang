package main

import (
	"fmt"
)
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last1:=0
	last2:=1
	return func() int{
	  temp:=last1 + last2
		last2,last1=last1,temp
	 return last1-last2
}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
