package main

import "fmt"

//通道（channel）是用来传递数据的一个数据结构。
//定义
//var 变量名 chan 数据类型

func main() {
	//创建管道
	//var intC chan int
	intC := make(chan int, 3)
	//向管道写入数据
	intC <- 10
	num := 211
	intC <- num

	//从管道读取数据
	num2 := <-intC
	fmt.Println(num2)
}
