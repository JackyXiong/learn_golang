/*
1.部分保留字：defer, fallthrough, go, goto, interface, map, struct, type, var, select, range, func
2.运算符：
3.初始化:
4.控制流
	* 不支持三元表达式
	* range 类似于迭代器，如果被迭代对象是普通类型，迭代出来的值都是从复制的被迭代对象取出；
	如果被迭代对象是引用类型，则其底层数据不会被复制，如slice map channel
	* swtich case 可省略break语句，默认自动终止, fallthrough 继续下一分支
	* goto, continue, break 可配合标签使用，continue只能用于for循环，break 可用于for switch select
*/

package main

import "fmt"

func testSwitch() {
	a := 1
	b := 0
	switch {
	case a > 0:
		fmt.Println(a, b)
		fallthrough //自动继续下一分支，不判断条件
	case a > 1:
		fmt.Println(a)
	}
}

func testGoto() {
	fmt.Println("---- goto ")
	goto GOTO
	// fmt.Println('不会打印') // 不会打印
GOTO:
	fmt.Println("goto 开始")
}

func main() {
	testSwitch()
	testGoto()
}
