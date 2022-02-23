package main

import (
	"fmt"
)

func formattedOutput() string {
	var stockcode = 123
	var enddate = "2022-02-22"
	var url = "Code=%d,endDate=%s,%s"
	var targetUrl = fmt.Sprintf(url, stockcode, enddate, "来自同一个package下的另一个go文件的输出")
	return targetUrl
}

func variableTest() (string, string) {
	// 声明变量的方式
	var str1, str2 = "i am a ", "boy" //1.用=直接赋值可以省略变量类型的声明
	var a string                      //2.带变量类型default is None
	str3 := "girl"                    //3.使用:=声明，它只能被用在函数体内，而不可以用于全局变量的声明与赋值。

	//注意：声明的变量必须是未被声明过的
	/*
		上述代码相当于
		var str3 = "girl"
	*/

	e, f := 123, "hello" // 并行赋值 e是int型 f是string型
	fmt.Println(e, f)

	return str1 + str2 + a, str3
}
