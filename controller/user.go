package controller

import (
	"github.com/SpotFir2/douyin/service"

	"github.com/gin-gonic/gin"
)

// 用户注册响应
type UserRegisterResponse struct {
	Response
	UserId uint64 `json:"user_id"` //用户id
	Token  string `json:"token"`   //用户鉴权token
}

// 用户登录响应
type UserLoginResponse struct {
	Response
	UserId uint64 `json:"user_id,omitempty"` //用户id
	Token  string `json:"token,omitempty"`   //用户鉴权token
}

// 用户信息响应
type GetUserInfoResponse struct {
	Response
	User User `json:"user,omitempty"` //用户信息
}

func InitUserRoute(router *gin.RouterGroup) {
	router.POST("register/", UserRegister)
	router.POST("login/", UserLogin)
	router.GET("/", GetUserInfo)
}

/*
UserRegister 用户注册
新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一
创建成功后返回用户 id 和权限token
POST /douyin/user/register/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18899952
*/
func UserRegister(c *gin.Context) {
	username := c.Query("username") //注册用户名，最长32个字符
	password := c.Query("password") //密码，最长32个字符
	//TODO 用户注册

	userRegisterService := service.UserRegisterService{
		Username: username,
		Password: password,
	}
	userRegisterResponse, _ := userRegisterService.Register()
	c.JSON(200, userRegisterResponse)
	return
}

/*
UserLogin 用户登录
通过用户名和密码进行登录
登录成功后返回用户 id 和权限 token
POST /douyin/user/login/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18900033
*/
func UserLogin(c *gin.Context) {
	username := c.Query("username") //注册用户名，最长32个字符
	password := c.Query("password") //密码，最长32个字符
	//TODO 用户登录

	userLoginService := service.UserLoginService{
		Username: username,
		Password: password,
	}
	userLoginResponse, _ := userLoginService.Login()
	c.JSON(200, userLoginResponse)
	return
}

/*
GetUserInfo 用户信息
获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数
GET /douyin/user/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18901232
*/
func GetUserInfo(c *gin.Context) {
	userId := c.Query("user_id") //用户id
	token := c.Query("token")    //用户鉴权token
	//TODO 用户信息

	getUserInfoService := service.GetUserInfoService{
		UserId: userId,
		Token:  token,
	}
	getUserInfoResponse, _ := getUserInfoService.GetUserInfo()
	c.JSON(200, getUserInfoResponse)
}
