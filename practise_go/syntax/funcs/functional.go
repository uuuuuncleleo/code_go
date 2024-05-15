package main

// Func1 方法可定义多个形参及返回值，统一类型变量可统一进行类型定义
func Func1(a, b, c int, d string) (string, error) {
	return "Func1", nil
}

// Func2 方法可定义返回值名称
func Func2(a int, b int) (str string, err error) {
	str = "hello, world"
	return
}

func Func3(a int, b int, c string) (string, error) {
	// nil为go语言中一个错误
	return "hello", nil
}

// Func4 可将方法赋给一个变量
func Func4() {
	myFunc3 := Func3
	str, _ := myFunc3(1, 2, "Func4 uses Func3")
	print(str)
}

// Func5 方法内部进行方法定义，内部方法只能被方法内部调用
func Func5() {
	fn := func(name string) string {
		return "hello, " + name
	}
	print(fn("go"))
}

func Func6() func(name string) string {
	return func(name string) string {
		return "hello, " + name
	}
}

func Func6Invoke() {
	fn := Func6()
	str := fn("go")
	print(str)
}

// Func7 发起匿名调用
func Func7() {
	fn := func(name string) string {
		return "hello, " + name
	}("Func7")
	print(fn)
}
