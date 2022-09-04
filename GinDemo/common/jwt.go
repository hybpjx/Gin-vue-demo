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
