//调用函数多次改变全局变量

package main

import "fmt"

var x = 0

func increment() int{

  x++
  return x

}

func main(){

  fmt.Println(increment())
  fmt.Println(increment())

}
//output:
//1
//2
