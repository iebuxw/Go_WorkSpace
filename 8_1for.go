//包声明
package main

import "fmt"

func main() {
	// for
	sum := 0
	for i := 0; i <= 10; i++ {// i只能在for或者else内部使用
		sum += i
	}
	fmt.Println(sum)

	// forEach
	strings := []string{"google", "runoob"}
	//for _, s := range strings {            // i或者s可以省略
	//for i := range strings {               // 可以只有i
	for i, s := range strings {
		fmt.Println(i, s)
	}

	// for range会对slice元素值一次拷贝到item。更改item中的值不会改变原slice
	// 而且item的地址是永远不变的
	slice := []int{1,2,3}
	for _, item := range slice {
		item++
	}
	fmt.Println(slice)

	// while
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// 无限循环
	sum3 := 1
	for {
		sum3++
	}
}
