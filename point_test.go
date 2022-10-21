package demo

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)                       //a =  1
	fmt.Println("&a = ", &a)                     //&a =  0xc042010098
	fmt.Println("*&a = ", *&a)                   //*&a =  1
	fmt.Println("b = ", b)                       //b =  0xc042010098
	fmt.Println("&b = ", &b)                     //&b =  0xc042004028
	fmt.Println("*&b = ", *&b)                   //*&b =  0xc042010098
	fmt.Println("*b = ", *b)                     //*b =  1
	fmt.Println("c = ", c)                       //c =  0xc042004028
	fmt.Println("*c = ", *c)                     //*c =  0xc042010098
	fmt.Println("&c = ", &c)                     //&c =  0xc042004030
	fmt.Println("*&c = ", *&c)                   //*&c =  0xc042004028
	fmt.Println("**c = ", **c)                   //**c =  1
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c) //***&*&*&*&c =  1
	fmt.Println("x = ", x)                       //x =  1
}
