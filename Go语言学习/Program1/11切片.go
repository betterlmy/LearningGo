package main

import (
	"fmt"
)

// Go语言切片是对数组的抽象
// Go数组的长度是不可改变的
//Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，
//与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

//切片的声明
// 切片不需要说明长度
func definationOfSlice() {

	//使用type创建切片
	//slice := make([]T, length, capacity)
	s := []int{1, 2, 3}
	println(s[0])
	printSlice("s", s)
}

// printSlice 格式化输出切片
func printSlice(name string, s []int) {
	fmt.Printf("%s:len = %d,cap = %d,slice=%v \n", name, len(s), cap(s), s)
}

func copyFromArray() {
	array := [...]int{4, 5, 6}
	s1 := array[:]
	printSlice("s1", s1)

	//part Array
	s2 := array[1:3] //从1到3-1
	printSlice("s2", s2)
	s3 := array[:2]
	printSlice("s3", s3)
}

func main() {
	//definationOfSlice()
	//copyFromArray()
	testSlice()
}

func testSlice() {
	slice := make([]int, 2, 3)
	printSlice("s", slice)
	slice = append(slice, 10, 10, 10, 101, 10)
	printSlice("s", slice)
}
