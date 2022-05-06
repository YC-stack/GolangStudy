package main

import (
	"fmt"
)

func printArray(myArray [4]int) {
	
	for index,value := range myArray {
		fmt.Println("index = ", index, " value = ", value)
	}

	myArray[0] = 111 // 注意，固定长度数组是值传递，形参是实参的一个副本，修改形参不会影响实参
}

func printArray1(myArray []int) {
	// _ 代表匿名变量
	for _,value := range myArray {
		fmt.Println(" value = ", value)
	}

	myArray[0] = 111 // 注意，动态数组是地址传递，修改形参同时修改实参
}


func main() {
	/* ----------------固定长度数组---------------- */
	var myArray1 [10]int // 默认值初始化为0
	myArray2 := [10]int {1,2,3,4} // 给前四个元素赋值
	myArray3 := [4]int {11,22,33,44}

	// 遍历固定长度数组，方式一
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	// 方式二，range返回数组的索引和值
	for index,value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	// 查看数组数据类型
	fmt.Printf("myArray1 type = %T\n", myArray1)
	fmt.Printf("myArray2 type = %T\n", myArray2)
	fmt.Printf("myArray3 type = %T\n", myArray3)

	printArray(myArray3)
	// 修改后再打印一次
	fmt.Println("---------------")
	for index,value := range myArray3 {
		fmt.Println("index = ", index, " value = ", value)
	}


	/* ----------------动态数组，切片，slice---------------- */
	myArray := []int{1,2,3,4} 
	// 查看数组数据类型
	fmt.Printf("myArray type = %T\n", myArray)

	printArray1(myArray)
	// 修改后再打印一次
	fmt.Println("---------------")
	for index,value := range myArray {
		fmt.Println("index = ", index, " value = ", value)
	}

	/* ----------------声明slice的四种方式---------------- */
	// 1.声明slice1是一个切片，并且初始化，长度为3
	slice1 := []int{1,2,3}

	// 2.声明slice2是一个切片，当没有分配空间，长度为0，再用make给slice2分配3个int空间，默认值是0
	//var slice2 []int
	//slice2 = make([]int,3)

	// 3.声明slice3是一个切片，同时用make给slice3分配3个int空间
	//var slice3 []int = make([]int, 3)

	// 4.通过 := 自动推导slice4是一个有3个int空间的切片
	//silce4 := make([]int, 3)

	// 判断一个切片是否为空
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 不是一个空切片")
	}

	/* ----------------slice的追加---------------- */
	// 声明一个长度为3，容量为5的切片，numsber = [0,0,0], len = 3, cap = 5
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, numbers = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers追加一个元素1，numsber = [0,0,0,1], len = 4, cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, numbers = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers追加一个元素2，numsber = [0,0,0,1,2], len = 5, cap = 5
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, numbers = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers追加一个元素3，超出容量后，会按照之前cap的两倍扩容，numsber = [0,0,0,1,2,3], len = 6, cap = 10
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, numbers = %v\n", len(numbers), cap(numbers), numbers)

	// 如果不指定容量，默认cap = len
	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, numbers2 = %v\n", len(numbers2), cap(numbers2), numbers2)


	/* ----------------slice的截取---------------- */
	s := []int{1,2,3}

	// 截取索引 [0,2) 左闭又开区间的元素
	s1 := s[0:2] // s1 = [1,2]

	fmt.Println(s1)

	// 注意：截取是浅拷贝， s和s1还是指向同一块地址，修改其中一个另一个也会同时修改
	s1[0]=100 

	fmt.Println(s)
	fmt.Println(s1)

	// 给s2开辟新的空间
	s2 := make([]int, 3)
	// copy是深拷贝，把s拷贝给s2，s和s2拥有独立的地址空间
	copy(s2,s)
	fmt.Println(s2)

}

