//用for循环输出1~100的二进制和16进制
//output 1~100's binary and hexadecimal in for loop
package main 

import "fmt"

func main(){
  for i := 1; i <= 100; i++{
    fmt.Printf("%d\t%b\t%#X\n",i,i,i)
  } 
}
//out put
//g1 1 0X1
//g2 10  0X2
//3 11  0X3
//4 100 0X4
//5 101 0X5
//6 110 0X6
//7 111 0X7
//8 1000  0X8
//9 1001  0X9
//10  1010  0XA
//11  1011  0XB
//12  1100  0XC
//13  1101  0XD
//14  1110  0XE
//15  1111  0XF
//16  10000 0X10
//17  10001 0X11
//18  10010 0X12
//19  10011 0X13
//20  10100 0X14
//21  10101 0X15
//22  10110 0X16
//23  10111 0X17
//24  11000 0X18
//25  11001 0X19
//26  11010 0X1A
//27  11011 0X1B
//28  11100 0X1C
//29  11101 0X1D
//30  11110 0X1E
//31  11111 0X1F
//32  100000  0X20
//33  100001  0X21
//34  100010  0X22
//35  100011  0X23
//36  100100  0X24
//37  100101  0X25
//38  100110  0X26
//39  100111  0X27
//40  101000  0X28
//41  101001  0X29
//42  101010  0X2A
//43  101011  0X2B
//44  101100  0X2C
//45  101101  0X2D
//46  101110  0X2E
//47  101111  0X2F
//48  110000  0X30
//49  110001  0X31
//50  110010  0X32
//51  110011  0X33
//52  110100  0X34
//53  110101  0X35
//54  110110  0X36
//55  110111  0X37
//56  111000  0X38
//57  111001  0X39
//58  111010  0X3A
//59  111011  0X3B
//60  111100  0X3C
//61  111101  0X3D
//62  111110  0X3E
//63  111111  0X3F
//64  1000000 0X40
//65  1000001 0X41
//66  1000010 0X42
//67  1000011 0X43
//68  1000100 0X44
//69  1000101 0X45
//70  1000110 0X46
//71  1000111 0X47
//72  1001000 0X48
//73  1001001 0X49
//74  1001010 0X4A
//75  1001011 0X4B
//76  1001100 0X4C
//77  1001101 0X4D
//78  1001110 0X4E
//79  1001111 0X4F
//80  1010000 0X50
//81  1010001 0X51
//82  1010010 0X52
//83  1010011 0X53
//84  1010100 0X54
//85  1010101 0X55
//86  1010110 0X56
//87  1010111 0X57
//88  1011000 0X58
//89  1011001 0X59
//90  1011010 0X5A
//91  1011011 0X5B
//92  1011100 0X5C
//93  1011101 0X5D
//94  1011110 0X5E
//95  1011111 0X5F
//96  1100000 0X60
//97  1100001 0X61
//98  1100010 0X62
//99  1100011 0X63 
//100 1100100 0X64
