package controller

import "github.com/gin-gonic/gin"

type RestController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Put(ctx *gin.Context)
	Select(ctx *gin.Context)
}