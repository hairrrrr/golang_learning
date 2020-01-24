package main 

import "fmt"

const (
  _ = iota 
  KB = 1 << (iota * 10)
  MB = 1 << (iota * 10)
  GB = 1 << (iota * 10)
  TB = 1 << (iota * 10)

)

func main(){

  fmt.Println("binary\t\t\t\t\t\tdecimal")
  fmt.Printf("%-41b\t%d\n",KB,KB)
  fmt.Printf("%-41b\t%d\n",MB,MB)  
  fmt.Printf("%-41b\t%d\n",GB,GB) 
  fmt.Printf("%-41b\t%d\n",TB,TB) 

}
//output:
//binary            decimal
//10000000000                               1024
//100000000000000000000                     1048576
//1000000000000000000000000000000           1073741824
//10000000000000000000000000000000000000000 1099511627776

