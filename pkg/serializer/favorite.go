package serializer

import "github.com/SpotFir2/douyin/model"

type GetFavoriteResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
}

type FavoriteActionResponse struct {
	Response
}
