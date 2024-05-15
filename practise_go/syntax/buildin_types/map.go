package main

func Map() {
	m1 := map[string]string{
		"key1": "value1",
	}
	println(m1)

	m2 := make(map[string]string, 4)
	m2["key2"] = "value2"
	println(m2)

	val1, ok := m1["key1"]
	println(val1, ok)
	val2, ok := m1["key2"]
	println(val2, ok)

	val2 = m2["key2"]
	println(val2)

	// 遍历map时，顺序是随机的
	for k, v := range m1 {
		println(k, v)
	}

	for k := range m1 {
		println(k, m1[k])
	}

	for _, v := range m1 {
		println(v)
	}

	delete(m1, "key1")
}
