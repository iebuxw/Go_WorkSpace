package main

import "fmt"

func main() {
	//定义切片1
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8} //声明一个未指定大小的数组来定义切片

	//切片截取
	number2 := numbers[1:6]
	number3 := numbers[:3]
	number4 := numbers[3:]

	fmt.Println("number2 == ", numbers)
	fmt.Println("number2 == ", number2)
	fmt.Println("number3 == ", number3)
	fmt.Println("number4 == ", number4)
	c := append(number2, 5)
	fmt.Println("number2 == ", number2)
	fmt.Println("c == ", c)

	//定义切片2
	slice2 := make([]int, 2)
	slice2[0] = 1
	slice2[1] = 2
	fmt.Println(slice2)
}
