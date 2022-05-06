package main

import (
	"fmt"
)

// 值传递
/* func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
 */

// 指针（地址）传递
func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 10
	var b int = 20

	// swap(a, b) // 值传递
	swap(&a, &b)  // 地址传递

	fmt.Println("a =", a, " b =",b)

	// 二级指针
	var p *int
	p = &a       // p指向a的地址

	var pp **int 
	pp = &p      // pp指向p的地址

	fmt.Println("p =", p, "&a =", &a)
	fmt.Println("&p =", &p, "pp =", pp)
}