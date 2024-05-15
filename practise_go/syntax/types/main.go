package main

import "fmt"

func main() {
	u1 := &User{}
	println(u1)
	u1 = new(User)
	println(u1)

	u2 := User{}
	fmt.Printf("%+v \n", u2)
	u2.Name = "Jerry"
	println(u2.Name)

	var u3 User
	fmt.Printf("%+v \n", u3)
	var u4 *User
	// nil
	println(u4)

	u5 := User{Name: "Jerry"}
	fmt.Printf("%+v \n", u5)
	// 按顺序写，需要写全并且顺序不能出错
	u5 = User{18, "Jerry", "Nick_Jerry"}
	fmt.Printf("%+v \n", u5)

}
