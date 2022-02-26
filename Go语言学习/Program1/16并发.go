package main

import (
	"fmt"
	"time"
)

//通过go关键字开启goroutine(routine 例行程序,运行期线程,轻量级的)
//在一个程序中,所有的goroutine共享地址空间
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	// 普通的并发
	go say("world")
	say("hello")

	//并发通讯
}
