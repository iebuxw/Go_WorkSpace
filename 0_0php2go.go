package main

import (
	"fmt"
	"github.com/syyongx/php2go"
	"github.com/tidwall/gjson"
)

// 文档地址：https://github.com/syyongx/php2go

func main() {
	slice1 := []string{"sss", "nnnn"}
	string2 := php2go.Implode(",", slice1)
	fmt.Println("php的Implode函数结果：", string2)
	ret := php2go.InArray("sss", slice1)
	fmt.Println("php的In_array函数结果：", ret)
	json := `{"name":"ddddd"}`
	value := gjson.Get(json, "name")
	println("php的json_decode结果：", value.String())
}
