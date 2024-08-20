//包声明
package main

import (
	"fmt"
	"time"
)

//main 函数是每一个可执行程序所必须包含的
func main() {
	if true {
		fmt.Println("222222")
	} else if 1 == 2 {
		fmt.Println("33333")
	} else {
		fmt.Println("44444")
	}

	// 条件前可以写一个语句，不能多个
	// 该语句声明的变量作用域仅在 if 之内，或者对应 else 块
	if v := 1; v < 10 {
		fmt.Println("44444")
	} else {
		fmt.Println(v)
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

	// 没有条件的switch，相当于if-then-else
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
	}
}
