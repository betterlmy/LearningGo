package main

import "fmt"

func main() {
	//fmt.Println(getSum())
	//infiniteCycle()
	//iterableOperations()
	//nineTimesNineTable()
	//breakLoopWithoutLable()
	breakLoopWithLable()
}

func countOne2Ten() int {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	return sum
}

//在Go中，for循环的init与post参数可选
// 类似于while操作
func getSum() int {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	return sum

}

// 无限循环
func infiniteCycle() {
	sum := 0
	for {
		sum++
		fmt.Println(sum)
	}
}

//针对可迭代元素的循环操作
func iterableOperations() {
	strings := []string{"google", "baidu"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := []int{1, 2, 3, 5}
	for index, x := range numbers {
		fmt.Printf("第%d位的值是%d\n", index+1, x)
	}
}

// nineTimesNineTable 输出九九乘法表
func nineTimesNineTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("   %d*%d=%d", i, j, i*j)
		}
		fmt.Println()
	}

}

// breakLoopWithoutLable 不带标签默认终止当前循环
func breakLoopWithoutLable() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("i = %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2=%d\n", i2)
			break
		}
	}
}

// breakLoopWithLable 终止当前标签的循环
//i = 1
//i2=11
func breakLoopWithLable() {
re1:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i = %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2=%d\n", i2)
			break re1
		}
	}
}
