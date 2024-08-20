package main

import "fmt"

// 定义一个泛型函数，它接受两个同类型的参数并返回它们的和
// 可以加约束，例如：any、comparable，comparable是可以用== and !=比较
func Sum[T int|float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Sum(1, 2))      // int 类型的参数
	fmt.Println(Sum(1.1, 2.2))  // float64 类型的参数
}