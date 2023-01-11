package main //指向自己的包 所有的包名都应该使用小写字母。
// `package main`表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 `main` 的包

import (
	fm "fmt" //包的重命名
	//mathClass "learngo/Program2"
)

func main() {

	// 1. 格式化
	//r1 := "asd"
	//var r2 = r1
	//
	//r1 = "aa"
	//fm.Println(r2)
	////fm.Println("hello world")
	////println(sumAB(5, 6))
	////println(GetList(7, 8))
	////println(mathClass.Add(1, 2)) // 调用其他package下的方法
	////fm.Println(concat("我叫", "李梦洋"))
	////fm.Println(formattedOutput())
	//a, _ := variableTest() //_用于抛弃 即该值没有用 但是仍返回了两个值
	//fm.Println(a)
	//
	//2. 常量
	fm.Println("面积是:", constTest())

}

// sumAB 计算a+b的和
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

func constTest() int {
	const (
		width, height = 100, 50
	)
	return width * height
}
