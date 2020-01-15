//用%v打印表变量的实例
//use format of %v to print variables- example
package main 
import "fmt"
func main(){
  a := 10 
  b := "golang"
  c := 4.17 
  d := true
  e := "Hello"
  g := 'M'
  h := "M"
  i := `M`
  f := `Do you like my hat?`
  //可以看出
  //1.用%v 打印字符串有两种声明格式：""和``
  //2.想打出单个字符 声明时可以使用：``或""，使用''不会打出单个字母，打出的是字母对应的utf-8的数字值
  //we can see:
  //1.if you want to print string,when you initialize the variables you can use "" or ``
  //2.if you want to print a single character,you can use `` or "",but do not use ''.in  situatuon of using '',you will get the value of utf-8
  fmt.Printf("%v \n",a)
  fmt.Printf("%v \n",b)
  fmt.Printf("%v \n",c)
  fmt.Printf("%v \n",d)
  fmt.Printf("%v \n",e)
  fmt.Printf("%v \n",f)
  fmt.Printf("%v \n",g)
  fmt.Printf("%v \n",h)
  fmt.Printf("%v \n",i)
}
//output
//10 
//golang 
//4.17 
//true 
//Hello 
//Do you like my hat? 
//77 
//M 
//M 
