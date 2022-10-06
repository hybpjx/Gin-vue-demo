//package dto
//
//import (
//	"GinDemo/common"
//	"GinDemo/model"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//type UserDto struct {
//	Name      string `json:"name"`
//	Telephone string `json:"telephone"`
//}
//
//func ToUserDto(claims *common.Claims, context *gin.Context) UserDto {
//	var user model.User
//	db := common.GetDb()
//	db.First(&user, claims.UserID)
//
//	// 如果用户不存在
//	if user.ID == 0 {
//		context.JSON(http.StatusUnauthorized, gin.H{
//			"code": 401,
//			"msg":  "用户不存在",
//		})
//		context.Abort()
//		return UserDto{
//			Name:      "无此人",
//			Telephone: "无次手机号",
//		}
//	}
//
//	return UserDto{
//		Name:      user.Name,
//		Telephone: user.Telephone,
//	}
//}

package dto

import (
	"GinDemo/model"
)

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
