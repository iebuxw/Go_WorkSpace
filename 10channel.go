package main

import (
	"fmt"
)

//通道（channel）是用来传递数据的一个数据结构。管道的本质是一个队列
//作用：协程之间通信
//var 变量名 chan 数据类型
/*
知识点：
	1、chan的缓冲与无缓冲，无缓冲我把他认为是只有一个缓冲
    2、chan的关闭，关闭后只能读不能写
    3、chan的range迭代，阻塞迭代直到关闭
    4、单向队列，收和发分开
    5、chan与select，作用是可以监听多个chan
    6、阻塞的条件：1、未关闭写写满了；2、未关闭读完了
*/

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

	//测试close和不需要sleep
	testData()

	//测试只读和只写的管道
	intE := make(chan int) //   chan   //读写

	go counter(intE) //生产者
	printer(intE)    //消费者

	fmt.Println("done")

	//测试select
	dofibonacci()
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

//测试管道的关闭
func testData() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i //这里会阻塞的写
		}
		//可以用defer
		close(c) //为什么不加close会死锁不知道
	}()

	//可以用for...range简写，下面有例子
	for {
		//ok为true说明channel没有关闭，为false说明管道已经关闭
		if data, ok := <-c; ok { //这里会阻塞的读
			fmt.Println(data)
		} else {
			fmt.Println("chan is closed")
			break
		}
	}

	fmt.Println("Finished")
}

//可以将 channel 隐式转换为单向队列，只收或只发，不能将单向 channel 转换为普通 channel
//   chan<- //只写
func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i //如果对方不读 会阻塞
	}
}

//   <-chan //只读
func printer(in <-chan int) {
	//for …… range语句可以处理Channel。阻塞迭代
	for num := range in { //range in产生的迭代值为Channel中发送的值，它会一直迭代直到channel被关闭
		fmt.Println(num)
	}
}

//测试select，select可以完成监控多个管道
func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: //意思是如果可写
			x, y = y, x+y
		case <-quit: //意思是如果可读
			fmt.Println("quit")
			return
		}
	}
}

func dofibonacci() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
