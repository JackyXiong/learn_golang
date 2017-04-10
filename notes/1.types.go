package main

import "fmt"

var a int // 不初始化，默认为当前类型的零值
var b int8 = 100
var c = 2
var d float64
var e, f, g int
var (
	h int
	i float32
)

const x int = 1 //定义常量

func defineVar() {
	e := 1  // 函数内部简洁形式
	a := 10 // 函数内部创建新的a变量，不影响外部
	fmt.Println(&e, a)
	e, f := 2, 1 // 重新对e赋值
	fmt.Println(&e, f)

}

func main() {
	fmt.Println(a, b, c, d)
	defineVar()
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g)
	fmt.Println(h, i)
	// x = 10 error: 修改常量
	defineType()
	s := "熊杰"
	q := []byte(s)
	fmt.Println(q)
}

/*
类型转换不支持隐式类型转换
字符串：不可变类型，内部是指针指向的字节数组
  * 默认空串，无法修改字节数组
  * s[i]访问某个字节
  * s1 = s[1:2] s1依然指向s。只是指针位置不同，修改了长度属性
  * 修改方式：先转换为[]byte，修改数组元素后，在转换回来

*/

func defineType() {
	type bigint int64 //定义新类型，不是原类型的别名，只是拥有相同的数据结构
	var a bigint = 12
	fmt.Println(a)
}
