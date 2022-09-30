package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response
// context 上下文
// httpStatus http 状态码
// code 自己定义的状态码
// data 返回的空接口
// msg 返回的信息
func Response(context *gin.Context, httpStatus int, data gin.H, msg string) {
	context.JSON(httpStatus, gin.H{
		"httpStatus": httpStatus,
		"data": data,
		"msg":  msg,
	})
}

func Success(context *gin.Context, data gin.H, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"msg":  msg,
	})
}

func Fail(context *gin.Context, data gin.H, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code": 400,
		"data": data,
		"msg":  msg,
	})
}

func UnprocessableEntity(context *gin.Context, data gin.H, msg string) {
	context.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": 422,
		"data": data,
		"msg":  msg,
	})
}
