//包声明
package main

//引入包
import "fmt"

//main 函数是每一个可执行程序所必须包含的
func main() {
	var str string = "yyyyy"
	str2 := "yyyyyaaa"
	int2 := 10
	var str3 = "zzzzzz"

	fmt.Println("str= ", str)
	fmt.Println("str2= ", str2)
	fmt.Println("str3= ", str3)

	//字符串连接
	fmt.Println("Hello " + "World!")

	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(targetUrl)

	switch int2 {
	case 0:
		fmt.Println("int2= 0")
	case 10:
		fmt.Println("int2= 10")
	default:
		fmt.Println("int2= def")
	}
}
