package main 

import "fmt"

func max(x int) int {

  return x + 1

}

func main(){

  max := max(1)
  fmt.Println(max)

}
//output:
//2

//don't do this bad coding practise
//不要这么写，变量名和函数名不要相同
