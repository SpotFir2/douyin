package service

import (
	"douyin/model"
	"douyin/pkg/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 生成token
func NewToken(user *model.User) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "douyin",
		"exp": time.Now().Add(time.Hour * 24 * 30),
		"sub": user.Name,
	})
	tokenString, err = token.SignedString(config.Confi.Salt)
	return
}

// PraseToken 解析token
func praseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.Confi.Salt, nil
	})
}

// GetUsetByToken 根据token获取用户
func GetUsetByToken(tokenString string) (*model.User, error) {
	token, err := praseToken(tokenString)
	if err != nil {
		return nil, errors.New("token无效")
	}
	username, err := token.Claims.GetSubject()
	if err != nil {
		return nil, errors.New("token非法")
	}
	return model.GetUserByUsername(username)
}

// CheckToken 检验token是否有效
func CheckToken(tokenString string) bool {
	token, err := praseToken(tokenString)
	return token.Valid && err == nil
}
