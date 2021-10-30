package main

import "fmt"

func main() {
	c := make(chan int)
	//向一个为nil(即没有make)的channel发送数据,无论如何都会阻塞 deadlock
	// var c chan int
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			//向一个已经关闭的channel再发送数据,会出现Panic错误
			// close(c)
		}

		close(c)
		//关闭channel后,依然可以从channel中获取数据
	}()

	/*
		for {
			//ok为true表示channel没有关闭,为false则表示已经关闭
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/

	//可以使用range来迭代不断操作channel
	for data := range c {
		/*
			channel中有数据,就会进入for循环读取
		*/
		fmt.Println(data)
	}
	fmt.Println("Finished...")
}
