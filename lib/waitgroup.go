package lib

import (
"fmt"
"sync"
"time"
)

func main() {
	// 创建一个等待组
	var wg sync.WaitGroup

	// 创建一个通道用于控制最大并发数
	sem := make(chan struct{}, 3)

	// 创建10个任务
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 将任务添加到通道中
		sem <- struct{}{}
		go func(i int) {
			// 执行任务
			fmt.Println("Task", i, "start")
			time.Sleep(time.Second)
			fmt.Println("Task", i, "done")
			// 任务完成后从通道中移除
			<-sem
			// 完成任务
			wg.Done()
		}(i)
	}
	// 等待所有任务完成
	wg.Wait()
}