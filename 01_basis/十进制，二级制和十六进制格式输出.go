//输出十进制，二进制与十六进制
//output format(decimal,binary,hexadecimal)
package main 

import "fmt"

func main(){
  //decimal 十进制
  fmt.Println(42)
  fmt.Printf("%d\n",42)
  //binary  二进制
  fmt.Printf("%d - %b\n",42,42)
  //hexadecimal 十六进制
  fmt.Printf("%x\n%#x\n%#X\n",42,42,42)
}
//42
//42
//42 - 101010
//2a
//0x2a
//0X2A
