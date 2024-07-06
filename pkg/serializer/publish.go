package serializer

import "github.com/SpotFir2/douyin/model"

type GetPublishListResponse struct {
	Response
	VideoList []*model.VideoAuthorBundle `json:"video_list,omitempty"` //用户发布的视频列表
}

type PublishActionResponse struct {
	Response
}

func GetPublishListResponseBuilder(statusCode int, videoList []*model.VideoAuthorBundle) *GetPublishListResponse {
	res := &GetPublishListResponse{
		Response: NewResponse(statusCode, CodeMsg[statusCode]),
	}
	if statusCode != CodeSuccess {
		return res
	} else {
		res.VideoList = videoList
	}
	return res
}

func PublishActionResponseBuilder(statusCode int, err string) *PublishActionResponse {
	res := &PublishActionResponse{
		Response: NewResponse(statusCode, err),
	}
	return res
}
