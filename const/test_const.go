package main

/*
	常量和枚举iota
*/

import (
	"fmt"
)

const(
	// BJ = iota //iota第一行默认值是0
	// SH		//1
	// GZ		//2
	// SZ		//3

	// BJ = 10*iota //iota第一行默认值是0
	// SH			//10
	// GZ			//20
	// SZ			//30

	BJ,SH = iota + 1,iota + 2 //iota第一行默认值是0   1 2
	GZ,SZ			// 2 3
	HZ,TZ = iota * 2,iota * 3 //4 6
	BD,FS			//6 9
)

func main() {
	const length int = 100
	fmt.Println("length = ", length)
	fmt.Println("BJ = ", BJ)
	fmt.Println("SH = ", SH)
	fmt.Println("GZ = ", GZ)
	fmt.Println("SZ = ", SZ)
	fmt.Println("HZ = ", HZ)
	fmt.Println("TZ = ", TZ)
	fmt.Println("BD = ", BD)
	fmt.Println("FS = ", FS)
}