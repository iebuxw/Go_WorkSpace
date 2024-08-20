package main

import "fmt"

//固定长度、固定类型
//声明数组   var arrayName [size]dataType
//初始化数组
//访问数组
//数组类型名：[size]type
//注意：数组的大小是类型的一部分，[5]int 和 [10]int 是不同的类型
func main() {
	//声明数组
	//var m = [5]int{1, 2, 3, 4, 5}			//方式一
	m := [...]string{"aaaa", "bbbb", "ccc"} //方式二（这种定义方式好用，推荐）

	//foreach
	for key, value := range m {
		fmt.Println(key, "=>", value)
	}

	//忽略索引
	for _, value := range m {
		fmt.Println(value)
	}

	//php数组函数
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