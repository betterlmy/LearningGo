package main

import (
	"fmt"
	"strconv"
)

// Go语言中 range关键字用于进行迭代array，slice channel或者map返回值为key-value对
func main() {
	nums := []int{1, 2, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Printf(strconv.Itoa(sum) + "\n")

	for _, str := range "who are u?" {
		fmt.Printf(string(str))
	}
}
