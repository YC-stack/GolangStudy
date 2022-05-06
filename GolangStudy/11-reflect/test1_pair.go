package main

import (
	"fmt"
)

func main() {
	var a string
	// a: pair<type: string, value: "abc">
	a = "abc"

	// 把a赋值给allType后，allType中的pair的type和value就分别指向了a中的pair的type和value
	// allType: pair<type: string, value: "abc">
	var allType interface{}
	allType = a

	value,_ := allType.(string) // 断言，得到allType的类型和值
	fmt.Println(value)

}