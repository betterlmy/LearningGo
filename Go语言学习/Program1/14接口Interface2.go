package main

type Human interface {
	name() string
	age() int
}
type Woman struct {
}

func (w Woman) name() string {
	return "XiaoFang"
}

func (w Woman) age() int {
	return 15
}

type Man struct {
}

func (m Man) name() string {
	return "LiLei"
}

func (m Man) age() int {
	return 12
}

func main() {
	var human Human
	human = new(Man)
	println(human.name())
}
