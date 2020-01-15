package main 
import "fmt"
func main(){

  var a int
  var b string
  var c float64
  var d bool

  //声明变量但不进行初始化,直接输出
  //declare the variables but not initialize them,print directly
  fmt.Printf("%v \n", a)
  fmt.Printf("%v \n", b)
  fmt.Printf("%v \n", c)
  fmt.Printf("%v \n", d)

  fmt.Println()

  //可以看到，不给变量赋值直接输出
  //int,float型会默认为0
  //bool型会默认为false
  //string型不进行任何输出
  //we can see,
  //int and float format's default is 0
  //bool format's default is false 
  //string format print nothing
}
//output:
//0 
// 
//0 
//false 
//
