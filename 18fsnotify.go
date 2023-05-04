package main

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)

func main() {
	// 创建文件/目录监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("watcher new err: [%+v]", err)
		return
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// 打印监听事件
				log.Print("event:", event)
				php /opt/www/yiic evidschedule
				// 可以启动新的goroutine或使用channel进行事件传递
			case _, ok := <-watcher.Errors:
				if !ok {
					return
				}
			}
		}
	}()
	// 监听当前目录
	err = watcher.Add("D:\\Users\\xiaowang\\桌面\\tmp2\\web_client")
	if err != nil {
		log.Fatal("err add dir: [%+v]", err)
	}
	<-done
}
