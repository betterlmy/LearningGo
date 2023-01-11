package main

import (
	"fmt"
	"unsafe"
)

// constTest 常量的学习
func constTest() int {
	const (
		weight = 6
		height = 5
	)
	var area int
	const a, b, c = 1, "a", true
	area = weight * height
	fmt.Println(a, b, c)

	const (

		//常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。
		//常量表达式中，函数必须是内置函数，否则编译不过：
		a1 = "abc"
		b1 = len(a1)
		c1 = unsafe.Sizeof(a1)
	)
	println(a1, b1, c1)

	return area
}

func iotaTest() {
	//iota常量 特殊常量 可以认为是一个可以被编译器修改的常量
	//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
	//const 中每新增一行常量声明将使 iota 计数一次
	//(iota 可理解为 const 语句块中的行索引)。
	const (
		a = iota
		b
		c
	)
	println(a, b, c)
}
func useOfIota() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //3
		e        //4
		f = 100  //5
		g        //6
		h = iota //7
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	//0 1 2 ha ha 100 100 7 8
}
func main() {
	//iotaTest()
	useOfIota()
}
