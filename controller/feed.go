package controller

import (
	"github.com/SpotFir2/douyin/service"
	"github.com/gin-gonic/gin"
)

func InitFeedRouter(router *gin.RouterGroup) {
	router.GET("/", GetFeed)
}

/*
GetFeed 视频流接口
不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
GET /douyin/feed
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145
*/
func GetFeed(c *gin.Context) {
	latestTime := c.Query("latest_time") //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	token := c.Query("token")            //用户登录状态下设置

	getFeedService := &service.GetFeedService{
		Token:      token,
		LatestTime: latestTime,
	}
	getFeedResponse := getFeedService.Get()
	c.JSON(200, getFeedResponse)
}
