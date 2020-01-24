package main 

import "fmt"

func main(){
  
  a := 1 

  fmt.Println(a)
  fmt.Println(&a)

  var b = &a 

  fmt.Println(b)
  fmt.Println(*b)

}
