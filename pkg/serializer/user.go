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
	UserAPI model.UserAPI `json:"user,omitempty"` //用户信息
}

func UserRegisterResponseBuilder(statusCode int, userId uint64, token string) *UserRegisterResponse {
	res := &UserRegisterResponse{}
	res.Response = NewResponse(statusCode, CodeUserMsg[statusCode])
	if statusCode != CodeSuccess {
		return res
	} else {
		res.UserId = userId
		res.Token = token
	}
	return res
}

func UserLoginResponseBuilder(statusCode int, userId uint64, token string) *UserLoginResponse {
	res := &UserLoginResponse{}
	res.Response = NewResponse(statusCode, CodeUserMsg[statusCode])
	if statusCode != CodeSuccess {
		return res
	} else {
		res.UserId = userId
		res.Token = token
		return res
	}
}

func GetUserInfoResponseBuilder(statusCode int, userAPI *model.UserAPI) *GetUserInfoResponse {
	res := &GetUserInfoResponse{}
	res.Response = NewResponse(statusCode, CodeUserMsg[statusCode])
	if statusCode != CodeSuccess {
		return res
	} else {
		res.UserAPI = model.UserAPI{
			ID:            userAPI.ID,
			Name:          userAPI.Name,
			FollowCount:   userAPI.FollowCount,
			FollowerCount: userAPI.FollowerCount,
			IsFollow:      userAPI.IsFollow,
		}
		return res
	}
}
