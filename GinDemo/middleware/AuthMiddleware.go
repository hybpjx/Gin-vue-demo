package middleware

import (
	"GinDemo/common"
	"GinDemo/dto"
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
		// 如果token 开头是空，或者开头不是以bearer结尾的 则报错
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(context,http.StatusUnauthorized,401,nil,"权限验证错误")

			context.Abort()
			return
		}

		//token格式错误
		tokenSlice := strings.SplitN(tokenString, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			response.UnprocessableEntity(context,nil,"token格式错误")

			context.Abort() //阻止执行
			return
		}
		//验证token
		claims, ok := common.ParseToken(tokenSlice[1])
		if !ok {
			response.UnprocessableEntity(context,nil,"token不正确")
			context.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
			response.UnprocessableEntity(context,nil,"token过期")
			context.Abort() //阻止执行
			return
		}

		/*
			//返回所有数据
			context.Set("userInfo", tokenStruck)
			context.Next()

			// 返回部分数据
			var user model.User
			db := common.GetDb()
			db.First(&user, tokenStruck.UserID)

			// 如果用户不存在
			if user.ID == 0 {
				context.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "用户不存在",
				})
				context.Abort()
				return
			}

			// 返回部分数据
			context.Set("name", user.Name)
			context.Set("telephone", user.Telephone)
			context.Next()
		*/

		userInfo := dto.ToUserDto(claims, context)
		// 返回部分数据
		context.Set("userInfo", userInfo)
		context.Next()

	}

}
