package main

/*golang支持函数有多个返回值*/

import "fmt"

// 两个形参a,b，一个int类型的返回值
func func1(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100

	return c
}

// 两个形参a,b，两个返回值都是int类型，匿名的
func func2(a string, b int) (int, int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 200,200
}

// 两个形参a,b，两个返回值r1,r2都是int类型，有名的，r1,r2默认值为0，作用域为函数体内
func func3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	r1 = 300
	r2 = 300

	return r1,r2
}

func main() {
	res1 := func1("func1", 111)
	fmt.Println("res1 =", res1)

	res21,res22 := func2("func2", 222)
	fmt.Println("res21 =", res21, "res22 = ", res22)

	res31,res32 := func3("func3", 333)
	fmt.Println("res2 =", res31, "res22 = ", res32)
}