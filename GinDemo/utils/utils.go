package utils

import (
	"math/rand"
	"time"
)

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
