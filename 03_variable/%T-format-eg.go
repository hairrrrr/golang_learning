//用%T来输出变量类型
//use %T format to print types of variables
package main

import "fmt"

func main() {

     a := 10
     b := "golang"
     c := 4.17
     d := true 
     e := "Hello"
     f := `Do you like my hat?`
     g := 'M'
     h := `M`

     fmt.Printf("%T \n", a)
     fmt.Printf("%T \n", b)
     fmt.Printf("%T \n", c)
     fmt.Printf("%T \n", d)
     fmt.Printf("%T \n", e)
     fmt.Printf("%T \n", f)
     fmt.Printf("%T \n", g)
     fmt.Printf("%T \n", h)
}
//output:
//int 
//string 
//float64 
//bool 
//string 
//string 
//int32 
//string
