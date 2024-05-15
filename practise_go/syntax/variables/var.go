package main

import "practise_go/syntax/variables/demo"

var (
	alpha = 123
	beta  = 12.3
)

// 常量定义使用 const, 定义后不可更改变量值
const (
	cannotChange = 123
)

func main() {
	var a int = 123
	println(a)

	var b float64 = 12.3
	println(b)

	var c string = "string变量"
	println(c)

	var d int
	println(d)

	// 此变量声明方式只能作用于局部变量（即方法内部），会自动进行类型推断
	e := 123
	println(e)

	println(demo.Global)
}
