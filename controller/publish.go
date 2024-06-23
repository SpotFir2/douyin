package controller

import "github.com/gin-gonic/gin"

type GetPublishListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"` //用户发布的视频列表
}

type PublishActionResponse struct {
	Response
}

/*
GetPublish 发布列表
用户的视频发布列表，直接列出用户所有投稿过的视频
GET /douyin/publish/list/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18901444
*/
func GetPublish(c *gin.Context) {
	//token := c.Query("token")    //用户鉴权token
	//userId := c.Query("user_id") //用户id
	//TODO 发布列表
}

/*
PublishAction 投稿接口
登录用户选择视频上传
POST /douyin/publish/action/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18875092
*/
func PublishAction(c *gin.Context) {
	//data, err := c.FormFile("data") //视频数据
	//token := c.PostForm("token") //用户鉴权token
	//title := c.PostForm("title") //视频标题
	//TODO 投稿接口
}
