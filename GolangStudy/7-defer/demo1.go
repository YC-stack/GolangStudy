package main

import (
	"fmt"
)

/*知识点一：关键字defer后跟一个表达式，会在该函数全部执行完后，执行该表达式，如果有多个defer，定义defer顺序为压栈顺序，执行defer顺序为出栈顺序*/

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

/*知识点二：return会先于defer调用*/
func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	defer func1()
	defer func2()
	defer func3()

	deferAndReturn()
}