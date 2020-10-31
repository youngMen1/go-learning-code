package main

import "fmt"

/*
Go 语言变量
Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。
声明变量的一般形式是使用 var 关键字：
var identifier type
可以一次声明多个变量：
var identifier1, identifier2 type
*/
func main() {

	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
}
