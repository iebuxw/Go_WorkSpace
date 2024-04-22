package main

import (
	"encoding/json"
	"fmt"
)

//特点：1、固定数量；2、键不能数字；3、值类型可以不同
//区别于数组，结构体可以存储不类型的数据
//结构体就像是类的一种简化形式
//结构体是一种自定义数据类型
//可以使用结构体或者结构体指针访问结构体成员，结构体.成员名 | 结构体指针.成员名
type Books struct { //推荐定义方式
	Title   string `json:"title"` //后面是struct tag，相当于序列化后的字段别名
	Author  string
	Subject string
	Book_id int
}

type Books3 struct {
	Books            // 继承
	Page  int
	page2 string
}

//func (var *Struct_Name) FuncName( var0, var1... )  return type {}
//方法和函数很相似，多了个接受类型，也就是上边的(var *Struct_Name)
//加了*代表引用拷贝，没加就是值拷贝
func (v *Books) getList() {
	v.Title = "PHP语言"
	fmt.Println("Title=", v.Title)
}

func (v Books) getList2() {
	v.Title = "PHP语言2"
	fmt.Println("Title2=", v.Title)
}

// 结构体指针作为参数
func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.Title) //指针.属性也是可以的
}

func main() {
	// 创建一个新的结构体，推荐定义方式
	b := Books{
		Title:   "Go 语言",
		Author:  "www.runoob.com",
		Subject: "Go 语言教程",
		Book_id: 6495407,
	}
	//访问成员
	fmt.Println(b.Title)

	//也可以用简化形式
	b2 := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	fmt.Println(b2)
	fmt.Printf("22%v\n", b2)
	fmt.Printf("333333%#v\n", b2) //打印结构体的详细信息

	//执行结构体方法
	(&b).getList()
	fmt.Println("&b.Title=", b.Title) //b也改变了(PHP语言)

	b.getList2()
	fmt.Println("main b.Title=", b.Title)  //Go 语言
	(&b).getList2()         //仍然是值拷贝，跟上面等价，go语言底层优化了，方便书写
	fmt.Println("main2 b.Title=", b.Title) //Go 语言

	//转json：序列化，Marshal函数使用了反射
	jsonStr, _ := json.Marshal(b)
	fmt.Println("json=", string(jsonStr))

	//解析json
	jsonDec := Books{}
	err := json.Unmarshal(jsonStr, &jsonDec)
	fmt.Printf("json decode=%#v err=%s\n", jsonDec, err)

	//结构体指针
	testp := &Books3{}
	//var testp *Books3 = new(Books3)  //跟上面的等价
	testp.Page = 10 // 等价 (*testp).Page = 10
	fmt.Println("testp", testp)

	/* 打印 Book 信息 */
	printBook(&b)

}
