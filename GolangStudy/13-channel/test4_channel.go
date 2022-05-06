package main

import (
	"fmt"
)

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	// select可以同时监控多个channel的状态
	for {
		select {
		case c <- x: // 如果c可写，则会进入该case
			t := x
			x = y
			y = t+y
		case <-quit: // 如果quit可读，则会进入该case
			// 当quit可读时，说明sub go中向quit中写入了数据
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {
		for i := 0; i < 10; i++ {
			data := <-c // 尝试从c中读数据
			fmt.Println(data)
		}
		// 读完之后，向quit中写入一个数据
		quit <- 0
	}()
	
	// main go
	fibonacii(c, quit)

}