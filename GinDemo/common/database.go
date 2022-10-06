package common

import (
	"GinDemo/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
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

	//DB, err = gorm.Open(mysql.Open(args), &gorm.Config{})
	DB, err = gorm.Open(driverName, args)
	if err != nil {
		//log.Fatal(err)
		panic("fail to connect database, err: " + err.Error())
	}

	_ = DB.AutoMigrate(&model.User{})

	fmt.Printf("连接成功：%v\n", DB)

	return DB
}

func GetDb() *gorm.DB {
	return DB
}
