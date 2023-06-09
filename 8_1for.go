//包声明
package main

import "fmt"

func main() {
	// for
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// forEach
	strings := []string{"google", "runoob"}
	//for _, s := range strings {            // i或者s可以省略
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
