package main

func IfOnly(age int) string {
	if age >= 18 {
		return "adult"
	}
	return "child"
}

func IfElse(age int) string {
	if age >= 18 {
		return "adult"
	} else {
		return "child"
	}
}

func IfElseIf(age int) string {
	if age >= 18 {
		return "adult"
	} else if age >= 12 {
		return "teenager"
	} else {
		return "child"
	}
}

func IfNewVariable(start, end int) string {
	//distance := end - start
	//if distance > 100 {
	//	return "too far"
	//} else {
	//	return "that's ok"
	//}

	// 另一种写法，但distance只作用与if-else语句中
	if distance := end - start; distance > 100 {
		println(distance)
		return "too far"
	} else {
		return "that's ok"
	}
}
