package common

import (
	"GinDemo/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//driverName := "mysql"
	host := "127.0.0.1"
	port := 3306
	username := "root"
	password := "admin*123"
	database := "ginInessential"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	
	var err error

	DB, err = gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	_ = DB.AutoMigrate(&model.User{})

	fmt.Printf("连接成功：%v\n", DB)

	return DB
}

func GetDb() *gorm.DB {
	return DB
}
