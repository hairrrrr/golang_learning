package main 
import "fmt"
func main(){
  //多个相同类型的变量初始化，可以用逗号把每个变量隔开，一起初始化
  //many variables of the same format can be initialized at once,each of which seperated by comma  
  var a,b string = "Hello", "World"
  var c,d,e int 
  var i,j float64
  var t,f bool

  //多个变量赋值，不管这些变量类型是否相同都可以用逗号隔开,一起赋值
  // many variables can be assigned at once,no matter they are of same format or not,each of which seperated by comma 
  c, d = 1, 2
  i, t = 3.14, true

  fmt.Printf("%v %v\n%v %v %v\n%v %v\n%v %v", a, b, c, d, e, i, j, t, f)
}
//output
//Hello World
//1 2 0
//3.14 0
//true false
