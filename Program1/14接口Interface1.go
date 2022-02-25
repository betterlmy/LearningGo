package main

import "fmt"

//Go的接口定义:
//把所有具有共性的方法定义到了一起 任何其他类型只要实现了这些方法 就是实现了这个接口

// Phone 定义一个接口,这个接口中有个方法是call
type Phone interface {
	call()
}
type NokiaPhone struct {
}

func (NokiaPhone) call() {
	fmt.Println("Im Nokia Phone")
}

type IPhone struct {
}

func (IPhone) call() {
	fmt.Println("Im IPhone")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
