package main 

import "fmt"

func main(){

  fmt.Println(x)
  fmt.Println(y)
  x := 0
}

var y = 1
//error:undefined x

//global var can be declared anywhere but local var must be declared before you use it
//全局变量可以声明在任意位置，局部变量声明必须在使用之前
