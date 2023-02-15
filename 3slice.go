package main

import "fmt"

//声明切片
//初始化切片
//访问切片
//切片截取
//切片扩充
//切片元素删除
//切片类型名：[]type
func main() {
	//定义切片1
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8} //声明一个未指定大小的数组来定义切片（这种定义方式好用，推荐）

	//切片截取，数组也可以这样截取
	number2 := numbers[1:6] // 索引1(包含) 到索引 6(不包含)
	number3 := numbers[:3]
	number4 := numbers[3:]
	number5 := numbers[:] // 取所有

	fmt.Println("number2 == ", numbers)
	fmt.Println("number2 == ", number2)
	fmt.Println("number3 == ", number3)
	fmt.Println("number4 == ", number4)
	fmt.Println("number5 == ", number5)

	//切片扩充
	c := append(number2, 5)
	fmt.Println("number2 == ", number2)
	fmt.Println("c == ", c)

	//定义切片2
	slice2 := make([]int, 2)
	//切片赋值
	slice2[0] = 1
	slice2[1] = 2
	fmt.Println(slice2)
}
