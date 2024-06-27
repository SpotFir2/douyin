package serializer

import "github.com/SpotFir2/douyin/model"

type GetPublishListResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"` //用户发布的视频列表
}

type PublishActionResponse struct {
	Response
}
