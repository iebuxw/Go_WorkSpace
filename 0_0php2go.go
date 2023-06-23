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
	json2 := `{"name":["ddddd"]}`
	json3 := `{"name":[{"ddddd":"dddggggggg"}]}`
	value := gjson.Get(json, "name")
	fmt.Println("json的json_decode结果：", value.String())
	fmt.Println("json2的json_decode结果：", gjson.Get(json2, "name"))
	fmt.Println("json3的json_decode结果：", fmt.Sprintf("%#v", gjson.Get(json3, "name")))
}
