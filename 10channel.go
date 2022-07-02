package main

import (
	"fmt"
)

//通道（channel）是用来传递数据的一个数据结构。管道的本质是一个队列
//作用：协程之间通信
//var 变量名 chan 数据类型

//var intB chan int
var intB = make(chan int, 50)

func main() {
	//创建管道
	//var intC chan int
	intC := make(chan int, 3)
	//向管道写入数据
	intC <- 10

	//从管道读取数据
	num2 := <-intC
	fmt.Println(num2) //输出10

	//关闭管道，关闭后只能读不能写
	close(intC)
}

func readData(intChan chan int) {
	//固定写法，读取管道并检测是否关闭
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
}
