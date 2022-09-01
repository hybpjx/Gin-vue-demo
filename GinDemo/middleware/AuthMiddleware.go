package middleware

import (
	"GinDemo/common"
	"GinDemo/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取 authorization headers
		tokenString := context.GetHeader("Authorization")
		// 如果token 开头是空，或者开头不是以bearer结尾的 则报错
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限验证错误",
			})
			context.Abort()
			return
		}

		tokenString = tokenString[7:]

		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		fmt.Println(token)
		fmt.Println(token.Valid)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token错误,或者token已失效",
			})
			context.Abort()

			return
		}

		// 验证后获取 clams 中的UserID,
		userID := claims.UserID
		DB := common.GetDb()
		var user model.User
		DB.First(&user, userID)

		// 如果用户不存在
		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "用户不存在",
			})
			context.Abort()
			return
		}

		// 如果用户存在 则有效
		context.Set("user", user)
		context.Next()
	}

}
