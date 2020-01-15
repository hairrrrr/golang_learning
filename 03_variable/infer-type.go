package main 
import "fmt"
func main(){
  //初始化变量时，可以省去变量类型,不管变量的类型是否相同
  //When initializing a variable or many variables, you can leave out the variable type,no matter what kind of the formats the variables are 
  var a = "Hello world"
  var b, c, d = 1, 2.2, false

  fmt.Println(a, b, c, d)
}
//output
//Hello world 1 2.2 false
