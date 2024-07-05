package service

import (
	"time"

	"github.com/SpotFir2/douyin/model"
	"github.com/SpotFir2/douyin/pkg/config"

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

// GetUserByToken 根据token获取用户
func GetUserByToken(tokenString string) (*model.User, error) {
	token, err := praseToken(tokenString)
	if err != nil {
		return nil, err
	}
	username, err := token.Claims.GetSubject()
	if err != nil {
		return nil, err
	}
	return model.GetUserByUsername(username)
}

// CheckToken 检验token是否有效
// true-有效 false-无效
func CheckToken(tokenString string) bool {
	token, err := praseToken(tokenString)
	return token.Valid && err == nil
}
