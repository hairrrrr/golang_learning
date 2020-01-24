package main 

import "fmt"

func zero(z *int){
  *z = 0
}
func main(){

  x := 1
  zero(&x)
  fmt.Println(x)

}
