package main

import (
	"GinDemo/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/register", controller.Register)
	r.POST("/api/login", controller.Login)
	return r
}
