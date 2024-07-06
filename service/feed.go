package service

import (
	"time"

	"github.com/SpotFir2/douyin/model"
	"github.com/SpotFir2/douyin/pkg/serializer"
)

type GetFeedService struct {
	Token      string
	LatestTime string
}

func (f *GetFeedService) Get() *serializer.GetFeedResponse {
	if f.LatestTime == "" { //如果未传入latesttime则将其设置为当前时间
		f.LatestTime = time.Now().Format("2006-01-02 15:04:05")
	}
	time, _ := time.Parse(f.LatestTime, "2006-01-02 15:04:05")
	latestTime := uint64(time.Unix()) //将time.Time格式化为uint64
	var (
		videoList []*model.VideoAuthorBundle
		user      *model.User
		err       error
	)
	if f.Token != "" { //token不为空则获取关注与点赞信息
		if !CheckToken(f.Token) {
			return serializer.GetFeedResponseBuilder(serializer.CodeUserTokenInvalid, 0, nil)
		}
		user, err = GetUserByToken(f.Token)
		if err != nil {
			return serializer.GetFeedResponseBuilder(serializer.CodeUserNotFound, 0, nil)
		} else {
			videoList, err = model.GetFeed(latestTime, user.ID)
		}
	} else { //token为空即用户未登录,关注与点赞恒为false
		videoList, err = model.GetFeed(latestTime)
	}
	if err != nil {
		return serializer.GetFeedResponseBuilder(serializer.CodeFeedNoMore, 0, nil)
	}
	return serializer.GetFeedResponseBuilder(serializer.CodeSuccess, uint64(videoList[len(videoList)-1].CreatedAt.Unix()), videoList)
}
