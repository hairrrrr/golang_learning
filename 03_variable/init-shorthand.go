package main 
import "fmt"
func main(){

  //这种初始化只能写在函数体内部
  //you can only do this inside a func
  a := "Hello World"
  b, c, d  := 1, false, 3.14
  e := 2
  f := true 

  fmt.Println(a, b, c, d, e, f)
}
//output
//Hello World 1 false 3.14 2 true
