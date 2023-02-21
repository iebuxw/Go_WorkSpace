//包声明
package main

import "fmt"

//main 函数是每一个可执行程序所必须包含的
//指针类型名，可以看做是跟int、string平级的类型，在函数参数类型限制那里可以填写指针类型: *type
//如：*int、*string、*Person，Person是定义的一个结构体类型名称
func main() {
	//定义指针
	//指针赋值
	//指针的使用
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var b int
	test(&b) //类似php里的引用
	fmt.Printf("b变量的值: %d\n", b)

}

func test(p *int) {
	*p = 10
}
