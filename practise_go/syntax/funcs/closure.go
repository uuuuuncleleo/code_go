package main

// Closure 闭包
func Closure(name string) func() string {
	// 闭包：引用了上下文（例如name变量）
	// 闭包使用不当可能会造成内存泄漏
	return func() string {
		return "hello, " + name
	}
}

func ClosureInvoke() {
	c := Closure("funcClosure")
	print(c())
}
