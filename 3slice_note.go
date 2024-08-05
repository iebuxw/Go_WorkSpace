package main

import "fmt"

/**

type slice struct {
    array unsafe.Pointer // 元素指针
    len   int // 长度
    cap   int // 容量
}

1、以上就是切片的底层结构，实际上是一个结构体，存了数组的指针，这也是切片是引用类型的原因
2、切片如果扩容，就会跟底层数组解绑，此时不再影响原底层数组，详见下面exchangeSlice2

 */

func main() {
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)
	exchangeSlice(slice)
	fmt.Println(slice)     // 函数里面改变反馈到外部

	slice2 := []int{1,2,3}
	fmt.Println(slice2)
	exchangeSlice2(slice2)
	fmt.Println(slice2)   // 扩容了，与底层数组解绑，内部修改不反馈到外部
}

func exchangeSlice(slice []int) {
	for k, v := range slice {
		slice[k] = v * 2
	}
}

func exchangeSlice2(slice []int) {
	slice = append(slice, 4, 5, 6) // 原始数组大小为4，超过4会扩容
	for k, v := range slice {
		slice[k] = v * 2
	}
	fmt.Println(slice)
}
