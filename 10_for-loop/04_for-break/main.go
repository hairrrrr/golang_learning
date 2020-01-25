package main 

import "fmt"

func main(){

  i := 1
  for ; ; i++{
    if i >= 10{
      break;
    } 
    fmt.Printf("%d ",i)
  } 

}
