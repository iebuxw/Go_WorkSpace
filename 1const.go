package main

import "fmt"

const (
	BEIJING  = 0
	SHANGHAI = 1
)

// iota初始值=0，每行加1
// 只能出现在const的括号中
const (
	AAA = iota
	BBB
)

// 如果没有指定公式，则和上一行保持一致，直到改变公式
const (
	DD = iota * 10 // iota=0, DD=0
	EE             // iota=1, EE=10
	FF = iota + 1  // iota=2, FF=3
)

func main() {
	const age int = 10

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)

	fmt.Println("age = ", age)

	fmt.Println("AAA = ", AAA)
	fmt.Println("BBB = ", BBB)

	fmt.Println("DD = ", DD)
	fmt.Println("EE = ", EE)
	fmt.Println("FF = ", FF)
}
