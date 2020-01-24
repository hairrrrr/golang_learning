package main 

import "fmt"

func main(){

  a := 1
  var b = &a
  *b++
  fmt.Println(a)

}
