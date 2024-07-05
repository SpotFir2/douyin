package service

import (
	"strconv"

	"github.com/SpotFir2/douyin/model"
	"github.com/SpotFir2/douyin/pkg/config"
	"github.com/SpotFir2/douyin/pkg/serializer"
	"github.com/SpotFir2/douyin/pkg/snowflake"
	"github.com/SpotFir2/douyin/utils"
)

type UserRegisterService struct {
	Username string
	Password string
}

type UserLoginService struct {
	Username string
	Password string
}

type GetUserInfoService struct {
	UserId string
	Token  string
}

// Register 用户注册
func (u *UserRegisterService) Register() *serializer.UserRegisterResponse {
	//校验用户名是否存在
	_, err := model.GetUserByUsername(u.Username)
	if err == nil {
		return serializer.UserRegisterResponseBuilder(serializer.CodeUserAlreadyExists, 0, "")
	}
	//校验用户名合法性
	if !utils.CheckUsername(u.Username) {
		return serializer.UserRegisterResponseBuilder(serializer.CodeUserNameInvalid, 0, "")
	}
	//校验密码合法性
	if !utils.CheckPassword(u.Password) {
		return serializer.UserRegisterResponseBuilder(serializer.CodeUserPwdInvalid, 0, "")
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
		return serializer.UserRegisterResponseBuilder(serializer.CodeUserRegisterFailed, 0, "")
	}
	token, err := NewToken(user)
	if err != nil {
		return serializer.UserRegisterResponseBuilder(serializer.CodeUserTokenGenerateFailed, 0, "")
	}
	return serializer.UserRegisterResponseBuilder(0, user.ID, token)
}

// Login 用户登录
func (u *UserLoginService) Login() *serializer.UserLoginResponse {
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		return serializer.UserLoginResponseBuilder(serializer.CodeUserNotFound, 0, "")
	}
	if utils.PwdMatch(u.Username, user.Password, config.Confi.Salt) {
		token, err := NewToken(user)
		if err != nil {
			return serializer.UserLoginResponseBuilder(serializer.CodeUserTokenGenerateFailed, 0, "")
		}
		return serializer.UserLoginResponseBuilder(0, user.ID, token)
	} else {
		return serializer.UserLoginResponseBuilder(serializer.CodeUserLoginFailed, 0, "")
	}
}

// GetUserInfo 用户信息
func (u *GetUserInfoService) GetInfo() *serializer.GetUserInfoResponse {
	userId, _ := strconv.ParseUint(u.UserId, 10, 64)
	user, err := model.GetUserById(userId)
	if err != nil {
		return serializer.GetUserInfoResponseBuilder(serializer.CodeUserNotFound, nil)
	}
	user2, err := GetUserByToken(u.Token)
	if err != nil {
		return serializer.GetUserInfoResponseBuilder(serializer.CodeUserTokenInvalid, nil)
	}
	userAPI := &model.UserAPI{
		ID:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      model.GetRelationByUserId(user2.ID, userId),
	}
	return serializer.GetUserInfoResponseBuilder(0, userAPI)
}
