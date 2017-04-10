package main

import "fmt"

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func test1() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(typeof(a), a)
	b := a[1:4:6]
	fmt.Println(typeof(b), b, len(b), cap(b))
	c := make([]int, 5, 10)
	fmt.Println(typeof(c), c, len(c), cap(c))
}

func test2() {
	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 10; i++ {
		fmt.Println("execute", i)
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
	fmt.Println(s)

	a1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := a1[0:2:8]    // {0,1}
	s2 := a1[4:]       // {4,5,6,7}
	s3 := copy(s1, s2) // s2 copy to s1, s1 = {4,5,6}
	fmt.Println(s1, s2, s3)

	a1 = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = a1[0:3:8] // {0,1,2}
	s2 = a1[4:]    // {4,5,6,7}
	copy(s2, s1)   // s0 copy to s2: s2 = {0,1,2,7}
	fmt.Println(s1, s2)

}

func test3() {
	m1 := map[string]int{
		"a": 1, "b": 2,
	}
	m2 := make(map[string]int, 10)
	m2["a"] = 1
	fmt.Println(m1, m2)

	v, ok := m2["a"]
	fmt.Println(v, ok) // v=0, ok=true
	v, ok = m2["b"]
	fmt.Println(v, ok) // v=0, ok=false

	m2["a"] = 3 // update key's value
	m2["b"] = 2 // add k-v
	fmt.Println(m2)

	delete(m2, "a") //delete ok
	fmt.Println(m2)
	delete(m2, "a") //delete failed, no error
	fmt.Println(m2, len(m2))
	for k := range m1 {
		fmt.Println(k, m1[k])
	}
	type user struct {
		name string
		age  int
	}
	m3 := map[int]user{
		1: {name: "user1", age: 10},
	}
	fmt.Println(m3)
	// m3[1].age = 12 // cannot assign to struct field m3[1].age in map
	u := m3[1]
	u.age = 12
	m3[1] = u // replace

	m4 := map[int]*user{ //use pointer
		1: &user{name: "user1", age: 10},
	}
	m4[1].age = 11
	fmt.Println(m4[1])
}

func test4() {
	type Node struct {
		id   int
		data *byte
		next *Node //指向自身的指针
	}
	n1 := Node{id: 1, data: nil}
	n2 := Node{id: 2, data: nil, next: &n1}
	n3 := Node{2, nil}  //  too few values in struct initializer
	fmt.Println(n1, n2) // {1 <nil> <nil>} {2 <nil> 0xc420084240}
	type Resource struct {
		id int
	}
	type User struct {
		name string
		Resource
	}
	type Manager struct {
		User  // 匿名字段，成员名是User，成员类型是User
		title string
	}
	m1 := Manager{User{name: "tom"}, "admin"}
	fmt.Println(m1)
	var m Manager
	m.id = 1 //访问匿名字段，编译器由内向外逐级查找所有层次的匿名字段
	m.name = "tom"
	m.title = "admin"
	fmt.Println(m)

	r := Resource{1}
	fmt.Println(r.id)
}

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println(a == b)
	test1()
	test2()
	test3()
	test4()
}
