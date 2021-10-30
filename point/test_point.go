package main

import "fmt"

func changeValueWithValue(p int) {
	p = 10
}

func changeValueWithPoint(p *int) {
	//Step2 p的地址
	println(p)
	//Step3 p的值(地址)指向的值
	println(*p)
	//Step4 修改p的值(地址)指向的值
	*p = 10
}

func main() {
	/*
		值传递 a=1
	*/
	a := 1
	changeValueWithValue(a)
	fmt.Println(a)
	/* 
		引用传递 a=10
	*/
	//Step1 a的地址
	println(&a)
	changeValueWithPoint(&a)
	fmt.Println(a)
}