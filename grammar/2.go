package main 
import "fmt"
  var a,b int = 1,2
  var c,d = 3.14, "Hi"

  //这种写法一般只用于声明全局变量
  var (
    e int 
    f bool
  )
  //全局变量声明后可以不使用，但是函数内声明得变量必须要使用
func main(){
  //:= 是初始化得操作符，不能用来赋值已经声明得变量，必须要有新变量出现
  i,j := 123,"hello"

  //没有初始化得变量默认为0

  fmt.Println(a,b)
  fmt.Printf("%v %q\n",c,d)
  fmt.Print(e,f)
  fmt.Println(i,j)

  //输出
  //1 2
  //3.14 "Hi"
  //0 false123 hello
}
