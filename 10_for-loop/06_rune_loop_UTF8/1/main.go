package main 

import "fmt"

func main(){

  for i := 250; i<= 340; i++{
    fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))
  }
  foo := "a"
  fmt.Println(foo)
  fmt.Printf("%T\n", foo)

}
