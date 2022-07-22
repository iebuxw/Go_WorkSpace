//包声明
package main

//引入包
import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//init函数会在main之前执行
//var的用法：var identifier type
//make的用法：make(type, length, capacity)//capacity就是容量，警戒值
func init() {
	fmt.Println("this is init...")
}

//main 函数是每一个可执行程序所必须包含的
func main() {
	fmt.Println("this is main...")

	var str1 string
	str1 = "我是str1"
	fmt.Println("str1= ", str1)

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

	//byte 等同于int8，即一个字节长度，常用来处理ascii字符
	//rune 等同于int32，即4个字节长度,常用来处理unicode或utf-8字符
	//包含中文的字符串应该先转成rune，在计算长度。[]rune应该是个切片
	str13 := "你好 world"
	fmt.Printf("len(str):%d\n", len(str13))               //返回len(str):12
	fmt.Printf("len(rune(str)):%d\n", len([]rune(str13))) //len(rune(str)):8

	str14 := "I am lilei"
	//string 转[]byte
	b := []byte(str14)
	//[]byte转string
	str14 = string(b)
	//string 转 rune
	r := []rune(str14)
	//rune 转 string
	str14 = string(r)

	//字符串函数
	var2 := "12233"
	fmt.Println("var2 len= ", len(var2))
	//strconv
	//字符串转整数
	int1, error2 := strconv.Atoi(var2)        //方式一
	var11, _ := strconv.ParseInt(var2, 10, 0) //方式二10代表基准是10进制，0代表转成int
	//方式三 sprintf
	fmt.Println(int1, error2, var11)
	//整数转字符串
	var4 := 1000
	var5 := strconv.Itoa(var4)
	var10 := strconv.FormatInt(int64(var4), 10) //这种也可以 Itoa 等价 FormatInt(int64(var4), 10)
	fmt.Println(var5, var10)
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

	//日期函数
	//日期时间字符串

	//获取当前时间戳
	timeUnix := time.Now().Unix()
	fmt.Println("当前时间为:", timeUnix)
	//获取年月日 时分秒
	fmt.Println("年月日时分秒：", time.Now().Format("2006-01-02 15:04:05")) // 当前时间
	//sleep 2s
	time.Sleep(time.Second * 2)

	//日期转时间戳
	timeStr := "2022-02-01 10:45:15"
	//要转换成时间日期的格式模板（go诞生时间，模板必须是这个时间）
	timeTmeplate := "2006-01-02 15:04:05"
	tim, _ := time.Parse(timeTmeplate, timeStr)
	chuo := tim.Unix()
	fmt.Println("时间戳为：", chuo)
	//时间戳转日期
	tm := time.Unix(int64(chuo), 0)
	timeStr2 := tm.Format(timeTmeplate)
	fmt.Println("日期为：", timeStr2)
}
