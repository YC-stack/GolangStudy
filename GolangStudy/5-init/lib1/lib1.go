package lib1

import (
	"fmt"
)

// 对外接口，函数名首字母大写代表对外开放，小写代表仅在包内可以使用
func Lib1Test() {
	fmt.Println("Lib1Test...")
}

// 包的初始化函数
func init() {
	fmt.Println("lib1 init...")
}