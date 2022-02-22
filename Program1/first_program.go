package main //指向自己的包 所有的包名都应该使用小写字母。
// `package main`表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 `main` 的包

import (
	fm "fmt" //包的重命名
	mathClass "learngo/Program2"
)

func main() {
	fm.Println("hello world")
	println(sumAB(5, 6))
	println(GetList(7, 8))
	println(mathClass.Add(1, 2)) // 调用其他package下的方法
	fm.Println(concat("我叫", "李梦洋"))
	fm.Println(formattedOutput())
}

//sumAB 计算a+b的和
func sumAB(a int, b int) int {
	return a + b
}

// GetList 将ab返回
func GetList(a int, b int) (int, int) {
	return a, b
}

// concat 连接两个字符串
func concat(str1 string, str2 string) string {
	return str1 + str2
}
