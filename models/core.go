package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlInit() {
	//连接数据库
	var err error
	dsn := "root:cyj18859062686.@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//模型映射
	DB.AutoMigrate(&User{}, &Task{})
	if err != nil {
		panic(err)
	}
}
