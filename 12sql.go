package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	//gorm.Model
	Id   int32
	Name string
	Age  int32
}

func main() {
	// https://github.com/go-gorm/postgres
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=testdrive port=5360 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	//db.AutoMigrate(&Student{})

	// Create
	db.Create(&Student{Name: "D43", Age: 101})

	// Read
	var student Student
	var student2 Student
	db.First(&student, 25) // 根据整型主键查找
	fmt.Println("学生id=25的信息", student)
	db.First(&student2, "name = ?", "D42") // 查找 name 字段值为 D42 的记录
	fmt.Println("学生name=D42的信息", student2)
	fmt.Println("学生name=D42的年龄", student2.Age)

	// Update - 将 student 的 age 更新为 200
	db.Model(&student).Update("Age", 206)
	// Update - 更新多个字段
	//db.Model(&student).Updates(Student{Age: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&student).Updates(map[string]interface{}{"Age": 200, "Code": "F42"})

	// Delete - 删除 student
	db.Delete(&student, 2)
}
