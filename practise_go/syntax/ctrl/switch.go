package main

func Switch(status int) string {
	switch status {
	case 0:
		return "start"
	case 1:
		return "running"
	case 2:
		return "stop"
	default:
		// 可以不写default
		return "unknown"
	}
}

func SwitchNoValue(age int) string {
	// switch不加value的用法，case后面跟bool类型的表达式，需尽量做到每个条件互斥
	switch {
	case age >= 18:
		return "adult"
	case age >= 12:
		return "teenage"
	}
	return "child"
}
