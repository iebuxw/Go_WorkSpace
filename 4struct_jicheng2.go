package main

import "fmt"

// 定义基础结构体
type BaseStruct struct {
	field1 string
}

// 定义一个在基础结构体中的方法
func (bs *BaseStruct) BaseMethod() {
	fmt.Println("This is a method in BaseStruct.")
}

// 定义子结构体，嵌入基础结构体
type SubStruct struct {
	BaseStruct           // 嵌入基础结构体
	field2     int
}

func main() {
	// 创建子结构体实例
	sub := SubStruct{
		BaseStruct: BaseStruct{
			field1: "Inherited field",
		},
		field2: 10,
	}

	// 直接访问基类字段
	fmt.Println(sub.field1) // 输出: Inherited field

	// 调用基类方法
	sub.BaseMethod() // 输出: This is a method in BaseStruct.
}