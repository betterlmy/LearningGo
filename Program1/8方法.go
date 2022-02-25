// 方法：指包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或指针 所有给定类型的方法属于该类型的方法集
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64 //定义圆的半径
}

func (c Circle) getArea() float64 {
	return c.radius * c.radius * math.Pi
}
func main() {
	var c1 Circle //声明一个对象
	c1.radius = 10.00
	fmt.Printf("圆的面积 = %f", c1.getArea())
}
