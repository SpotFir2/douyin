package controller

import "github.com/gin-gonic/gin"

/*
CommentAction 评论操作
登录用户对视频进行评论
POST /douyin/comment/action/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902509
*/
func CommentAction(c *gin.Context) {
	//token := c.Query("token")              //用户鉴权token
	//videoId := c.Query("video_id")         //视频id
	//actionType := c.Query("action_type")   //1-发布评论，2-删除评论
	//commentText := c.Query("comment_text") //用户填写的评论内容，在action_type=1的时候使用
	//commentId := c.Query("comment_id")     //要删除的评论id，在action_type=2的时候使用
	//TODO 评论操作接口
}

/*
GetCommentList 评论列表
查看视频的所有评论，按发布时间倒序
GET /douyin/comment/list/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902517
*/
func GetCommentList(c *gin.Context) {
	//token := c.Query("token")      //用户鉴权token
	//videoId := c.Query("video_id") //视频id
	//TODO 评论列表接口
}
