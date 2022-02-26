package main

import "fmt"

//	channel是一个传递数据的数据结构
//	两个goroutine之间可以通过传递一个指定类型的值进行同步运行和通讯
//	<-用于指定通道的方向 ch <- v 将数据v发送给通道ch  v:= <- ch 将通道ch的内容传递给变量v
//	默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	//test1()
	test2()

}

// test1 不带缓冲器的测试
func test1() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}

// test2 带缓冲区的测试
func test2() {
	ch := make(chan int, 3)
	// 因为ch是带缓冲的通道,我们可以不用取出就可以存入到同道中人
	ch <- 5

	ch <- 10
	ch <- 20
	fmt.Println(len(ch))
}
