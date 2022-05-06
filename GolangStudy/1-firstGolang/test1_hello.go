package main // 程序的包名

/* import "fmt"
import "time" */

import (
	"fmt"
	"time"
)

func main()  {
	// golang表达式加不加 ";" 都可以，建议不加
	fmt.Println("hello Go!")

	time.Sleep(1*time.Second)

}
