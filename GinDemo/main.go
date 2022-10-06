package main

import (
	"GinDemo/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//初始化配置
	InitConfig()

	//初始化数据库
	common.InitDB()

	//使用gin
	r := gin.Default()
	r = CollectRoute(r)

	port := viper.GetString("server.port")

	if port != "" {
		_ = r.Run(":" + port)
	} else {
		err := r.Run()
		if err != nil {
			return 
		}
	}

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
