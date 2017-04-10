package main

import (
	"fmt"
)

//1 ---------------------
type Circle struct {
	r float64
}

func (c Circle) getArea(msg string) float64 {
	fmt.Println("msg is: ", msg)
	return c.r * c.r * 3.14
}

//2 ---------------------
type User struct {
	id   int
	name string
}

type Admin struct {
	id   int
	name string
}

type Manager struct {
	User
	Admin
}

func (self User) ToString() string {
	return fmt.Sprintf("User %p, %v", &self, self)
}

// func (self Admin) ToString() string { //error
// return fmt.Sprintf("User %p, %v", &self, self)
// }

//3 --------------------
func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, *self)
}

//4 --------------------
func (self User) ChangeName(newName string) {
	fmt.Println(self.name) // tom
	self.name = newName
	fmt.Println(self.name) //amy
}

//5 --------------------
//6 --------------------

func main() {
	c := Circle{2}
	fmt.Println(c.getArea("get area"))
	m := Manager{User{1, "tom"}, Admin{2, "jack"}}
	fmt.Println(m.ToString())

	u := User{1, "tom"}
	u.Test()
	(*User).Test(&u) // 类型调用需要显示传递receiver

	fmt.Println(u.name) //tom
	u.ChangeName("amy")
	fmt.Println(u.name) //tom
}
