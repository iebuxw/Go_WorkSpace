package main

import (
	"log"
	"time"
)

func main() {
	// 用defer func代替defer
	start := time.Now()
	log.Printf("开始时间为：%v", start)
	// Now()此时已经copy进去了
	// defer log.Printf("时间差：%v", time.Since(start))
	defer func() {
		log.Printf("时间差：%v", time.Since(start))
	}()
	//不受这3秒睡眠的影响
	time.Sleep(3 * time.Second)
	log.Printf("函数结束")
}
