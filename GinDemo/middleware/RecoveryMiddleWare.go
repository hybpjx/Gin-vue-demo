package middleware

import (
	"GinDemo/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(context, nil, fmt.Sprint("系统错误",err))
			}
		}()

		context.Next()
	}

}
