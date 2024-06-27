package serializer

import "github.com/SpotFir2/douyin/model"

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
	User model.UserAPI `json:"user,omitempty"` //用户信息
}
