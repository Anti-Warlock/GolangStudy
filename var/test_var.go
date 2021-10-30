package main

import (
	"fmt"
)

/*
	四种命名变量的方式
*/

//方式四测试
// e:=20
// f:="blog"

func main() {
	//方式一,默认值为0：
	var a int
	fmt.Println("a = ", a)
	//方式二：
	var b int = 100
	fmt.Println("b = ", b)
	//方式三,省略类型,字段匹配：
	var c = 10
	var d = "abcd"
	fmt.Printf("c = %d,type of c = %T\n", c, c)
	fmt.Printf("d = %s,type of d = %T\n", d, d)
	//方式四,省略var和类型,自动匹配,只能用做局部变量
	e := 20
	f := "blog"
	fmt.Printf("e = %d,type of e = %T\n", e, e)
	fmt.Printf("f = %s,type of f = %T\n", f, f)

}
