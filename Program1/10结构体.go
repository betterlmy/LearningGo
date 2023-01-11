package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	booka := Books{title: "Go语言教程", author: "李梦洋", subject: "编程", bookId: 1}
	fmt.Printf("书籍名字是《%s》", booka.title)
}
func (book Books) getBookName() string {
	return book.title
}
