package main

import "fmt"

func pointerAndAddress() {
	//
	a := 100
	fmt.Println(&a)
	ptr := &a //ptr是指针类型的变量
	fmt.Println(ptr)
}
func main() {
	pointerAndAddress()
}
