package main

import "fmt"

func main() {
	a, b := test1(1, 2)
	fmt.Println(a, b)
}

func test1(a int, b int) (a1 int, b2 int) {
	a4 := a + 1
	b4 := b + 1
	return a4, b4
}
