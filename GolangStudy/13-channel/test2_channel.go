package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个channel，传输数据类型为Int，缓冲容量为3
	c := make(chan int, 3)

	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	// 注意：当channel满了继续发送的话会，发送方会阻塞，当channel空了继续接收的话，接收方会阻塞
	go func() {
		defer fmt.Println("sub goroutine end")

		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("sub goroutine send: ", i, ", len(c)= ", len(c), ", cap(c)= ", cap(c))
		}
	}()

	time.Sleep(2*time.Second)

	for i := 0; i < 10; i++ {
		num := <-c
		fmt.Println("main goroutine receive: ", num)
	}

	fmt.Println("main goroutine end")
}