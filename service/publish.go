package service

import (
	"mime/multipart"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/SpotFir2/douyin/model"
	"github.com/SpotFir2/douyin/pkg/config"
	"github.com/SpotFir2/douyin/pkg/serializer"
	"github.com/SpotFir2/douyin/pkg/snowflake"
	"github.com/gin-gonic/gin"
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

// GetList 获取用户发布视频列表
func (p *GetPublishListService) GetList() *serializer.GetPublishListResponse {
	var (
		videoList []*model.VideoAuthorBundle
		err       error
		user      *model.User
	)
	intUserId, _ := strconv.Atoi(p.UserId)
	if p.Token != "" {
		if !CheckToken(p.Token) {
			return serializer.GetPublishListResponseBuilder(serializer.CodeUserTokenInvalid, nil)
		}
		user, err = GetUserByToken(p.Token)
		if err != nil {
			return serializer.GetPublishListResponseBuilder(serializer.CodeUserNotFound, nil)
		} else {
			videoList, err = model.GetPublishListByUserId(uint64(intUserId), user.ID)
		}
	} else {
		videoList, err = model.GetPublishListByUserId(uint64(intUserId))
	}
	if err != nil {
		return serializer.GetPublishListResponseBuilder(serializer.CodePublishGetFailed, nil)
	}
	return serializer.GetPublishListResponseBuilder(serializer.CodeSuccess, videoList)
}

// Action 用户视频投稿
func (p *PublishActionService) Action() *serializer.PublishActionResponse {
	//校验token合法
	if !CheckToken(p.Token) {
		return serializer.PublishActionResponseBuilder(serializer.CodeUserTokenInvalid, serializer.CodeMsg[serializer.CodeUserTokenInvalid])
	}
	user, err := GetUserByToken(p.Token)
	if err != nil {
		return serializer.PublishActionResponseBuilder(serializer.CodeUserNotFound, serializer.CodeMsg[serializer.CodeUserNotFound])
	}
	videoPath := "uploads/" + strconv.FormatUint(user.ID, 10) + "/" + time.Now().Format("2006-01-02")
	//创建文件夹用于存储文件
	err = os.MkdirAll(videoPath, 0750)
	if err != nil && !os.IsExist(err) {
		return serializer.PublishActionResponseBuilder(serializer.CodePublishActionFailed, err.Error())
	}
	ctx := &gin.Context{}
	videoPath += "/" + p.File.Filename
	//接收上传文件并存储于文件夹中
	err = ctx.SaveUploadedFile(p.File, videoPath)
	if err != nil {
		return serializer.PublishActionResponseBuilder(serializer.CodePublishActionFailed, err.Error())
	}
	videoUrl := config.Confi.Site.Domain + "/" + videoPath
	coverPath := videoPath + ".jpg"
	//使用ffmpeg截取视频第一帧用于封面,command参数构建防止os注入
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-frames:v", "1", coverPath)
	err = cmd.Run()
	if err != nil {
		return serializer.PublishActionResponseBuilder(serializer.CodePublishActionFailed, err.Error())
	}
	coverUrl := config.Confi.Site.Domain + "/" + coverPath
	videoBasics := &model.Basics{
		ID: uint64(snowflake.GenerateId()),
	}
	video := &model.Video{
		Basics:        *videoBasics,
		AuthorId:      user.ID,
		PlayUrl:       videoUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         p.Title,
	}
	//在数据库中存储视频记录
	err = video.Save()
	if err != nil {
		return serializer.PublishActionResponseBuilder(serializer.CodePublishActionFailed, err.Error())
	}
	return serializer.PublishActionResponseBuilder(serializer.CodeSuccess, serializer.CodeMsg[serializer.CodeSuccess])
}
