package main

import (
	"GinDemo/controller"
	"GinDemo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.DELETE(":id", categoryController.Delete)
	categoryRoutes.PUT(":id", categoryController.Put)
	categoryRoutes.GET(":id", categoryController.Select)
	// 局部修改
	//categoryRoutes.PATCH("")

	postRoutes := r.Group("/posts")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.DELETE(":id", postController.Delete)
	postRoutes.PUT(":id", postController.Put)
	postRoutes.GET(":id", postController.Select)
	postRoutes.POST("page/list", postController.PageList)


	return r
}
