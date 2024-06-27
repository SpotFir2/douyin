package controller

import "github.com/gin-gonic/gin"

/*
GetFavoriteList 喜欢列表
用户的所有点赞视频
GET /douyin/favorite/list/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902464
*/
func GetFavoriteList(c *gin.Context) {
	//userId := c.Query("user_id") //用户id
	//token := c.Query("token")    //用户鉴权token
	//TODO 喜欢列表
}

/*
FavoriteAction 赞操作
登录用户对视频的点赞和取消点赞操作
POST /douyin/favorite/action/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902441
*/
func FavoriteAction(c *gin.Context) {
	//token := c.Query("token")            //用户鉴权token
	//videoId := c.Query("video_id")       //视频id
	//actionType := c.Query("action_type") //1-点赞，2-取消点赞
	//TODO 赞操作
}
