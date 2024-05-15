package main

import "unicode/utf8"

func main() {
	//Hello()
	//print("go")
	String()
	Byte()
}

func String() {
	var strDemo string = "hello你好"
	println(len(strDemo))
	println(utf8.RuneCountInString(strDemo))
}

func Byte() {
	var a byte = 'a'
	println(a)

	var str string = "hello"
	var bs []byte = []byte(str) // byte切片
	var str1 = string(bs)
	println(bs)
	println(str1)

}
