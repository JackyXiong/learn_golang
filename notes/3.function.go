package main

import (
	"fmt"
	"reflect"
)

func test1(x int, y int) (sum, diff int) { // { 不能换行
	sum = x + y
	diff = x - y
	return x, y
	// return // 使用sum 和 diff 命名隐式返回
}

func test2(fn func() int) int { // fn 作为参数，类型是函数，其返回值是int
	return fn() // 函数作为值返回
}

func test3(s string, nums ...int) {
	var sum int
	for index, n := range nums { // index 从 0 开始
		fmt.Println(index, n)
		sum += n
	}
	fmt.Println(s, nums)                 // print test3 [1,2,3]
	fmt.Println(fmt.Sprint("sum=", sum)) // print sum=6

}

func test4(x int) (z int) {
	defer func() {
		fmt.Println("execute last")
		z += 1
		fmt.Println("z=", z)
	}()
	z = 100 + x
	fmt.Println("execute first")
	fmt.Println("z=", z)
	return
}

func test5(x int) int {
	x = x + 1
	return x
}

func test6(x *int) int {
	*x = *x + 1
	return *x
}

func a(x int, y string) {}

func main() {
	fmt.Println(test1(10, 5))
	// var s = make([]int, 2)
	// s = test1(10, 5)                             //  multiple-value test1() in single-value context
	fmt.Println(test2(func() int { return 10 })) // 匿名函数作为参数
	test3("test3", 1, 2, 3)
	fmt.Println(test4(10))

	x := 1
	x1 := test5(x)
	fmt.Println(x, x1) // 1 2

	x = 1
	x1 = test6(&x)
	fmt.Println(x, x1) // 2 2
	type A func(int, string)
	fmt.Println(reflect.TypeOf(a))

}
