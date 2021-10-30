package main

import (
	"fmt"
	_ "runtime"
	"time"
)

func main() {

	/*
		//用go程创建一个形参和返回值都为空的函数,
		go func() {
			defer fmt.Println("A.defer")

			// return //A.defer

			//匿名函数
			func ()  {
				defer fmt.Println("B.defer")
				//此处return只退出了当前函数,并没有退出当前goroutine
				// return //B.defer A A.defer

				//退出当前goroutine
				runtime.Goexit()
				fmt.Println("B")
			}()//括号表示调用这个函数
			fmt.Println("A")
		}()
	*/

	//用go创建一个带参数和返回值的函数
	go func(a, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
