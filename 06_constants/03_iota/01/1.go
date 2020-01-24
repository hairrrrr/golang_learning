package main 

import "fmt"

const(
  a = iota;
  b = iota;
  c = iota;
)
func main(){

  fmt.Println(a, b,c);
}
//output: 
//0 1 2 
