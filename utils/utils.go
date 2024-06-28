package utils

import (
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

/*
PwdEncrypt 密码加盐加密
str-密码明文 salt-盐值
*/
func PwdEncrypt(str, salt string) string {
	str = str + salt
	result, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(result)
}

/*
PwdMatch 密码匹配
str1-输入密码明文 str2-数据库密码密文 salt-盐值
true-成功 false-失败
*/
func PwdMatch(str1, str2, salt string) bool {
	str1 = str1 + salt
	if err := bcrypt.CompareHashAndPassword([]byte(str2), []byte(str1)); err != nil {
		return false
	}
	return true
}

/*
CheckUsername 校验用户名
如果用户名长度大于32则不合法
true-成功 false-失败
*/
func CheckUsername(username string) bool {
	return strings.Count(username, "") <= 32
}

/*
CheckPassword 校验用户名
如果用户名长度大于32则不合法
true-成功 false-失败
*/
func CheckPassword(password string) bool {
	return strings.Count(password, "") <= 32
}
