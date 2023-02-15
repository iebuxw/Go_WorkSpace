package main

import "fmt"

//声明数组
//初始化数组
//访问数组
//数组类型名：[size]type
func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	fmt.Println("wang di xin ku kan yi xia")
	fmt.Println(n)

	//foreach
	m := [...]string{"aaaa", "bbbb", "ccc"} //数组（这种定义方式好用，推荐）
	for key, value := range m {
		fmt.Println(key, "=>", value)
	}

	//忽略索引
	for _, value := range m {
		fmt.Println(value)
	}

	arr1 := []int{3, 3}
	arr2 := [5]int{3, 3}
	arr3 := [5]int{3, 3}
	testArr01(arr1)
	fmt.Println(arr1)
	testArr02(arr2)
	fmt.Println(arr2)
	testArr03(&arr3)
	fmt.Println(arr3)

	fmt.Println("切片从数组取值", arr3[1:2]) //注意只有一个元素，包含头不包含尾

	//数组函数
	//count			=>  len()
	//in_array
	//implode
	//array_push
	//array_shift
	//array_unshift
	//array_pop
	//array_map
	//array_diff
	//array_intersect
	//array_col
	//array_merge
	//array_combine
	//array_multisort
	//array_sum
	//array_unique
	//isset
	//array_values
	//array_keys
	//array_search
	//array_slice
	//array_filter
	//sort
}

//数组作为函数的参数
func testArr01(arr []int) { //形参未指定大小(这里应该是切片)
	arr[1] = 4
	fmt.Println(arr)
}
func testArr02(arr [5]int) { //形参指定大小
	arr[1] = 4
	fmt.Println(arr)
}
func testArr03(arr *[5]int) { //使用指针方式，也就是引用传递
	//arr[1] = 5
	(*arr)[1] = 4
	fmt.Println(*arr)
}
