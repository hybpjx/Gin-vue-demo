package controller

import (
	"GinDemo/common"
	"GinDemo/model"
	"GinDemo/response"
	"GinDemo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Register 用户注册模块
func Register(context *gin.Context) {

	db := common.GetDb()

	// 使用绑定结构体的方法 获取请求参数


	var requestUser = model.User{}
	context.Bind(&requestUser)
	//json.ne
	fmt.Println(requestUser)
	//	1. 获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 2. 数据验证
	if len(telephone) != 11 {
		response.UnprocessableEntity(context, nil, "手机号必须是11位")
		return
	}

	if len(password) < 6 && len(password) > 16 {
		response.UnprocessableEntity(context, nil, "密码必须是6-16位之间")
		return
	}

	// 如果名称没有传值，则传入随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, telephone, password)

	// 3. 判断手机号是否存在

	if isTelephoneExist(db, telephone) {
		response.UnprocessableEntity(context, nil, "用户已经注册了")

		return
	}

	// 4. 创建用户

	// 密码是不能明文保存的
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {

		response.Response(context, http.StatusInternalServerError, 500, nil, "密码加密错误")
		return
	}

	newUser := model.User{
		Name:      name,
		Password:  string(hasedPassword),
		Telephone: telephone,
	}

	db.Create(&newUser)

	// 发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")

		log.Fatalf("token generate error: %v\n", err)
		return
	}

	// 5. 返回结果
	response.Success(
		context,
		gin.H{
			"token": token,
		},
		"注册成功",
	)

}

// Login 用户登陆模块
func Login(context *gin.Context) {

	DB := common.GetDb()

	//	1. 获取参数
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")
	// 2. 数据验证
	if len(telephone) != 11 {
		response.UnprocessableEntity(context, nil, "手机号必须是11位")

		return
	}

	if len(password) < 6 && len(password) > 16 {

		response.UnprocessableEntity(context, nil, "密码必须是6-16位之间")

		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)

	// 等于0 即 查不到 所以等于0 才会返回不存在
	if user.ID == 0 {

		response.UnprocessableEntity(context, nil, "用户不存在")

		return
	}

	// 判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {

		response.Response(context, http.StatusBadRequest, 400, nil, "密码错误")

		return
	}

	// 发送token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")

		log.Fatalf("token generate error: %v\n", err)
		return
	}

	// 5. 返回结果
	response.Success(
		context,
		gin.H{
			"token": token,
		},
		"登录成功",
	)

}

// Info 登录后 获取token 获得用户信息
func Info(context *gin.Context) {
	//获得所有信息 UserInfo 是一个结构体
	userInfo, _ := context.Get("userInfo")

	response.Success(
		context, gin.H{
			"userInfo": userInfo,
		},
		"返回成功",
	)

	/*
		name, _ := context.Get("name")
		telephone, _ := context.Get("telephone")
		context.JSON(http.StatusOK, gin.H{
			"code":      200,
			"name":      name,
			"telephone": telephone,
		})
	*/

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
