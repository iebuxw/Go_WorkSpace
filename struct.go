package main

import "fmt"

// Books 定义结构体，结构体可以存储不类型的数据.结构体就像是类的一种简化形式
/*需要固定的定义，用结构体
存临时数据用map*/
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

//func (var *Struct_Name) FuncName( var0, var1... )  return type {}
func (v *Books) getList() {
	fmt.Println("title")
	fmt.Println(v.title)
}

func main() {
	// 创建一个新的结构体
	b := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	fmt.Println(b)

	// 也可以使用 key => value 格式
	fmt.Println(Books{
		title:   "Go 语言",
		author:  "www.runoob.com",
		subject: "Go 语言教程",
		book_id: 6495407,
	})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	//访问成员
	fmt.Println(b.title)

	//执行结构体方法
	b.getList()
}
