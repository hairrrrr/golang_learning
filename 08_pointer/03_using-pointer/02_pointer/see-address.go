package main 

import "fmt"

func zero(z *int){
  
  fmt.Println(z)
  fmt.Printf("%p\n",z)

  *z = 0

}
func main(){

  x := 1
  fmt.Println(&x)
  fmt.Printf("%p\n",&x)

  zero(&x)

  fmt.Println(x)

}
