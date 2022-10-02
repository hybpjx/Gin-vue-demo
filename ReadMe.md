# 1. 实现用户注册
安装gin
> go get -u github.com/gin-gonic/gin

安装gorm
> go get -u gorm.io/gorm

```go
package main

import (
	"gorm.io/driver/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null"`
	Password  string `gorm:"size:255 not null"`
}

func InitDB() *gorm.DB {
	//driverName := "mysql"
	host := "127.0.0.1"
	port := 3306
	username := "root"
	password := "admin"
	database := "ginInessential"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)


	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
    _ = DB.AutoMigrate(&model.User{})
	fmt.Printf("连接成功：%v\n", db)
	return db
}

func RandomString(n int) string {
	var letters = []byte("ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789")
	result := make([]byte, n)

	//设置随机数种⼦，加上这⾏代码，可以保证每次随机都是随机
	rand.Seed(time.Now().Unix())

	for i := range result {
		//如果不加上述 rand.Seed(time.Now().Unix())每次遍历获取都是重复的一些随机数据
		result[i] = letters[rand.Intn(len(letters))]
	}
	//log.Println(result)
	return string(result)
}
func main() {
	err := InitDB()
	if err != nil {
		return 
	}
	
	r := gin.Default()
	r.POST("/api/register", func(context *gin.Context) {
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
			name = RandomString(10)
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

	})

	_ = r.Run(":9000")
}

```
将功能解耦
分别拆解为以下图
注册返回jwt

github.com/dgrijalva/jwt-go

jwt.go
```go
package common

import (
	"GinDemo/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 1、指定加密秘钥
var jwtKey = []byte("a_secret_crest")

// Claims 2. 创建Claims结构体
type Claims struct {
	/*
		这个结构体就是用来保存信息的，
		需要内嵌jwt.StandardClaims，
		这些信息会被保存在我们生成好的token当中。

		Claim是什么?
		翻译过来就是声明，断言，在payload里面的每一对键值对都是一等一对claim,Claim可以有很多种，JWT官方提供了7个：

		iss (issuer)：签发人
		exp (expiration time)：过期时间
		sub (subject)：主题
		aud (audience)：受众
		nbf (Not Before)：生效时间
		iat (Issued At)：签发时间
		jti (JWT ID)：编号
	*/
	UserID         uint
	StandardClaims jwt.StandardClaims
}

// ReleaseToken 3、生成token
func ReleaseToken(user model.User) (string, error) {
	// 有效期 7 天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //设置这个token的有效期
			IssuedAt:  time.Now().Unix(),     //发放时间
			Issuer:    "zic.tech",            // 发行方
			Subject:   "user token",          //主题
		},
	}

	//使用指定的签名方式创建签名对象
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	////fmt.Println(token)
	////使用上面指定的钥匙(secret)签名并获取完整的签名后的字符串
	//tokenString, err := tokenStruct.SignedString(jwtKey)
	//
	//if err != nil {
	//	return "", err
	//}

	return tokenStruct.SignedString(jwtKey)

}

func (c Claims) Valid() error {
	return nil
}

//// ParseToken 4. 解析token
//func ParseToken(tokenString string) (*Claims,bool) {
//
//	tokenObj,_ := jwt.ParseWithClaims(tokenString,&Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtKey,nil
//	})
//	if key,_ := tokenObj.Claims.(*Claims); tokenObj.Valid {
//		return key,true
//	}else{
//		return nil,false
//	}
//}

// ParseToken 4. 解析token
func ParseToken(tokenString string) (*Claims, bool) {

	tokenObj, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if key, _ := tokenObj.Claims.(*Claims); tokenObj.Valid {
		return key, true
	} else {
		return nil, false
	}
}

```


# 2. 实现登录功能

```go
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
```
# 3. 设置登录后返回用户模块
```go
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
```
路过一个中间件
```go
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

```

# 4. 返回统一的格式
```go
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
func Response(context *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	context.JSON(httpStatus, gin.H{
		"code": code,
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

```
# 5. 从文件中读取配置
[github.com/spf13/viper](https://github.com/spf13/viper)

> go get -u github.com/spf13/viper


# 6. 前端
设置 eslint
安装bootsrtip
[https://code.z01.com/bootstrap-vue/docs/](https://code.z01.com/bootstrap-vue/docs/)
基于模块化
你可以使用Webpack、 Parcel 、 rollup.js等方法引入到项目中，通过yarn 或者 npm 来获取安装Vue.js、 BootstrapVue 和 Bootstrap v4:
**With npm**
> npm install vue bootstrap-vue bootstrap

**With yarn**
> yarn add vue bootstrap-vue bootstrap
> yarn add bootstrap@4

然后在你的应用程序入口点注BootstrapVue：
```javascript

// app.js
import Vue from 'vue'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

// Install BootstrapVue
Vue.use(BootstrapVue)
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)
```
导入
```javascript
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'


createApp(App).use(store).use(router).mount('#app')

```
# 7. 安装 前端验证插件

[https://vuelidate.js.org/#getting-started](https://vuelidate.js.org/#getting-started)

# 8. 安装axios

```python
yarn add axios
yarn add vue-axios
```
或者
```python
npm install axios vue-axios
```
```vue
// // 点击事件的方法
  function register() {
    // 验证数据
    if (!(this.user.telephone)) {
      this.showTelephoneValidate = true;
      return
    } else if (!(this.user.password)) {
      this.showPasswordValidate = true;
      return
    }

    // 请求
    const api = "http://localhost:1016/api/auth/register"
    axios.post(api, { ...this.user }).then(res => {

      // 保存token
      console.log(res.data)
      // localStorage.setItem("token", res.data.data.token)
      storageService.set(storageService.USER_TOKEN, res.data.data.token)
      // 没有this
      $router.replace({ name: "home" })
    }).catch(err => {
      console.log(err)
      const openCenter = () => {
        ElMessage({
          showClose: true,
          message: err.response.data.msg,
          center: true,
          type: 'error',
        })
      }



      openCenter()
      return

    })



    console.log("注册成功")



  }
```
# 9. 登录成功不显示登录和注册

# 10. 封装axios
[https://github.com/axios/axios](https://github.com/axios/axios)
```vue
import storageService from "@/service/storageService";
import axios from "axios";


export default axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
    timeout: 1000,
    headers: { Authorization: `$Bearer ${storageService.get(storageService.USER_TOKEN)}` }
});
```
创建两个服务
storageServices
```vue
// 本地缓存服务

const PREFIX = "ginessential_"

// user 模块
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;


// 储存
const set = (key, data) => {
    localStorage.setItem(key, data)
}

// 读取

const get = (key) => localStorage.getItem(key)

export default {
    set,
    get,
    USER_TOKEN,
    USER_INFO
}
```
userServices.js
```vue
import request from "@/utils/requests";

// 用户注册
const register = ({ name, telephone, password }) => {
    return request.post("auth/register", { name, telephone, password })
}

export default { register }
```

# 11.后端文章分类接口
```go
package controller

import (
	"GinDemo/common"
	"GinDemo/model"
	"GinDemo/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDb()
	_ = db.AutoMigrate(model.Category{})

	return CategoryController{
		DB: *db,
	}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	_ = ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填项")
		return
	}

	if err := c.DB.Create(&requestCategory).Error; err != nil {
		response.Fail(ctx, nil, "分类已存在 请更换name 重新添加")
		return
	}

	response.Success(ctx, gin.H{"category": requestCategory}, "返回成功")
	return
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var deleteCategory model.Category
	category := c.DB.First(&deleteCategory, categoryID)
	// 调用错误方法 如果不为nil 则代表有错误
	err := category.Error
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	category.Delete("id", categoryID)

	response.Success(ctx, nil, "删除成功")
	return
}

func (c CategoryController) Put(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category
	_ = ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填项")
		return
	}
	// 获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category
	category := c.DB.First(&updateCategory, categoryID)

	// 调用错误方法 如果不为nil 则代表有错误
	err := category.Error
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	category.Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
	return
}

func (c CategoryController) Select(ctx *gin.Context) {
	// 获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var SelectCategory model.Category

	category := c.DB.First(&SelectCategory, categoryID)

	// 调用错误方法 如果不为nil 则代表有错误
	err := category.Error
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"SelectCategory": SelectCategory}, "查询成功")
	return
}

```


# 12. go中使用uuid

## 安装
```go
go get github.com/satori/go.uuid
go get github.com/jinzhu/gorm
```
## 使用
```go
package model

import uuid "github.com/satori/go.uuid"

type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36):primary_key"`
	...
}
```