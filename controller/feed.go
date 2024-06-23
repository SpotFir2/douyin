package controller

import "github.com/gin-gonic/gin"

type GetFeedResponse struct {
	Response
	NextTime  uint64  `json:"next_time,omitempty"`  //本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []Video `json:"video_list,omitempty"` //视频列表
}

/*
GetFeed 视频流接口
不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
GET /douyin/feed
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145
*/
func GetFeed(c *gin.Context) {
	//latestTime := c.Query("latest_time") //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	//token := c.Query("token")            //用户登录状态下设置
	//TODO 视频流
}
