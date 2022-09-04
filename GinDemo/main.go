package main

import (
	"GinDemo/common"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库
	common.InitDB()

	//使用gin
	r := gin.Default()
	r = CollectRoute(r)

	_ = r.Run(":9000")
}
