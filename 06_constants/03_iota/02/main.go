package main 

import "fmt"

const (
  a = iota
  b
  c
)

const (
  d = iota 
  e  
  f
)

func main(){

  fmt.Println(a, b, c)
  fmt.Println(d, e, f)

}
