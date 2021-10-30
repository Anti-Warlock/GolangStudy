package main

import "fmt"

//本质是一个指针,是一个接口类型
type AnimalIF interface{
	Sleep()
	GetColor() string //获取颜色
	GetType() string //获取类型
}
//具体的类
type Cat struct{
	//实现接口的三个方法就相当于实现了接口，不需要把接口类型定义到具体的类中 AnimalIF
	//相当于interface指针指向了Cat对象
	color string //猫的颜色
}

/* 
	如果没有实现全部的方法，则表示没有完全实现接口，接口的指针无法指向具体的类！！！
*/

func (this Cat) Sleep(){
	fmt.Println("Cat is Sleeping...")
}
func (this Cat) GetColor() string{
	return this.color
}
func (this Cat) GetType() string{
	return "Cat"
}

//具体的类
type Dog struct{
	color string
}
func (this Dog) Sleep(){
	fmt.Println("Dog is Sleeping...")
}
func (this Dog) GetColor() string{
	return this.color
}
func (this Dog) GetType() string{
	return "Dog"
}

func showAnimal(animal AnimalIF){
	//根据不同的子类调用不同子类的方法,多态
	animal.Sleep()
	fmt.Println("animal color is",animal.GetColor())
	fmt.Println("animal type is",animal.GetType())
}

func main(){

	// var animal AnimalIF //接口的数据类型,父类指针
	// animal = Cat{"Black"}
	// animal.Sleep()
	// fmt.Println("animal color is",animal.GetColor())
	// fmt.Println("animal type is",animal.GetType())
	// animal = Dog{"Yellow"}
	// animal.Sleep()
	// fmt.Println("animal color is",animal.GetColor())
	// fmt.Println("animal type is",animal.GetType())


	cat := Cat{"Black"}
	dog := Dog{"Yellow"}
	showAnimal(cat)
	showAnimal(dog)
}