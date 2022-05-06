package main

import (
	"fmt"
)

// 定义父类 Human
type Human struct {
	Name string
	Sex string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

// 定义子类 Superman
type Superman struct {
	Human // 把类名写在这儿，表示继承父类 Human 的属性

	Level int // 子类自己的属性
}

// 重定义父类中的方法
func (this *Superman) Eat() {
	fmt.Println("Superman.Eat()...")
}

// 子类自己的方法
func (this *Superman) Fly() {
	fmt.Println("Superman.Fly()...")
}
func (this *Superman) PrintSuperman() {
	fmt.Println("Name = ", this.Name, "Sex = ", this.Sex, "Level = ", this.Level)
}

func main() {
	// 创建父类对象
	h := Human{"Zhangsan", "female"}

	h.Eat()
	h.Walk()

	// 创建子类对象
	// 方式一
	// s := Superman{Human{"Lisi", "female"}, 100}

	// 方式二
	var s Superman
	s.Name = "Lisi"
	s.Sex = "female"
	s.Level = 100

	s.Eat()  // 子类的方法
	s.Walk() // 父类的方法
	s.Fly()  // 子类的方法

	s.PrintSuperman()
}