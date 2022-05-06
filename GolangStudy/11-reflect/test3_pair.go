package main

import (
	"fmt"
)

// 定义接口
type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 定义具体类型
type Book struct {

}

// 重写接口
func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}

func main() {
	// b: pair<type: Book, value: Book{}的地址>
	b := &Book{}

	// r: pair<type: , value: >
	var r Reader

	// r: pair<type: Book, value: Book{}的地址>
	r = b
	r.ReadBook() // 调用的是具体数据类型Book的方法

	// w: pair<type: , value: >
	var w Writer
	// w: pair<type: Book, value: Book{}的地址>
	w = r.(Writer)  // 将Reader断言转换为Writer，为什么能成功？因为Reader和Writer有相同的type
	w.WriteBook()

}