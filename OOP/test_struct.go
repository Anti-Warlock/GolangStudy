package main

import "fmt"

//声明一种新的数据类型myInt
type myInt int

//定义一个结构体(类似于Java对象)
type Book struct{
	title string
	author string
}

func changeBook(book Book){
	//值传递
	book.author = "aa"
}

func changeBookOne(book *Book){
	//引用传递
	book.author = "aa"
}

func main() {

	// var a myInt = 10
	// fmt.Printf("a = %d,type = %T\n",a,a)

	var myBook Book
	myBook.title = "GoLang"
	myBook.author = "zs"

	fmt.Printf("%v\n",myBook)
	
	//值传递
	changeBook(myBook)
	
	fmt.Printf("%v\n",myBook)
	
	//引用传递
	changeBookOne(&myBook)
	
	fmt.Printf("%v\n",myBook)
}