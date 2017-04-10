package main

import "fmt"

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// 1-------------------
type Stringer interface {
	String() string
}

type Printer interface {
	Stringer // 嵌入
	Print()
}

type User struct {
	id   int
	name string
}

func (self User) String() string {
	return fmt.Sprintf("user %d. %s", self.id, self.name)
}

func (self User) Print() {
	fmt.Println(self.String())
}

func test1() {
	var t Printer = User{1, "Tom"} //User 实现了String 和 Print
	t.Print()
	t = &User{1, "Tom"} //*User 实现了String 和 Print
	t.Print()
}

func test2() {
	u := User{1, "tom"}
	var vi, pi interface{} = u, &u
	fmt.Println(vi, u)      // user 1. tom
	fmt.Println(typeof(vi)) // main.User
	fmt.Println(typeof(pi)) // *main.User
	fmt.Println(vi.(User))  // user 1. tom
	if v, ok := vi.(User); ok {
		fmt.Println(v, ok)
	}

	switch value := vi.(type) {
	case User:
		fmt.Println("---", value)
	}
}

func test3() {
	var p Printer = User{1, "amy"}
	var o Stringer = p
	fmt.Println(o.String())
}

func main() {
	test1()
	test2()
	test3()
}
