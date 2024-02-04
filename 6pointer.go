//包声明
package main

import "fmt"

//指针类型名，可以看做是跟int、string平级的类型，
//语法：var var_name *var-type
//指针主要是减少变量的拷贝，提供便捷性。php的引用传递是指针的简化版本
//var ip *int        /* 指向整型*/
//var fp *float32    /* 指向浮点型 */
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
