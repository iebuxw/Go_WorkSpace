package main

import "fmt"

//声明数组
//初始化数组
//访问数组
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
	m := [...]string{"aaaa", "bbbb", "ccc"} //数组
	for key, value := range m {
		fmt.Println(key, "=>", value)
	}

	//忽略索引
	for _, value := range m {
		fmt.Println(value)
	}

	//数组函数
	//count
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
