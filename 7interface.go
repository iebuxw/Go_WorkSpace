package main

import "fmt"

// 定义数据库操作的接口
type Database interface {
	Connect()
	Query(sql string) string
}

// 实现 MySQL 数据库操作
type MySQL struct {
	Name  string
}

func (m MySQL) Connect() {
	fmt.Println("Connect to MySQL database")
}

func (m MySQL) Query(sql string) string {
	return "MySQL: " + sql
}

// 实现 PostgreSQL 数据库操作
type PostgreSQL struct {
	Name  string
}

func (p PostgreSQL) Connect() {
	fmt.Println("Connect to PostgreSQL database")
}

func (p PostgreSQL) Query(sql string) string {
	return "PostgreSQL: " + sql
}

//接口：定义规范。把所有的具有共性的方法定义在一起
//接口是隐式实现的，如果一个类型实现了一个接口定义的所有方法，那么它就自动地实现了该接口
//类型断言
//空接口：万能类型，可以接受任何数据
func main() {
	// 使用接口进行数据库操作
	var db Database

	// 使用 MySQL
	db = MySQL{"mysql"}
	db.Connect()
	result := db.Query("SELECT * FROM users")
	fmt.Println(result)

	// 使用 PostgreSQL
	db = PostgreSQL{"postgresql"}
	db.Connect()
	result = db.Query("SELECT * FROM users")
	fmt.Println(result)

	//类型断言进行类型转换
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}

	// 空接口可以接受任何数据，元素看起来是字符串或者int，但类型是空接口
	mixedData := []interface{}{"foobar", 123456}   // 往切片里放不同类型的元素
	fmt.Println(mixedData)

	mixedData2 := make(map[string]interface{})     // 往map里放不同类型的元素
	mixedData2["username"] = "YamiOdymel"
	mixedData2["time"]     = 123456
}