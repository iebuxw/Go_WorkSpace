package main

import (
	"fmt"
)

//通道（channel）是用来传递数据的一个数据结构。
//定义
//var 变量名 chan 数据类型
//管道的本质是一个队列

//var intB chan int
var intB = make(chan int, 50)

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
	fmt.Println(num2) //输出10

	// 创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	// time.Sleep(time.Second * 10)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

// write Data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Println("writeData ", i)
		// time.Sleep(time.Second)
	}
	close(intChan) // 关闭，关闭之后只能读不能写
}

// read data
//2.1 什么时候发生阻塞？
//	向一个值为nil的管道写或读数据
//	无缓冲区时单独的写或读数据
//	缓冲区为空时进行读数据
//	缓冲区满时进行写数据
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan //从管道获取值。1、会检查是否有在写入，如果有就不会进break 2、会阻塞
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	// readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)
}
