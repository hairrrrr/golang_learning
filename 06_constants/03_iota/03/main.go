package main 

import "fmt"

const (
  _ = iota
  b = iota*10 
  c = iota*10
)
func main(){

  fmt.Println(b, c)

}
