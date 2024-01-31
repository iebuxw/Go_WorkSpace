package main

import (
	"fmt"
)

func main() {
	// 创建切片
	slice := []int{1, 2, 10, 5}

	// in_array
	found := false
	for _, value := range slice {
		if value == 10 {
			found = true
			break
		}
	}
	fmt.Println("10在切片里:", found)

	// count
	fmt.Println("切片个数:", len(slice))

	// array_unshift
	slice = append([]int{100}, slice...)
	fmt.Println("头部加一个元素:", slice)

	// array_shift
	slice = slice[1:]
	fmt.Println("头部去掉一个元素:", slice)

	// array_push
	slice = append(slice, 6)
	fmt.Println("尾部加一个元素:", slice)

	// array_pop
	slice = slice[:len(slice)-1]
	fmt.Println("尾部去掉一个元素:", slice)

	// array_merge
	a3 := []int{1, 2, 10, 5}
	a4 := []int{11, 2, 101, 51}
	combined := append(a3, a4...)
	fmt.Println("合并两个切片:", combined)

	// array_diff
	a5 := []int{1, 2, 10, 5}
	a6 := []int{11, 2, 101, 51}
	diff := make([]int, 0)
	for _, v := range a5 {
		found := false
		for _, v2 := range a6 {
			if v == v2 {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, v)
		}
	}
	fmt.Println("两个切片差集:", diff)

	// array_intersect
	a7 := []int{1, 2, 101, 5}
	a8 := []int{11, 2, 101, 51}
	intersect := make([]int, 0)
	for _, v := range a7 {
		for _, v2 := range a8 {
			if v == v2 {
				intersect = append(intersect, v)
				break
			}
		}
	}
	fmt.Println("两个切片交集:", intersect)

	// array_unique
	a10 := []int{1, 2, 2, 10, 5}
	// 创建一个map来存储唯一元素
	uniqueMap := make(map[int]bool)
	var uniqueSlice []int

	for _, v := range a10 {
		if _, ok := uniqueMap[v]; !ok {
			uniqueMap[v] = true
			uniqueSlice = append(uniqueSlice, v)
		}
	}

	// 输出去重后的切片
	fmt.Println(uniqueSlice)
}