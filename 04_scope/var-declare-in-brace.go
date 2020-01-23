//使用在大括号内声明的变量也是非法的

package main 

import "fmt"

func main(){

  x := 10
  {
    fmt.Println(x)
    y := "Hello World"
    fmt.Println(y)
  }
  //undefined y
  fmt.Println(y)

}
