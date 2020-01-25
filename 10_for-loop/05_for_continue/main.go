package main 

import "fmt"

func main(){

  i := 1
  for{
    if i % 2 == 0{
      i++
      continue
    } else{
      fmt.Println(i)
      i++
    }
    if i >= 10{
      break;
    }
  }

}
