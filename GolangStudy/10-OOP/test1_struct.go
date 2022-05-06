package main

import (
	"fmt"
)

// 声明一种新的数据类型 myInt，是int的别名
type myInt int

// 声明一个结构体
type Book struct {
	title string
	auth string
}

// 注意，直接传结构体是值传递
func changeBook(book Book) {
	book.auth = "Lisi"
}

// 指针传递
func changeBook1(book *Book) {
	book.auth = "Wangwu"
}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "Zhangsan"
	fmt.Printf("%v\n", book1)

	// 修改，值传递
	changeBook(book1)
	fmt.Printf("%v\n", book1)

	// 修改，指针传递
	changeBook1(&book1)
	fmt.Printf("%v\n", book1)

}