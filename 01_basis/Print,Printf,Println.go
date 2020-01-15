//三种输出函数的简单区别
//Print,Printf,Println-simple difference
package main 
import "fmt"
func main(){
  var a int = 10
  var b bool = true
  var c string = "hi!"
  var d float32 = 3.1415926
 //这样输出是可以的，但是不能输入换行符，他们是连在一起的
  fmt.Print(a, b, c, d)

  fmt.Println(a)
  fmt.Println(b)
 //函数体内定义的变量必须都要使用
  fmt.Println(c)
   fmt.Println(d)
  //用Println是可以自动换行的
  
  //这样是不可以的
  //fmt.Println("%v\n%v\n%q\n%v",a,b,c,d)
  //fmt.Print("%v\n%v\n%q\n%v\n",a,b,c,d)

  //要想打出格式，这里得用Printf
  fmt.Printf("%v %v %q %v",a,b,c,d)
}
//output:
//10 truehi!3.141592510
//true
//hi!
//3.1415925
//10 true "hi!" 3.1415925



