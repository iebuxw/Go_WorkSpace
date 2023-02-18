package main

import "fmt"

func main() {
	a, b := test1(1, 2)
	fmt.Println(a, b)

	a, b = test2(3, 4)
	fmt.Println(a, b)
}

func test1(a int, b int) (int, int) {
	a1 := a + 1
	b2 := b + 1
	return a1, b2
}

// 带参数的返回值
func test2(a int, b int) (a1 int, b2 int) {
	a1 = a + 1
	b2 = b + 1
	return
}
