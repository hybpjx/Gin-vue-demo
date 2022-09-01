package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Userinfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Macke(user *Userinfo) (token string, err error) { //生成jwt
	claims := jwt.MapClaims{ //创建一个自己的声明
		"name": user.Username,
		"pwd":  user.Password,
		"iss":  "lva",
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Second * 4).Unix(),
		"iat":  time.Now().Unix(),
	}

	then := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//fmt.Println(then) //打印&{ 0xc0000040a8 map[alg:HS256 typ:JWT] map[exp:1637212218 iat:1637212214 iss:lvjianhua name:zhansan nbf:1637212214 pwd:pwd]  false}

	token, err = then.SignedString([]byte("gettoken"))

	return
}

func secret() jwt.Keyfunc { //按照这样的规则解析
	return func(t *jwt.Token) (interface{}, error) {
		return []byte("gettoken"), nil
	}
}

//解析token
func ParseToken(token string) (user *Userinfo, err error) {
	user = &Userinfo{}
	tokn, _ := jwt.Parse(token, secret())

	claim, ok := tokn.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("解析错误")
		return
	}
	if !tokn.Valid {
		err = errors.New("令牌错误！")
		return
	}
	//fmt.Println(claim)
	user.Username = claim["name"].(string) //强行转换为string类型
	user.Password = claim["pwd"].(string)  //强行转换为string类型
	return
}

func main() {
	var use = Userinfo{"zic", "admin*123"}
	tkn, _ := Macke(&use)
	fmt.Println("_____", tkn)
	// time.Sleep(time.Second * 8)超过时间打印令牌错误
	user, err := ParseToken(tkn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Username)
}
