package main

import "fmt"

func main() {

}

func init() {
	// 基础初始化
	// 仅设置三个 后面七个默认为0
	balance := [10]float32{1, 3, 5}
	fmt.Println(balance)

	// 未知数组长度
	balance2 := [...]float32{7, 8, 9}
	fmt.Println(balance2)

	//初始化部分
	// 长度为7 其中b[3]=5.0 b[6]=7.0
	balance3 := [...]float32{3: 5.0, 6: 7.01}
	fmt.Println(balance3)

	//for i := 0; i < len(balance3); i++ {
	//	println(balance3[i])
	//}
}
