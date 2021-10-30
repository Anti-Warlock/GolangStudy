package main

import "fmt"

//类名首字母大写,表示其他包也能够访问
type Hero struct {
	//类的属性首字母大写,表示属性是能够被外部访问的,否则只能被内部访问
	Name  string
	Ad    int
	Level int
}

//方法名首字母大写表示是公共方法,否则是私有方法
func (this Hero) Show() {
	fmt.Printf("Name = %s,Ad = %d,Level = %d\n", this.Name, this.Ad, this.Level)
}
func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{
		Name:  "zs",
		Ad:    100,
		Level: 1,
	}

	hero.Show()

	//指针传递
	hero.Name = "ls"

	hero.Show()
}
