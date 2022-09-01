package common

import (
	"GinDemo/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 1、指定加密秘钥
var jwtKey = []byte("a_secret_crest")

/*
2. 创建Claims结构体
这个结构体就是用来保存信息的，
需要内嵌jwt.StandardClaims，
这些信息会被保存在我们生成好的token当中。
*/
type Claims struct {
	UserID         uint
	StandardClaims jwt.StandardClaims
}

// 3、生成token
func ReleaseToken(user model.User) (string, error) {
	// 有效期 7 天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //设置这个token的有效期
			IssuedAt:  time.Now().Unix(),     //发放时间
			Issuer:    "kaiyuanshinian.tech", // 发行方
			Subject:   "user token",          //主题
		},
	}

	//使用指定的签名方式创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.StandardClaims)

	//使用上面指定的钥匙(secret)签名并获取完整的签名后的字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

//4. 解析token

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims.StandardClaims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjI2MDg1NDEsImlhdCI6MTY2MjAwMzc0MSwiaXNzIjoia2FpeXVhbnNoaW5pYW4udGVjaCIsInN1YiI6InVzZXIgdG9rZW4ifQ.yE8kO0w9yYFUkhmcZRFi2_3ss9yCpT9_IYweWUZ9KjU
