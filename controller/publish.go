package controller

import (
	"github.com/SpotFir2/douyin/pkg/serializer"
	"github.com/SpotFir2/douyin/service"
	"github.com/gin-gonic/gin"
)

func InitPublishRouter(router *gin.RouterGroup) {
	router.GET("list/", GetPublishList)
	router.POST("action/", PublishAction)
}

/*
GetPublishList 发布列表
用户的视频发布列表，直接列出用户所有投稿过的视频
GET /douyin/publish/list/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18901444
*/
func GetPublishList(c *gin.Context) {
	token := c.Query("token")    //用户鉴权token
	userId := c.Query("user_id") //用户id

	getPublishListService := &service.GetPublishListService{
		Token:  token,
		UserId: userId,
	}
	getPublishListResponse := getPublishListService.GetList()
	c.JSON(200, getPublishListResponse)
}

/*
PublishAction 投稿接口
登录用户选择视频上传
POST /douyin/publish/action/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18875092
*/
func PublishAction(c *gin.Context) {
	data, err := c.FormFile("data") //视频数据
	if err != nil {
		c.JSON(200, serializer.PublishActionResponseBuilder(serializer.CodePublishActionFailed, err.Error()))
	}
	token := c.PostForm("token") //用户鉴权token
	title := c.PostForm("title") //视频标题
	//TODO 投稿接口
	publishActionService := &service.PublishActionService{
		Token: token,
		Title: title,
		File:  data,
	}
	publishActionResponse := publishActionService.Action()
	c.JSON(200, publishActionResponse)
}
