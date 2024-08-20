package main

import (
	"fmt"
)

//通道（channel）是用来传递数据的一个数据结构。管道的本质是一个队列
//作用：协程之间通信
//通道缓冲区：相当于队列元素个数
//var 变量名 chan 数据类型
//通道的关闭必须由发送者
//通道通常不需要关闭，除非用range，range需要关闭，否则会死锁
/*

应用场景：
	1、超时控制
	2、定时任务
	3、生产消费
	4、控制并发数
*/

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

