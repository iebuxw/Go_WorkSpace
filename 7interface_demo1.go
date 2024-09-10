package main

import "fmt"

type Book struct {
	Name  string
	Price float64
	Num   int
}

func (b *Book) GetPrice() float64 {//最终实现接口的数据类型是结构体指针
	return b.Price
}

func (b *Book) GetNum() int {
	return b.Num
}

type Phone struct {
	Brand    string
	Discount float64
	Price    float64
	Num      int
}

func (p *Phone) GetPrice() float64 {
	return p.Price * p.Discount
}

func (p *Phone) GetNum() int {
	return p.Num
}

func main() {
	// 底下这个这会导致编译错误
	//var goods2 Goods = Phone{Brand: "华为", Price: 6000, Num: 1, Discount: 0.8}

	// 正确的做法是使用指针，因为最终实现Goods接口的是结构体指针
	// Phone实现了Goods接口，所以可以将&Phone{...}赋值给Goods类型的变量goods
	var goods Goods = &Phone{Brand: "华为", Price: 6000, Num: 1, Discount: 0.8}

	// 当调用 goods.GetPrice() 时，Go语言会在运行时查找 goods 接口变量内部存储的
	// 实际值的类型,（在这个例子中是 *Phone），并调用该类型上对应的GetPrice()方法
	fmt.Println(goods.GetPrice())
}
