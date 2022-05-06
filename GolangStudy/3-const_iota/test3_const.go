package main

import "fmt"

// const 来定义枚举类型
const (
	// 可以在const()添加一个关键字 iota，每行的iota都会累加1，第一行的iota默认值为0
	// 注意：iota只能在const()中使用，实现累加效果
	BEIJING = iota // iota = 0
	SHANGHAI       // iota = 1
	SHENZHEN       // iota = 2
)

const (
	a, b = iota+1, iota+2 // iota = 0
	c, d                  // iota = 1
	e, f                  // iota = 2

	g, h = iota*2, iota*3 // iota = 3
	i, k                  // iota = 4
)

func main() {
	// 常量（只读属性）
	const length int = 10
	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

}

