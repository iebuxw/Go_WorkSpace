package main

import "fmt"

// 接口就是定义规范
// 不过貌似在代码中，interface部分就算注释了也能运行，interface和struct没有直接绑定
type sw interface {
	//定义交换机公共方法
	login()
	logout()
}

type Huawei struct {
	//定义华为结构体
	Name string
}

type Zte struct {
	//定义中兴结构体
	Name string
}

func (c Huawei) login() {
	//华为交换机登录的公共方法
	fmt.Println(c.Name)
}

func (c Zte) login() {
	//zte交换机登录的公共方法
	fmt.Println(c.Name)
}

func main() {
	c := Huawei{
		Name: "huawei",
	}
	c.login()

	d := Zte{
		Name: "zte",
	}
	d.login()
}

//万能数据类型  interface{}
//怎么判断传进来的是什么类型？用类型断言，根据不同的类型进行不同的处理。作用：判断并转换空接口实际对应的类型
//value, ok := s.(string)		//类型断言判断是否是string
