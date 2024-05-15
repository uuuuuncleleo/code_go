package main

import "fmt"

func Array() {
	// 定义数组(名称:=[cap]int{val1, val2...}，value值<=cap值
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1: %v, len: %d, cap: %d\n", a1, len(a1), cap(a1))
	// value数量<cap值时，int型默认填充0
	a2 := [3]int{1, 2}
	fmt.Printf("a2: %v, len: %d, cap: %d\n", a2, len(a2), cap(a2))
	// 未定义value时，已分配好地址，默认填充0
	var a3 [3]int
	fmt.Printf("a3: %v, len: %d, cap: %d\n", a3, len(a3), cap(a3))

	fmt.Printf("a1第一个值：%d\n", a1[0])
}
