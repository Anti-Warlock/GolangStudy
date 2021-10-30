package main

import "fmt"

func main() {
	//长度表示切片左指针到右指针之间的距离
	//容量表示切片左指针到底层数组末尾的距离
	var numbers = make([]int, 3, 5) //声明一个切片,长度为3,容量为5
	fmt.Printf("len = %d,cap = %d,numbers = %v\n", len(numbers), cap(numbers), numbers)

	//报错,超过长度了
	// numbers[3] = 3

	//通过append向numbers切片追加一个元素3,[0,0,0,1],len=4,cap=5
	numbers = append(numbers, 3)
	fmt.Printf("len = %d,cap = %d,numbers = %v\n", len(numbers), cap(numbers), numbers)

	//通过append向numbers切片追加一个元素4,[0,0,0,1],len=5,cap=5
	numbers = append(numbers, 4)
	fmt.Printf("len = %d,cap = %d,numbers = %v\n", len(numbers), cap(numbers), numbers)

	//向一个容量已经满了的slice继续追加元素5
	numbers = append(numbers, 5)
	/*
		底层会自动帮我们扩容slice原本容量大小的长度,翻倍
				len   cap   slice
		扩容前 	 5	   5     [0,0,0,3,4]
		扩容后	 6     10    [0,0,0,3,4,5]
	*/

	fmt.Printf("len = %d,cap = %d,numbers = %v\n", len(numbers), cap(numbers), numbers)

	println("******")

	//如果只声明了切片的长度,默认容量跟长度相等
	var numbers1 = make([]int, 3)
	fmt.Printf("len = %d,cap = %d,numbers = %v\n", len(numbers1), cap(numbers1), numbers1)
	numbers1 = append(numbers1, 3)
	fmt.Printf("len = %d,cap = %d,numbers = %v\n", len(numbers1), cap(numbers1), numbers1)

	//slice截取
	s := []int{1, 2, 3}
	//截取机制(类似Python),含头不含尾
	//1,2
	s1 := s[0:2]
	fmt.Println(s1)

	s1[0] = 100
	/*
		指向的是用一个数组,修改值后,两个都发生改变了
	*/
	//[100,2,3]
	fmt.Println(s)
	//[100,2]
	fmt.Println(s1)

	//s2 [0,0,0]
	s2 := make([]int, 3)
	//copy将底层数组的slice一起拷贝,深拷贝
	//将s中的值依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)

	/*
		再次尝试修改,因为指向的是两个不同的slice,所以值没有互相影响
	*/
	s2[0] = 50
	fmt.Println(s)
	fmt.Println(s2)
}
