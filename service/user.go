package service

import (
	"strconv"

	"github.com/SpotFir2/douyin/controller"
	"github.com/SpotFir2/douyin/model"
	"github.com/SpotFir2/douyin/pkg/config"
	"github.com/SpotFir2/douyin/pkg/snowflake"
	"github.com/SpotFir2/douyin/utils"
)

type UserRegisterService struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginService struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserInfoService struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

// Register 用户注册
func (u *UserRegisterService) Register() (*controller.UserRegisterResponse, error) {
	//校验用户名是否存在
	_, err := model.GetUserByUsername(u.Username)
	if err == nil {
		res := &controller.UserRegisterResponse{}
		res.StatusCode = 1
		res.StatusMsg = "用户名已存在"
		return res, nil
	}
	user := &model.User{
		FollowCount:   0,
		FollowerCount: 0,
	}
	user.Name = u.Username
	user.Password = utils.PwdEncrypt(u.Password, config.Confi.Salt)
	user.ID = uint64(snowflake.GenerateId())
	err = user.Save()
	if err != nil {
		res := &controller.UserRegisterResponse{}
		res.StatusCode = 1
		res.StatusMsg = "用户注册失败"
		return res, nil
	}
	token, err := NewToken(user)
	if err != nil {
		res := &controller.UserRegisterResponse{}
		res.StatusCode = 1
		res.StatusMsg = "用户token生成失败"
		return res, nil
	}
	res := &controller.UserRegisterResponse{}
	res.StatusCode = 0
	res.StatusMsg = "用户注册成功"
	res.UserId = user.ID
	res.Token = token
	return res, nil
}

// Login 用户登录
func (u *UserLoginService) Login() (*controller.UserLoginResponse, error) {
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		res := &controller.UserLoginResponse{}
		res.StatusCode = 1
		res.StatusMsg = err.Error()
		return res, nil
	}
	if utils.PwdMatch(u.Username, user.Password, config.Confi.Salt) {
		res := &controller.UserLoginResponse{}
		res.StatusCode = 0
		res.StatusMsg = "用户登录成功"
		res.UserId = user.ID
		token, _ := NewToken(user)
		res.Token = token
		return res, nil
	} else {
		res := &controller.UserLoginResponse{}
		res.StatusCode = 1
		res.StatusMsg = "用户登录失败"
		return res, nil
	}
}

// GetUserInfo 用户信息
func (u *GetUserInfoService) GetUserInfo() (*controller.GetUserInfoResponse, error) {
	userId, _ := strconv.ParseUint(u.UserId, 10, 64)
	user, err := model.GetUserById(userId)
	if err != nil {
		res := &controller.GetUserInfoResponse{}
		res.StatusCode = 1
		res.StatusMsg = "用户名不存在"
	}
	user2, _ := GetUsetByToken(u.Token)
	res := &controller.GetUserInfoResponse{}
	res.StatusCode = 0
	res.StatusMsg = "用户信息获取成功"
	res.User.Id = user.ID
	res.User.Name = user.Name
	res.User.FollowCount = user.FollowCount
	res.User.FollowerCount = user.FollowerCount
	res.User.IsFollow = model.GetRelationByUserId(user2.ID, userId)
	return res, nil
}
