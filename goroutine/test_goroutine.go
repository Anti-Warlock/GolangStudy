package main

import (
	"fmt"
	"time"
)

//从goroutine
func newTask() {
	i := 0
	//死循环
	for {
		i++
		fmt.Printf("new GoRoutine:i=%d\n",i)
		time.Sleep(1 * time.Second)
	}
}

//主goroutine
func main() {
	//创建一个go程,去执行方法
	go newTask()
	
	//主go程退出后,整个进程会直接退出
	fmt.Println("main GoRoutine exit...")
	/* 
		i := 0
		for {
			i++
			fmt.Printf("main GoRoutine:i=%d\n",i)
			time.Sleep(1 * time.Second)
		}
	*/
}