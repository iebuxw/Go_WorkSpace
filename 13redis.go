package main

//需要设置gopath
//在gopath下执行 go get github.com/garyburd/redigo/redis
import (
	"fmt"
	"github.com/garyburd/redigo/redis" //不同的包有不同的api，按文档使用就好了
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("连接redis失败")
		return
	}
	fmt.Println("连接redis成功")
	//需要关闭
	defer func(c redis.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Println("关闭失败")
		}
	}(c)
	_, err2 := c.Do("Set", "goTestKey", "hello")
	if err2 != nil {
		fmt.Println("set失败")
		return
	}
	//返回的是接口类型，需要转成string
	goTestKey, _ := redis.String(c.Do("Get", "goTestKey"))

	fmt.Println("获取成功,goTestKey:", goTestKey)

	PoolTest()
}

// PoolTest 连接池方式
func PoolTest() {
	var pool *redis.Pool //Pool是一个结构体，为什么要定义为指针，因为这个包的Get函数跟指针绑定的

	pool = &redis.Pool{ //&符号是指针操作，然后初始化结构体
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 200,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	//从pool中取出一个连接
	conn2 := pool.Get()
	//需要关闭
	defer func(conn2 redis.Conn) {
		err := conn2.Close()
		if err != nil {
			fmt.Println("关闭失败2")
		}
	}(conn2)

	_, err3 := conn2.Do("Set", "goTestKey2", "world")
	if err3 != nil {
		fmt.Println("set2失败")
		return
	}
	goTestKey2, _ := redis.String(conn2.Do("Get", "goTestKey2"))

	fmt.Println("获取成功,goTestKey2:", goTestKey2)
}
