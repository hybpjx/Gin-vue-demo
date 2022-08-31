package main

import (
	"GinDemo/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()


	r := gin.Default()
	r = CollectRoute(r)

	_ = r.Run(":9000")
}
