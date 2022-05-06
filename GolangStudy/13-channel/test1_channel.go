package main

import (
	"fmt"
)

func main() {
	// 定义一个channel，传输的数据类型为int，没有缓存（只能传输一个数据）
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine end")

		fmt.Println("goroutine start...")

		c <- 666 // 将666写入c
	}()

	num := <-c // 从c中读取数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main goroutine end")
}