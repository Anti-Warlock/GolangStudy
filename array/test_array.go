package main

import "fmt"

func main() {
	//固定长度的数组,默认值都是0
	var arrOne [10]int

	arrTwo := [10]int{1, 2, 3, 4}
	arrThree := [4]int{5,6,7,8}

	//slice动态数组
	arrFour := []int{1,2,3,4}

	/*
		遍历数组
	*/
	//方法一
	for i := 0; i < 10; i++ {
		println(arrOne[i])
	}
	// //方法二
	// for i := 0; i < len(arrOne); i++ {
	// 	println(arrOne[i])
	// }
	//方法三,如果不想用index可以用_代替(匿名)
	for index, value := range arrTwo {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}

	//查看数组类型
	//不同长度的数组都是不同的数据类型,不能作为公共数据类型
	//[10]int 和 [4]int 是两种不同的数据类型
	fmt.Printf("arr1 type = %T\n",arrOne)
	fmt.Printf("arr2 type = %T\n",arrTwo)
	fmt.Printf("arr3 type = %T\n",arrThree)

	printArray(arrThree)
	// printArray(arrTwo)//报错
	println(arrThree[1])//值不会改变

	println("***slice动态数组***")
	fmt.Printf("arr4 type = %T\n",arrFour)
	printArr(arrFour)

	//修改数组的值后
	println(arrFour[1])
}

//定长数组传参
func printArray(arr [4]int){
	//值传递
	for index, value := range arr {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}
	arr[1] = 66	
}

//slice动态数组传参
func printArr(arr []int){
	//引用传递
	//_表示匿名
	for _, value := range arr {
		fmt.Printf("value=%d\n", value)
	}
	arr[1] = 66	
}