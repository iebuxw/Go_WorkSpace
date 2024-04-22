package main

import (
	"fmt"
)

/*
func function_name( [parameter list] ) [return_types] {
   函数体
}

不支持默认参数
*/
func main() {
	/* 定义局部变量 */
	var a = 100
	var b = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf( "最大值是 : %d\n", ret )
}

func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}