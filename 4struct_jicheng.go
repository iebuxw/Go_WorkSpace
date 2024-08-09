package main

import "fmt"

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	// 注意跟Animal  *Animal这样写的区别，一个是继承，一个是属性
	// 注意Animal这样写的区别，一个是值类型，一个是指针类型
	*Animal        //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}

	d1.move()
	d1.Animal.move()     // 显示调用也可以
}


