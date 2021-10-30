package main

// _ "GolangStudy/init/lib1" //匿名导入,只执行init方法,不执行特定API
// mylib2 "GolangStudy/init/lib2" //别名导入
// . "GolangStudy/init/lib2" //.导入,不建议使用，容易产生歧义

func main() {
	//方法首字母大写表示公共方法可以被引用,小写表示只能包内调用
	// lib1.Lib1Test()
	// lib2.Lib2Test()

	/*
		lib1 init
		lib2 init
		lib1 test
		lib2 test
	*/

	// lib1.Lib1Test()
	// mylib2.Lib2Test()

	// Lib2Test()
}
