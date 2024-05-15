package main

// Defer 延迟调用defer
func Defer() {
	// defer 类似于栈，后进先出（先定义的后执行，后定义的先执行）
	defer func() {
		println("第一个defer")
	}()

	defer func() {
		println("第二个defer")
	}()
}

// DeferClosure 作为闭包引入：执行defer对应方法的时候才确定
func DeferClosure() {
	j := 0
	defer func() {
		// 闭包写法
		println(j)
	}()
	j = 1
}

// DeferClosureV1 作为参数传入：定义defer的时候就确定了
func DeferClosureV1() {
	j := 0
	// 参数传递
	defer func(val int) {
		println(val)
	}(j)
	j = 1
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

type MyStruct struct {
	name string
}

func DeferReturnV2() *MyStruct {
	res := &MyStruct{
		name: "alpha",
	}
	defer func() {
		res.name = "beta"
	}()

	return res
}

func DeferClosureLoopV1() {
	for m := 0; m < 5; m++ {
		defer func() {
			print(m)
		}()
	}
}

func DeferClosureLoopV2() {
	for m := 0; m < 5; m++ {
		defer func(val int) {
			print(val)
		}(m)
	}
}

func DeferClosureLoopV3() {
	for m := 0; m < 5; m++ {
		n := m
		defer func() {
			print(n)
		}()
	}
}
