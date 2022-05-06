package main

import (
	"fmt"
)


func printMap(cityMap map[string]string) {
	// 遍历
	for key,value := range cityMap {
		fmt.Println("key = ", key, " value = ", value)
	}
}

func changeMap(cityMap map[string]string) {
	// 注意，map是引用传递
	cityMap["UK"] = "London"
}

func main() {
	/*-------------------------------map的三种声明方式-------------------------------*/
	// 方式一，声明myMap1是一种map类型，键是string类型，值是string类型，此时是一个空map，还没有分配空间
	var myMap1 map[string]string
	// 使用make给myMap1分配10个键值对的空间
	myMap1 = make(map[string]string, 10)

	// 插入键值对，注意，map是乱序的
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// 方式二，声明的同时分配空间
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	// 方式三，声明的同时初始化
	myMap3 := map[string]string {
		"one": "java",
		"two": "c++",
		"three": "python",  // 注意，最后这里有","
	}

	fmt.Println(myMap3)
	fmt.Println("--------------------")

	/*-------------------------------map的使用-------------------------------*/
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	printMap(cityMap)
	fmt.Println("--------------------")

	// 删除
	delete(cityMap, "China")
	printMap(cityMap)
	fmt.Println("--------------------")

	// 修改
	cityMap["USA"] = "DC"
	changeMap(cityMap)
	printMap(cityMap)


}