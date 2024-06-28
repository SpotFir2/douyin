package controller

import (
	// TODO 拆分响应体与controller 尝试序列化响应结构
	"github.com/SpotFir2/douyin/service"

	"github.com/gin-gonic/gin"
)

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

	userRegisterService := service.UserRegisterService{
		Username: username,
		Password: password,
	}
	userRegisterResponse := userRegisterService.Register()
	c.JSON(200, userRegisterResponse)
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

	userLoginService := service.UserLoginService{
		Username: username,
		Password: password,
	}
	userLoginResponse := userLoginService.Login()
	c.JSON(200, userLoginResponse)
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

	getUserInfoService := service.GetUserInfoService{
		UserId: userId,
		Token:  token,
	}
	getUserInfoResponse := getUserInfoService.GetUserInfo()
	c.JSON(200, getUserInfoResponse)
}
