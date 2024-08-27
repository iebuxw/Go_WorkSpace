package main

import "fmt"

/**

type slice struct {
    array unsafe.Pointer // 元素指针
    len   int // 长度
    cap   int // 容量
}

1、以上就是切片的底层结构，实际上是一个结构体，存了数组的指针，这也是切片是引用类型的原因
2、切片如果扩容，就会跟底层数组解绑，此时不再影响原底层数组，详见下面AddItem

*/

func ChangeFirstItem(lgC []string) {
	lgC[0] = "C"
}

func AddItem(lgC []string) {
	lgC = append(lgC, "Rust")
}

func main() {
	lgA := []string{"C++", "JavaScript", "Python", "PHP"}
	fmt.Println("修改前：", lgA)
	ChangeFirstItem(lgA)		//修改会在函数外部提现
	fmt.Println("修改后：", lgA)

	lgB := []string{"C++", "JavaScript", "Python", "PHP"}
	fmt.Println("添加前：", lgB)
	AddItem(lgA)				//扩容了，发生了复制，指向了新的底层数组
	fmt.Println("添加后：", lgB)
}
