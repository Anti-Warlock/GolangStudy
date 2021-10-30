package main

import (
	"fmt"
)

func main() {
	//声明一个切片,并且初始化,默认值是1,2,3,长度是3
	// slice1 := []int{1, 2, 3}

	//声明slice2是一个切片,但是并没有给slice分配空间
	var slice1 []int
	//给slice1分配空间,长度为3,默认值为0
	// slice1 = make([]int, 3)
	// slice1[0] = 100
	// slice1[0] = 1//报错,因为没有给slice1开辟空间

	//声明slice1是一个切片,同时分配空间,初始化值是0
	// var slice1 = make([]int,3)

	//简写版,常用
	//声明slice1是一个切片,同时分配空间,初始化值是0,通过:=推断slice1是一个切片
	// slice1 := make([]int,3)

	fmt.Printf("len = %d,slice = %v\n", len(slice1), slice1)

	//判断一个slice是否为0,没有元素(是否分配了空间)
	if slice1 == nil {
		println("slice1是一个空切片")
	} else {
		println("slice1不是一个空切片")
	}
}
