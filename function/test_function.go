package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Printf("a = %s,b = %d", a, b)
	c := 100
	return c
}

//返回多个返回值,匿名
func foo2(a string, b int) (int, int) {
	fmt.Printf("a = %s,b = %d", a, b)
	return 10, 50
}

//返回多个返回值,有形参名称
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Printf("a = %s,b = %d\n", a, b)
	//r1 r2属于方法的形参,初始化的默认值是0
	//r1 r2作用域是函数体
	fmt.Printf("r1 = %d,r2 = %d\n", r1, r2)

	r1 = 1000
	r2 = 2000
	return
}

//省略写法
func foo4(a, b int) (r1, r2 int) {
	fmt.Printf("a = %d,b = %d", a, b)
	r1 = 1000
	r2 = 2000
	return
}

func main() {
	c := foo1("a", 1)
	fmt.Printf("foo1 result = %d\n", c)
	d, e := foo2("a", 1)
	fmt.Printf("foo2 result = %d,%d\n", d, e)
	f, g := foo3("a", 1)
	fmt.Printf("foo3 result = %d,%d\n", f, g)
	h, i := foo4(0, 1)
	fmt.Printf("foo4 result = %d,%d\n", h, i)
}
