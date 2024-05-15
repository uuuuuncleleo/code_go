package main

import "fmt"

// Slice 切片
func Slice() {
	// 推荐写法：s := make([]types, 0, capacity)
	// 提前预估好容量capacity，避免进行扩容

	// 直接定义好切片，并分配地址，value值及容量
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))

	// 创建一个长度为3，容量为4的切片
	s2 := make([]int, 3, 4)
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	// 追加一个元素
	s2 = append(s2, 7)
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	// 扩容
	s2 = append(s2, 8)
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	// make只传入一个参数n，默认创建一个长度为n，容量为n的切片
	s3 := make([]int, 4)
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))
}

func subSlice() {
	//
	s1 := []int{2, 4, 6, 8, 10}
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	s2 := s1[0:2]
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	s3 := s1[1:]
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))
	s4 := s1[:3]
	fmt.Printf("s4: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))
}

func shareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s2[0] = 5
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s2 = append(s2, 6)
	s2[0] = 7
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

}
