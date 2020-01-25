package main 

import "fmt"

func main(){
  a := "1"
  b := 49
  fmt.Printf("%v - %d\n", a, []byte(a)) 
  fmt.Printf("%d - %v - %v\n", b, b, string(b))
  fmt.Printf("%d - %v\n", []byte("1"), []byte("1"))
}
