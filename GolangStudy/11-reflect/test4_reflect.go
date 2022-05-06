package main

import (
	"fmt"
	"reflect"
)

// 反射一个变量的类型和值
func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("value : ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 3.14

	reflectNum(num)

}