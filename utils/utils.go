package utils

import (
	"log"

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
*/
func PwdMatch(str1, str2, salt string) bool {
	str1 = str1 + salt
	if err := bcrypt.CompareHashAndPassword([]byte(str2), []byte(str1)); err != nil {
		return false
	}
	return true
}
