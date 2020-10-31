package main

import "fmt"

/**
Go 语言常量
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
常量的定义格式：
const identifier [type] = value
你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
显式类型定义： const b string = "abc"
隐式类型定义： const b = "abc"
多个相同类型的声明可以简写为：
const c_name1, c_name2 = value1, value2
*/
func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	// 多重赋值
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
}
