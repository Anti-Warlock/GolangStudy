package main

import (
	"fmt"
	"time"
)

/*
	当channel没有关闭时：
		1.当channel已经满了,再向里面写数据,会发生阻塞
		2.当channel为空时,再从里面读数据,会发生阻塞
*/

func main() {
	//声明一个容量为3的缓冲channel
	c := make(chan int, 3)

	fmt.Println("长度 = ", len(c), ",容量 = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		/*
			当i取4时,会发生阻塞,只有当主go程取出元素后,阻塞才会停止
		*/
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("当前元素", i, "长度 = ", len(c), ",容量 = ", cap(c))
		}
		// for i := 0; i < 3; i++ {
		// 	c <- i
		// 	fmt.Println("当前元素", i, "长度 = ", len(c), ",容量 = ", cap(c))
		// }
	}()

	time.Sleep(2 * time.Second)

	// for i := 0; i < 3; i++ {
	// 	num := <-c
	// 	fmt.Println("num = ", num)
	// }

	for i := 0; i < 4; i++ {
		//channel中没有内容,会发生阻塞
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("主go程结束")
}

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     //创建一个有缓存的channel
//     ch := make(chan int, 3)

//     //len(ch)缓冲区剩余数据个数， cap(ch)缓冲区大小
//     fmt.Printf("len(ch) = %d, cap(ch)= %d\n", len(ch), cap(ch))

//     //新建协程
//     go func() {
// 		defer fmt.Println("子go程结束...")
//         for i := 0; i < 4; i++ {
//             ch <- i //往chan写内容
//             fmt.Printf("子协程[%d]: len(ch) = %d, cap(ch)= %d\n", i, len(ch), cap(ch))
//         }
//     }()

//     //延时
//     time.Sleep(2 * time.Second)

//     for i := 0; i < 4; i++ {
//         num := <-ch //读管道中内容，没有内容前，阻塞
//         fmt.Println("num = ", num)
//     }

// 	fmt.Println("主go程结束...")
// }
