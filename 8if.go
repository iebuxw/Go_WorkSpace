//包声明
package main

import "fmt"

//main 函数是每一个可执行程序所必须包含的
func main() {
	if true {
		fmt.Println("222222")
	} else if 1 == 2 {
		fmt.Println("33333")
	} else {
		fmt.Println("44444")
	}

	// switch语句
	int2 := 10
	switch int2 {
		case 0:
			fmt.Println("int2= 0")
		case 10:
			fmt.Println("int2= 10")
		default:
			fmt.Println("int2= def")
	}
}
