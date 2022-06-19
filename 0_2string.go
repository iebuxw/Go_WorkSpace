//包声明
package main

//引入包
import (
	"fmt"
	"strconv"
	"strings"
)

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

	//字符串函数
	var2 := "12233"
	fmt.Println("var2 len= ", len(var2))
	//strconv
	//字符串转整数
	int1, error2 := strconv.Atoi(var2)
	fmt.Println(int1, error2)
	//整数转字符串
	var4 := 1000
	var5 := strconv.Itoa(var4)
	fmt.Println(var5)
	//字符串是否包含
	var6 := strings.Contains("wokkkdddssss", "ok")
	fmt.Println(var6)
	//strpos，不存在则返回-1
	var7 := strings.Index("wokkkdddssss", "ok")
	fmt.Println(var7)
	//字符串替换,-1代表全部替换
	var8 := strings.Replace("wokkkdddokssss", "ok", "go22", -1)
	fmt.Println(var8)
	//字符串截取
	//implode
	//explode
	var9 := strings.Split("wok,kkd,ddo,kssss", ",")
	fmt.Println("var9 =", var9)
	//转大小写
	fmt.Println(strings.ToLower("GOOOO"))
	fmt.Println(strings.ToUpper("dssaaa"))
	//trim
	fmt.Printf("str = %q", strings.TrimSpace("  ds   saaa    "))
	//是否已什么开头
	fmt.Printf("res = %t", strings.HasPrefix("asodddddd", "aso"))
	//是否已什么结尾
	fmt.Printf("res = %t", strings.HasSuffix("asodddddse", "dse"))
}
