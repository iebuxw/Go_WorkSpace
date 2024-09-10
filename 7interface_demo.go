package main

import "fmt"

type Quacker interface {
	Quack()
}

type Duck struct{}

func (d Duck) Quack() {
	fmt.Println("Quack!")
}

func main() {
	duck := Duck{}
	// 可以将Duck的实例赋值给Quacker类型的变量
	var quacker Quacker = duck
	quacker.Quack() // 输出: Quack!
}
