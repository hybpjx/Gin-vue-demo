package controller

import (
	"GinDemo/common"
	"GinDemo/model"
	"GinDemo/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(context *gin.Context) {

	db := common.GetDb()

	//	1. 获取参数
	name := context.PostForm("name")
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")
	// 2. 数据验证
	if len(telephone) != 11 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号必须是11位",
		})
		return
	}

	if len(password) < 6 && len(password) > 16 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码必须是6-16位之间",
		})
		return
	}

	// 如果名称没有传值，则传入随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, telephone, password)

	// 3. 判断手机号是否存在

	if isTelephoneExist(db, telephone) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户已经注册了",
		})
		return
	}

	// 4. 创建用户
	newUser := model.User{
		Name:      name,
		Password:  password,
		Telephone: telephone,
	}

	db.Create(&newUser)

	// 5. 返回结果
	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	//fmt.Println(user.ID)
	// 等于0 即 查不到
	if user.ID != 0 {
		return true
	}
	return false
}
