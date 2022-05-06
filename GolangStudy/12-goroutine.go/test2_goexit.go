package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 用go创建goroutine，执行一个匿名函数，形参为空，返回值为空
	go func() {
		defer fmt.Println("A.defer")

		// 子函数定义并调用
		func() {
			defer fmt.Println("B.defer")

			// 退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}() // 注意加()表示定义函数后并调用该函数

	// 用go创建goroutine，执行一个匿名函数，形参为int,int，返回值为bool
	go func(a int, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20) // 传参

	// 主goroutine死循环
	for {
		time.Sleep(1*time.Second)
	}
}