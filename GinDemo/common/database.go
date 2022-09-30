package common

import (
	"GinDemo/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//driverName := "mysql"
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	loc:= viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
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
