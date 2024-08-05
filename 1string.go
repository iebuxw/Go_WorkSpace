//包声明
package main

//引入包
import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//var的用法：var identifier type
//make的用法：make(type, length, capacity)  //capacity就是容量，警戒值
func main() {
	//变量的定义
	//var str string   				//方式一，没有初始化就是零值
	//var str string = "yyyyy"   	//方式二，指定类型
	//var str3 = "zzzzzz"			//方式三，类型推导
	str2 := "yyyyyaaa" //推荐，只能在函数内，不能声明全局变量
	fmt.Println("str2= ", str2)

	// 多变量声明
	//var str, str3 int = 10, 20	// 指定类型
	//var str, str3 = 10, 20		// 类型推导
	/*var(
		str3 int = 10				// 指定类型
		str = 10					// 类型推导
	)*/

	//``创建原始字符串字面量，类似php的EOF
	fmt.Println(`dfdretgfgggg`)

	//字符串连接用+
	fmt.Println("Hello " + "World!")

	var stockcode = 123
	var enddate = "2020-12-31"
	var targetUrl = fmt.Sprintf("Code=%d&endDate=%s", stockcode, enddate)
	fmt.Println(targetUrl)

	//byte 等同于int8，即一个字节长度，常用来处理ascii字符
	//rune 等同于int32，即4个字节长度,常用来处理unicode或utf-8字符
	//包含中文的字符串应该先转成rune，再计算长度
	str13 := "你好 world"
	fmt.Printf("len(str):%d\n", len(str13)) //返回len(str):12
	//计算中文的长度
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
	//strlen
	var2 := "12233"
	fmt.Println("var2 len= ", len(var2))

	//字符串转整数intval
	int1, error2 := strconv.Atoi(var2) //方式一
	//方式二10代表基准是10进制，0代表转成int
	var11, _ := strconv.ParseInt(var2, 10, 0)
	fmt.Println(int1, error2, var11) //方式三 sprintf

	//整数转字符串
	var4 := 1000
	var5 := strconv.Itoa(var4)
	//这种也可以 Itoa 等价 FormatInt(int64(var4), 10)，第二个参数10代表10进制
	var10 := strconv.FormatInt(int64(var4), 10)
	fmt.Println(var5, var10)

	//字符串是否包含
	var6 := strings.Contains("wokkkdddssss", "wo")
	fmt.Println("is contains", var6)

	//strpos，不存在则返回-1
	var7 := strings.Index("wokkkdddssss", "ok")
	fmt.Println(var7)

	//str_replace字符串替换,-1代表全部替换
	var8 := strings.Replace("wokkkdddokssss", "ok", "go22", -1)
	fmt.Println(var8)

	//字符串截取，substr
	sstr := "Hello HaiCoder!"
	sstr1 := sstr[0:4] // 包含0不包含4
	sstr2 := sstr[:8]
	sstr3 := sstr[4:9]
	fmt.Println("sstr1 =", sstr1, "sstr2 =", sstr2, "sstr3 =", sstr3)

	//中文截取
	sstr_z := "Hello是大是大非!"
	sstr_r := []rune(sstr_z)
	fmt.Println("中文截取后结果 =", string(sstr_r[1:8])) // 包含1不包含8

	//explode
	var9 := strings.Split("wok,kkd,ddo,kssss", ",")
	fmt.Println("var9 =", var9)

	//implode
	slice := []string{"apple", "banana", "orange"}
	result := strings.Join(slice, ",")
	fmt.Println(result) // 输出：apple,banana,orange

	//strtolower和strtoupper转大小写
	fmt.Println(strings.ToLower("GOOOO"))
	fmt.Println(strings.ToUpper("dssaaa"))

	//trim
	fmt.Printf("str = %q", strings.TrimSpace("  ds   saaa    "))

	//startWith，是否已什么开头
	fmt.Printf("res = %t", strings.HasPrefix("asodddddd", "aso"))

	//endWith，是否已什么结尾
	fmt.Printf("res = %t", strings.HasSuffix("asodddddse", "dse"))

	//日期函数
	//获取当前时间戳
	timeUnix := time.Now().Unix()
	fmt.Println("当前时间为:", timeUnix)

	//获取当前日志，年月日 时分秒
	fmt.Println("年月日时分秒：", time.Now().Format("2006-01-02 15:04:05"))

	//日期转时间戳
	timeStr := "2022-02-01 10:45:15"
	//要转换成时间日期的格式模板（go诞生时间，模板必须是这个时间）
	timeTmeplate := "2006-01-02 15:04:05"
	tim, _ := time.Parse(timeTmeplate, timeStr)
	chuo := tim.Unix()
	fmt.Println("时间戳为：", chuo)

	//时间戳转日期
	tm := time.Unix(chuo, 0)
	timeStr2 := tm.Format(timeTmeplate)
	fmt.Println("日期为：", timeStr2)

	//sleep 2s
	time.Sleep(time.Second * 2)

}
