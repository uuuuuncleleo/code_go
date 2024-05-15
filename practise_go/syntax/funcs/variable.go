package main

// YourNames Go方法支持不定参数，不定参数为最后一个参数，可以传入任意多个值
func YourNames(name string, aliases ...string) {
	// 必须是最后一个参数才能声明为不定参数
	// 不定参数在方法内部可以被当成切片使用
	if len(aliases) > 0 {
		println(aliases[0])
	}
}

func YourNamesInvoke() {
	YourNames("alpha")
	YourNames("alpha", "beta")
	YourNames("alpha", "beta", "gamma")
}
