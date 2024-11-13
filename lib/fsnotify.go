package lib

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)
func main() {
	// 创建一个新的fsnotify监视器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error creating watcher:", err)
	}
	defer watcher.Close()
	// 添加要监视的文件或目录
	err = watcher.Add("D:\\go_code\\Go_WorkSpace")
	if err != nil {
		log.Fatal("Error adding directory to watcher:", err)
	}
	// 启动一个goroutine来处理监视事件
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// 处理文件系统事件
				fmt.Println("Event:", event)
			case err := <-watcher.Errors:
				// 处理错误
				log.Println("Error:", err)
			}
		}
	}()
	// 阻塞主goroutine，以便持续监视文件系统事件
	select {}
}