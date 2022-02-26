package main

import (
	"fmt"
	"math"
)

func main() {
	//println(funcAsParameters(1.44))
	useOfanonymousFunc()
}

// 将函数作为实参
func funcAsParameters(x float64) float64 {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	return getSquareRoot(x)
}

// 函数的闭包（匿名函数）
//匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
func anonymousFunc() func() int {
	// 此处返回一个函数 函数名字不知道 但是返回值是int类型
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func useOfanonymousFunc() {
	nextNumber := anonymousFunc()
	fmt.Println("1....")
	fmt.Println(nextNumber()) //1
	fmt.Println(nextNumber()) //2
	fmt.Println(nextNumber()) //3

	fmt.Println("2...")
	nextNumber1 := anonymousFunc()
	fmt.Println(nextNumber1()) //1
	fmt.Println(nextNumber1()) //2
}
