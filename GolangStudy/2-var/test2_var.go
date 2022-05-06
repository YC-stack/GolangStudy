package main

/*四种变量声明方式*/

import  "fmt"


// 方法一、二、三可以用于声明全局变量
var gA int
var gB int = 100
var gC = 100

// 方法四不能用于声明全局变量
// gD := 100 


func main()  {
	// 方法一：只声明不赋值，默认值为0
	var a int
	fmt.Println("a = ",a)

	// 方法二：声明一个变量，并初始化一个值
	var b int =100
	var bb string = "abc"
	fmt.Println("b = ", b)
	fmt.Println("bb = ", bb)

	// 方法三：在初始化的时候，可以省略数据类型，根据值自动匹配当前变量的数据类型
	var c = 100
	var cc = "abc"
	fmt.Println("c = ",c)
	fmt.Println("cc = ",cc)

	// 方法四：（常用的方法）省去关键字var，直接自动匹配，注意：方法四只能用在函数体内声明变量，不能声明全局变量
	d := 100
	e := "abc"
	f := 3.14
	fmt.Println("d = ",d)
	fmt.Println("e = ",e)
	fmt.Println("f = ",f)

	// 声明多个变量
	var g,h int =100,200 // 相同数据类型
	var  i,j = 100,"abc"   // 不同数据类型
	fmt.Println("g = ",g)
	fmt.Println("h = ",h)
	fmt.Println("i = ",i)
	fmt.Println("j = ",j)

	var(
		k int = 100
		l string = "abc"
		m bool = true
	)

	fmt.Println("k = ",k)
	fmt.Println("l = ",l)
	fmt.Println("m = ",m)

}