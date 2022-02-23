package main

import "fmt"

func main() {
	//switchTest()
	//typeSwtich()
	//fallthroughTest2()
}
func switchTest() {
	a := 2
	switch a {
	case 1:
		{
			fmt.Println("case1")
		}
	default:
		fmt.Println("default")

	}
}

// typeSwtich() 使用switch语句检测变量的类型
func typeSwtich() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是%T", i)
	case int, string: //多条件匹配 满足其中一个即可
		fmt.Println("x的类型是int或str")
	}
}

// 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
func fallthroughTest1() {
	switch { //默认为true
	case true:
		println(1)
	case false:
		println(2)
	default:
		println(3)
	}
}

func fallthroughTest2() {
	switch { //默认为true
	case true:
		println(2)
		fallthrough
	case false:
		println(3)
		fallthrough
	default:
		println(3)
	}
}
