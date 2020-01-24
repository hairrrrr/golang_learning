package main 

import "fmt"

const (
  a = iota 
  b = 3.14 
  c = iota 
  d

)

func main(){

  fmt.Println(a, b, c, d)

}
//output
//0 3.14 2 3
