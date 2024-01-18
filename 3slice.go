package main

import "fmt"

//声明切片
//初始化切片
//访问切片
//切片截取
//切片扩充
//切片元素删除（头部、中间、尾部，利用切片本身的特性：类似[1:]）
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

	//删除中间一个元素 [1,2,3,4,5]  => [1,2,4,5]
	idx := 2
	a2 := append(number2[:idx], number2[idx+1:]...)
	fmt.Println("a2 == ", a2)

	//复制，避免引用赋值
	a3 := make([]int, 4)
	copy(a3, a2)
	a3[2] = 10
	fmt.Println("a3 == ", a3)

	//定义切片2
	slice2 := make([]int, 2)
	//切片赋值
	slice2[0] = 1
	slice2[1] = 2
	fmt.Println(slice2)
}
