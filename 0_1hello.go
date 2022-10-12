package main

import "fmt"

// 在main之前调用
func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("Hello World!")
}
