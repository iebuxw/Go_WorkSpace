package main

// flag 包中，提供了命令行参数解析的功能。
//flag.Type(flag 名, 默认值, 帮助信息) *Type
import (
	"flag"
	"fmt"
)

//go run 15flag.go -name="李四" -age=18 -married=true
//go run 15flag.go -name "李四" -age 18
func main() {
	// 定义命令行参数
	var name string
	var age int
	var married bool
	// 使用 flag.TypeVar 方法
	flag.StringVar(&name, "name", "张三", "名字信息")
	flag.IntVar(&age, "age", 0, "年龄信息")
	flag.BoolVar(&married, "married", false, "婚恋情况")

	// 命令行参数解析
	flag.Parse()

	// 返回的是指针，访问需使用*
	fmt.Println(name, age, married) // 李四 18 false
}
