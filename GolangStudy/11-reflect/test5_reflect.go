package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID int
	Name string
	Age int
}

func (this *User) Call() {
	fmt.Println("User is called")
	fmt.Printf("%v\n", this)
}

func GetFieldAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("intputType is :", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	// 通过type获取User里面的字段，NumField为字段的个数，reflect包的各种接口参考go文档
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type获取User里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	user := User{1, "Zhangsan", 18}

	GetFieldAndMethod(user)
}