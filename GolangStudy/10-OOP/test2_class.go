package main

import (
	"fmt"
)
/*----------------------------------------定义一个类----------------------------------------*/
// 先定义一个结构体，首字母大写表示该类在其他包也可以访问，小写则只能在本包访问
type Hero struct {
	// 定义属性，首字母大写，表示该属性在其他包也可以访问，小写则只能在本包访问
	Name string
	AD int
	Level int
}

// 然后定义该类的方法，并用 this *类名 绑定到结构体，方法名首字母大写，表示该方法在其他包也可以访问，小写则只能在本包访问

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name, "AD = ", this.AD, "Level = ", this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

// 注意，如果用 this， 不用this * 的话，是值传递，无法修改
func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	// 创建一个对象
	hero := Hero{Name: "Zhangsan", AD: 100, Level: 1}

	hero.Show()

	hero.SetName("Lisi")

	hero.Show()
}