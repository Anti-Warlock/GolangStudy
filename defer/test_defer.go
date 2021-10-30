package main



func deferFunc() int{
	println("defer func called...")
	return 0
}

func returnFunc() int{
	println("return func called...")
	return 0
}

func deferAndReturn() int{
	defer deferFunc()
	return returnFunc()
}

func main() {
	//单个defer,类似于java的finally
	defer println("end1")
	defer println("end2")
	/*
		多个defer,遵循栈先入后出的特性,先执行end2,再执行end1
	*/

	println("hello go 1!")
	println("hello go 2!")

	/* 
		defer和return的执行顺序
		return先执行
		defer后执行
	*/
	deferAndReturn()
}
