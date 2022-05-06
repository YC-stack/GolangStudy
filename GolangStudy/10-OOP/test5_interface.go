package main

import (
	"fmt"
)
// interface{} 是空接口，是万能数据类型，所有数据类型底层都会实现interface{}
func MyFunc(arg interface{}) {
	fmt.Println("MyFunc is called...")
	fmt.Println("arg = ", arg)

	// Golang为interface{}提供“类型断言“机制，判断interface{}此时引用的具体数据类型
	value, ok := arg.(string) // interface{}类型断言写法，判断arg是不是string，返回arg的值value，和判断结果ok

	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
	}
}

type Book struct {
	Title string
}

func main() {
	// 任意数据类型都可以传入MyFunc
	book := Book{"Golang"}
	MyFunc(book)
	fmt.Println("-------------------")

	MyFunc(100)
	fmt.Println("-------------------")

	MyFunc(3.14)
	fmt.Println("-------------------")

	MyFunc("abc")
}