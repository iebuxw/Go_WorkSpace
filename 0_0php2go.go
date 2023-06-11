package main

import (
	"fmt"
	"github.com/syyongx/php2go"
)

// 文档地址：https://github.com/syyongx/php2go

func main() {
	slice1 := []string{"sss", "nnnn"}
	string2 := php2go.Implode(",", slice1)
	fmt.Println("php的Implode函数结果：", string2)
}
