package main

import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan int)

	//子go程
	go func() {
		for i := 0; i < 6; i++ {
			//当c中无数据中,会发生阻塞
			//当api中的channel开始向c中写入数据后,子go程开始从c中读取数据
			fmt.Println(<-c)
			fmt.Println("从c中读取数据...")
		}
		//从c中读取数据完毕后,会向quit中写入0
		quit <- 0
	}()

	//主go程
	fibonacii(c, quit)
}

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		//select可以监控多个channel的状态
		select {
		//如果能向channel c中写入数据,则执行该case语句
		case c <- x:
			x, y = x+y, x
		//如果能从channel quit中读取数据,则执行该case语句
		case <-quit:
			fmt.Println("quit...")
			return
		//如果上面都没有成功,则执行default
		default:
			break
		}
	}
}
