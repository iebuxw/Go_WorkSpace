package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func say1(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 并发执行程序
	go say("world")
	//go say1("hello")
	//say("world")
	fmt.Println(1111)
	//当main()函数结束时，程序也将结束。它不等待其他goroutine完成。所以需要say1，不然没有输出
	say1("hello3")
}
