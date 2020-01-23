//在函数内使用其他函数内定义的变量
//无法通过编译，报错：undefined x

package main 

import "fmt"

func main(){
  x := 10
  fmt.Println(x)
  foo()
}

func foo(){
  //no acess to x 
  //this does not  compile
  fmt.Println(x)

}
