package main

// 导包时会先递归调用包的init()函数
import (
	"GolangStudy/5-init/lib1" // 需要添加包在SRC目录下的路径
	"GolangStudy/5-init/lib2"

	// 三种特殊导包方法
	// _ "GolangStudy/5-init/lib1" // 匿名导包，这样导包无法使用包中的方法，但会调用该包的init函数，而不会报错
	// mylib2 "GolangStudy/5-init/lib2" // 别名导包，mylib2是lib2的别名，在包名长的时候可以取个简单的别名
	// . "GolangStudy/5-init/lib1"      // 将该包中的全部方法导入到当前包，可以直接使用该包中API，而不用包名.API（为避免函数同名歧义，不推荐使用该方法）
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}