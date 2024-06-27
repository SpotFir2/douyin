package serializer

import "github.com/SpotFir2/douyin/model"

type RelationActionResponse struct {
	Response
}

type GetFollowListResponse struct {
	Response
	UserList []model.User `json:"user_list,omitempty"` //用户列表
}

type GetFollowerListResponse struct {
	Response
	UserList []model.User `json:"user_list,omitempty"` //用户列表
}
