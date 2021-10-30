package main

import "fmt"

func main(){
	//定义一个无缓冲的chanel
	c := make(chan int)

	go func ()  {
		defer fmt.Println("goroutine结束...")
		fmt.Println("goroutine正在运行")
		//将66写到chanel中
		c <- 66
	}()

	//从管道c中接收数据,并赋值给num
	num := <-c

	fmt.Println("num = ", num)
	fmt.Println("main goroutine结束")

	/* 
		channel会起到一个同步的作用
		主go程和副go程需要同时达到从管道写/读值的代码时,才不会发生阻塞,否则当一方代码没有到达时,会发生阻塞,等待另一方读/写完成
	*/
}