package main

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"path/filepath"
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
				if event.Op&fsnotify.Create == fsnotify.Create {
					// 获取新文件完整路径
					filePath, err := filepath.Abs(event.Name)
					if err != nil {
						log.Errorf("Error: %v", err)
						continue
					}

					// 执行shell命令
					cmd := exec.Command("php", "/opt/www/yiic", "handleevidnotice", "--file="+filePath)
					err = cmd.Run()
					if err != nil {
						log.Errorf("Error: %v", err)
						continue
					}
					log.Println("Executed shell command for", filePath)
				}
				// 可以启动新的goroutine或使用channel进行事件传递
			case _, ok := <-watcher.Errors:
				if !ok {
					return
				}
			}
		}
	}()
	// 监听当前目录
	err = watcher.Add("/data/alarm")
	if err != nil {
		log.Fatal("err add dir: [%+v]", err)
	}
	<-done
}
