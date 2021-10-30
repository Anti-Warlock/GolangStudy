package main

import "fmt"

//interface{}表示万能数据类型,类似于java里面的object
func myFunc(arg interface{}) {
	fmt.Println("myFunc is Called...")
	fmt.Println(arg)

	//interface{}如何判断引用的数据类型到底是什么
	//类型断言机制
	value,ok := arg.(string)
	if !ok{
		fmt.Println("value is not string")
	}else{
		fmt.Println("value is string")
		fmt.Printf("value type is %T\n",value)
	}
}

type Book struct{
	auth string
}

func main() {
	book := Book{"GoLang"}
	myFunc(book)
	myFunc("abc")
	myFunc(1)
	myFunc(3.14)
}