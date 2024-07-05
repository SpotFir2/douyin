package service

import (
	"mime/multipart"
	"strconv"

	"github.com/SpotFir2/douyin/model"
	"github.com/SpotFir2/douyin/pkg/serializer"
)

type GetPublishListService struct {
	Token  string
	UserId string
}

type PublishActionService struct {
	Token string
	Title string
	File  *multipart.FileHeader
}

func (p *GetPublishListService) GetList() *serializer.GetPublishListResponse {
	if !CheckToken(p.Token) {
		return serializer.GetPublishListResponseBuilder(serializer.CodeUserTokenInvalid, nil)
	}
	user, err := GetUserByToken(p.Token)
	intUserId, _ := strconv.Atoi(p.UserId)
	var (
		videoList []*model.VideoAuthorBundle
		err1      error
	)
	if err != nil {
		videoList, err1 = model.GetPublishListByUserId(uint64(intUserId))
	} else {
		videoList, err1 = model.GetPublishListByUserId(uint64(intUserId), user.ID)

	}
	if err1 != nil {
		return serializer.GetPublishListResponseBuilder(serializer.CodePublishGetFailed, nil)
	}
	return serializer.GetPublishListResponseBuilder(serializer.CodeSuccess, videoList)
}
