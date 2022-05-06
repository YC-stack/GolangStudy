package main

import (
	"fmt"
)
/* ------------------------Golang用接口interface实现多态------------------------ */
// 定义接口，接口本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string // 注意写返回值
	GetType() string 
}

// 定义具体的类 猫
type Cat struct {
	Color string
}

// 绑定猫类，重写接口
func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 定义具体的类 狗
type Dog struct {
	Color string
}

// 绑定狗类，重写接口
func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.Color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// 传什么子类，调该子类重写的接口，实现多态
func ShowAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}


func main() {
	var animal AnimalIF  // 接口的数据类型，父类指针

	animal = &Cat{"White"} // 父类指针指向子类对象
	animal.Sleep() // 调用的就是Cat的方法

	animal = &Dog{"Black"} // 父类指针指向子类对象
	animal.Sleep() // 调用的就是Dog的方法

	fmt.Println("----------------")
	// 函数传参的方式实现多态
	cat := Cat{"White"}
	dog := Dog{"Black"}

	
	ShowAnimal(&cat) // 主要要传地址

	ShowAnimal(&dog)
}