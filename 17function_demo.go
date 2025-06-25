package main

import "fmt"

func main() {
	a, b := test1(1, 2)
	fmt.Println(a, b)

	a, b = test2(3, 4)
	fmt.Println(a, b)
}

func test1(a int, b int) (int, int) {
	a1 := a + 1
	b2 := b + 1
	return a1, b2
}

// 带参数的返回值
func test2(a int, b int) (a1 int, b2 int) {
	a1 = a + 1
	b2 = b + 1
	return
}

// 参数是切片、map直接值传递就行
// 修改元素（影响原切片）
func UpdateSlice(s []int) {
	s[0] = 100  // 直接影响原slice
}

// 扩容操作（需返回值）
func AppendSlice(s []int) []int {
	return append(s, 200)
}

func UpdateMap(m map[string]int) {
	m["key"] = 100  // 直接影响原map
}

// 参数是struct，只读时值传递
type Point struct{ X, Y int }
func ReadPoint(p Point) int {
	return p.X
}

// 参数是大struct，写时传指针
type BigStruct struct { data [1024]byte }
func UpdateBig(s *BigStruct) {
	s.data[0] = 1  // 避免复制开销
}

// ‌核心原则‌：优先使用值传递，仅在修改大结构体时用指针
