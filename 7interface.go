package main

type sw interface {
	//定义交换机公共方法
	login()
}

type Huawei struct {
	//定义华为结构体
}

type Zte struct {
	//定义中兴结构体
}

func (*Huawei) login() {
	//华为交换机登录的公共方法
}

func (*Huawei) sfunc2() {
	//华为交换机的特殊方法
}

func (*Zte) login() {
	//zte交换机登录的公共方法
}

func (*Zte) sfunc1() {
	//zte交换机的特殊方法
}

func NewSwHuawei() sw {
	//华为结构体初始化，返回值是sw接口
	return &Huawei{}
}

func NewZte() sw {
	//中兴结构体初始化，返回值是sw接口
	return &Zte{}
}

func LoginSw(s sw) {
	//某个交换机登录的公共函数
	switch s.(type) {
	case *Zte:
		s.login()
		s.(*Zte).sfunc1()
	case *Huawei:
		s.login()
		s.(*Huawei).sfunc2()
	}
}
