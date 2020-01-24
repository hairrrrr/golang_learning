package main 

import "fmt"

func No_pointer(z int){

  fmt.Printf("%p\n",&z)
  fmt.Println(&z)

  z = 0

}

func main(){

  x := 1

  fmt.Printf("%p\n",&x);
  fmt.Println(&x)

  No_pointer(x)
  fmt.Println(x)
  
}
//output:
//0xc000090010
//0xc000090010
//0xc000090018
//0xc000090018
//1

