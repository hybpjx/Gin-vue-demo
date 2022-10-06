package middleware

import (
	"GinDemo/common"
	"GinDemo/model"
	"GinDemo/response"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取 authorization headers
		tokenString := context.GetHeader("Authorization")

		//fmt.Printf("目前的tokenString：%v\n", tokenString)

		// 如果token 开头是空，或者开头不是以bearer结尾的 则报错
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(context, http.StatusUnauthorized, nil, "权限验证错误")

			context.Abort()
			return
		}

		//token格式错误
		tokenSlice := strings.SplitN(tokenString, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			response.UnprocessableEntity(context, nil, "token格式错误")

			context.Abort() //阻止执行
			return
		}
		//验证token
		claims, ok := common.ParseToken(tokenSlice[1])
		if !ok {
			response.UnprocessableEntity(context, nil, "token不正确")
			context.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
			response.UnprocessableEntity(context, nil, "token过期")
			context.Abort() //阻止执行
			return
		}

		//token通过验证, 获取claims中的UserID
		userId := claims.UserID
		DB := common.GetDb()
		var user model.User
		DB.First(&user, userId)

		// 验证用户是否存在
		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			context.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		context.Set("userInfo", user)

		context.Next()

	}

}
