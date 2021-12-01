package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	number2 := numbers[1:6]
	number3 := numbers[:3]
	number4 := numbers[3:]

	fmt.Println("number2 == ", number2)
	fmt.Println("number3 == ", number3)
	fmt.Println("number4 == ", number4)
	append(number2, 5)
	fmt.Println("number2 == ", number2)
	fmt.Println("c == ", c)

	copy(number2, number3)
	fmt.Println("number2 == ", number2)

}
