package main

import (
  "fmt"
)

type Vertex struct{
  Lat,Long float64
}

var m=map[string]Vertex{
  "Bell labs":Vertex{
    30.435,-72.5634,
  },
  "Google":Vertex{
    54.689,-73.321,
  },
}

func main(){
  fmt.Println(m)
}
