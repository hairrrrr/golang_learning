package main 

import "fmt"

func main(){

  x := 13 % 2
  if x == 1{
    fmt.Println(x, "Odd")
  }else{ 
      fmt.Println("Even")
    }
      for i := 1; i<=100; i++{
        if i % 2 == 1{
          fmt.Println(i, "Odd")
        }else{
          fmt.Println(i, "Even")
        }
    }
}
