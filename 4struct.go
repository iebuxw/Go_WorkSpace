package main

import (
	"encoding/json"
	"fmt"
)

// Books 定义结构体，结构体可以存储不类型的数据.结构体就像是类的一种简化形式
//结构体类型名：直接就是定义好的结构体名称
/*需要固定的定义，用结构体
存临时数据用map*/
type Books struct { //推荐定义方式
	Title   string `json:"title"` //后面是struct tag，相当于序列化后的字段别名
	Author  string
	Subject string
	Book_id int
}

type Books2 struct {
	Title   string `json:"title"` //后面是struct tag，相当于序列化后的字段别名
	Author  string
	Subject string
	Book_id int
}

type Books3 struct {
	Books2
	Page  int
	page2 string
}

//func (var *Struct_Name) FuncName( var0, var1... )  return type {}
//方法和函数很相似，但是概念有所区别
//表示Books结构体有个方法叫getList
//加了*代表引用拷贝，没加就是值拷贝
func (v *Books) getList() {
	fmt.Println("title")
	v.Title = "PHP语言"
	fmt.Println("Title=", v.Title)
}

func (v Books) getList2() {
	v.Title = "PHP语言2"
	fmt.Println("Title2=", v.Title)
}

func main() {
	// 创建一个新的结构体
	b := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407} //推荐定义方式
	fmt.Println(b)
	fmt.Printf("22%v\n", b)
	fmt.Printf("333333%#v\n", b) //打印结构体的详细信息

	// 也可以使用 key => value 格式
	fmt.Println(Books{
		Title:   "Go 语言",
		Author:  "www.runoob.com",
		Subject: "Go 语言教程",
		Book_id: 6495407,
	})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{Title: "Go 语言", Author: "www.runoob.com"})

	//访问成员
	fmt.Println(b.Title)

	//执行结构体方法
	(&b).getList()
	fmt.Println("&b.Title=", b.Title) //b也改变了(PHP语言)

	b.getList2()
	fmt.Println("main b.Title=", b.Title)  //Go 语言
	(&b).getList2()                        //仍然是值拷贝，跟上面等价
	fmt.Println("main2 b.Title=", b.Title) //Go 语言

	//转json：序列化，Marshal函数使用了反射
	jsonStr, _ := json.Marshal(b)
	fmt.Println("json=", string(jsonStr))

	//为啥底下两种写法都可以？容易混淆，注意是等价的
	book2 := Books2{}
	//book2 := &Books2{}
	book2.Title = "java"
	fmt.Println("book2=", book2)

	//结构体继承
	book3 := Books3{book2, 111, "page2"}
	fmt.Println("book3=", book3)
}
