package main 

import "fmt"

func main(){

  for i := 1; i < 5; i++{
    for j := 1; j < 5; j++{
      fmt.Printf("%d - %d\t", i, j)
    }
    fmt.Println()
  }
}
