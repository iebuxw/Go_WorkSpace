package main

import "fmt"

//“动态数组”

//声明切片    var identifier []type
//初始化切片
//访问切片
//切片截取
//切片扩充
//切片元素删除（头部、中间、尾部，利用切片本身的特性：类似[1:]）
//切片类型名：[]type
func main() {
	//定义切片1
	//声明后需要make初始化的类型：切片、map、chan
	var countryCapitalSlice []string
	countryCapitalSlice = make([]string, 2, 10)
	countryCapitalSlice[0] = "ddddd"
	fmt.Println(countryCapitalSlice)

	//定义切片2
	//声明一个未指定大小的数组来定义切片（自动扩容会有性能问题，所以最好指定合理容量）
	//这时字面量（Literals）的方式
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	//定义切片3(推荐)
	slice2 := make([]int, 2, 10)// 2是长度，10是容量
	//切片赋值
	slice2[0] = 1
	slice2[1] = 2
	fmt.Println(slice2)

	//切片截取，数组也可以这样截取
	//注意这种截取会共享底层数组，number2的改变会引起number的改变
	number2 := numbers[1:6] // 索引1(包含) 到索引 6(不包含)
	number3 := numbers[:3]
	number4 := numbers[3:]
	number5 := numbers[:]   // 取所有，可以用来复制

	fmt.Println("number2 == ", numbers)
	fmt.Println("number2 == ", number2)
	fmt.Println("number3 == ", number3)
	fmt.Println("number4 == ", number4)
	fmt.Println("number5 == ", number5)

	//切片扩充
	c := append(number2, 5)
	fmt.Println("number2 == ", number2)
	fmt.Println("c == ", c)

	//删除中间一个元素 [1,2,3,4,5]  => [1,2,4,5]
	idx := 2
	a2 := append(number2[:idx], number2[idx+1:]...)
	fmt.Println("a2 == ", a2)

	// 复制，避免引用赋值
	// 方法1：使用copy函数创建独立副本
	number22 := make([]int, 5)
	copy(number22, numbers[1:6])

	// 方法2：使用append创建独立副本
	// number22 := append([]int{}, numbers[1:6]...)

	/**
		切片不要和nil比较，判断切片是否为空要用len(s) == 0来判断
		var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
		s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
		s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	 */
	s3 := make([]int, 0)
	//s3这个切片并不是 nil，而是指向了一个长度和容量都为0的底层数组
	//if s3 == nil {// 这种判断有问题
	if len(s3) == 0 {
		fmt.Println("fffff")
	}

	var s []int
	var ss = []int{}
	fmt.Println(len(s),len(ss))//0,0
	fmt.Println(s == nil) //true
	fmt.Println(ss == nil)//false

	// 会panic，类似js的Cannot read property ... of undefined
	s4 := make([]int, 0, 4)
	s4[0] = 1
}
