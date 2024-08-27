package main

import "fmt"

type Goods interface {
	GetPrice() float64
}

type Book struct {
	Price float64
}

// 注意：GetPrice 方法有一个指针接收器
func (b *Book) GetPrice() float64 {//最终实现接口的数据类型是结构体指针
	return b.Price
}

func main() {
	var b1 Book
	b1.Price = 10.0
	// 尝试直接使用 b1（错误）
	// var g Goods = b1 // 这会导致编译错误

	// 正确的做法是使用指针，因为最终实现Goods接口的是结构体指针
	// Goods类型的变量等于一个指针，感觉很怪，先这样，以后再理解
	var g Goods = &b1 // 这不会导致编译错误
	g1 := &b1         // 这样也行
	fmt.Println(g.GetPrice()) // 输出: 10
	fmt.Println(g1.GetPrice()) // 输出: 10
}
