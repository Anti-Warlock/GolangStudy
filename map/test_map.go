package main

import "fmt"

func main() {

	/*
		第一种声明方式
	*/
	//声明一个map,key为int,value为string
	var myMap1 map[int]string
	//给map分配空间,长度为10,扩容机制类似于slice
	myMap1 = make(map[int]string, 10)
	myMap1[1] = "java"
	myMap1[2] = "c++"
	myMap1[3] = "python"

	fmt.Println(myMap1)

	/*
		第二种声明方式
	*/
	//给map分配空间,不指定长度,动态扩容
	myMap2 := make(map[int]string)
	myMap2[1] = "php"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	/*
		第三种声明方式,直接定义map内部的KV
	*/
	myMap3 := map[int]string{
		1: "java",
		2: "go",
		3: "python",
	}

	fmt.Println(myMap3)

	/*
		map的增删改查
	*/
	cityMap := make(map[string]string)
	cityMap["China"] = "BeiJing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	for key, value := range cityMap {
		fmt.Printf("key = %s,value = %s\n", key, value)
	}

	//删除
	delete(cityMap, "Japan")

	//修改
	cityMap["USA"] = "DC"
	fmt.Println(cityMap)

	cityMap["UK"] = "London"
	printMap(cityMap)
}

func printMap(cityMap map[string]string) {
	//cityMap是一个引用传递
	for key, value := range cityMap {
		fmt.Printf("key = %s,value = %s\n", key, value)
	}
}
