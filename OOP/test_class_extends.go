package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this Human) Eat() {
	fmt.Println("Human Eat...")
}

func (this Human) Walk() {
	fmt.Println("Human Walk...")
}

type SuperMan struct{
	Human //SuperMan继承了Human类的方法和属性
	level int
}

func (this SuperMan) Eat(){
	fmt.Println("SuperMan Eat...")
}

func (this SuperMan) Fly(){
	fmt.Println("SuperMan Fly...")
}

func main() {
	human := Human{"zs","male"}
	human.Eat()
	human.Walk()

	/*
		定义一个子类对象 
	*/
	//方法一
	superMan := SuperMan{Human{"ls","female"},1}
	//方法二
	// var superMan SuperMan
	// superMan.name = "ls"
	// superMan.sex = "female"
	// superMan.level = 1
	superMan.Eat() //子类重写后的方法
	superMan.Walk() //父类的方法
	superMan.Fly() //子类的新方法
}