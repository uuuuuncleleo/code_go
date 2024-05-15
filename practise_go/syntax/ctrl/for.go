package main

func Loop1() {
	// 键入fori，编译器可自动补全for语句
	for i := 0; i < 10; i++ {
		print(i)
	}

	// 等价于上面一种写法
	for i := 0; i < 10; {
		print(i)
		i++
	}
}

func Loop2() {
	// 类似于while循环
	i := 0
	for i < 10 {
		print(i)
		i++
	}
}

func Loop3() {
	// 无限循环
	// 谨慎使用，若无法退出，可能导致CPU占用至100%
	for {
		print("hello")
	}
}

func ForArray() {
	println("遍历数组")
	arr := [3]string{"A", "B", "C"}

	// 一般使用此形式进行遍历
	for index, value := range arr {
		println(index, value)
	}

	// 遍历array，value
	for index := range arr {
		println(index, arr[index])
	}
}

func ForSlice() {
	println("遍历切片")
	arr := []string{"A", "B", "C"}
	for index, value := range arr {
		println(index, value)
	}

	// 遍历array，忽略index
	for _, value := range arr {
		println(value)
	}
}

func ForMap() {
	println("遍历Map")
	m := map[string]string{
		"A": "alpha",
		"B": "beta",
	}
	for key, value := range m {
		println(key, value)
	}
	// 遍历map，忽略key
	for _, value := range m {
		println(value)
	}
	// 遍历map，忽略value
	for key := range m {
		println(key, m[key])
	}
}

type User struct {
	Name string
}

func LoopBug() {
	users := []User{
		{Name: "Tom"},
		{Name: "Jerry"},
	}
	m := make(map[string]*User, 2)
	for _, u := range users {
		// 指向了u的地址
		m[u.Name] = &u
	}
	for key, value := range m {
		// 打印value只会打印地址中的值，即最后一次赋给地址的值
		print(key, value)
	}
}

func ForBreak() {
	i := 0
	for {
		if i > 10 {
			break
		} else {
			print("for里面；", i)
			i++
		}
	}
	println(i)
}

func ForContinue() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		} else {
			println(i)
		}
	}
}
