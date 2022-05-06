package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// sub goroutine
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close关键字，用于关闭channel
		// 注意：channel关闭后无法向channel再发送数据，但接收方可以继续接收从channel接收数据，直到将数据取完，ok返回false
		close(c)
	}()
	
	// main goroutine死循环一直尝试从channel中接收数据，直到channel关闭
	/* for {
		// data为从channel接收的数据，ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	} */

	// 上面的for循环可以简写，使用range来迭代操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main goroutine end")
}