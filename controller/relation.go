package controller

import "github.com/gin-gonic/gin"

/*
RelationAction 关注操作
POST /douyin/relation/action/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902556
*/
func RelationAction(c *gin.Context) {
	//token := c.Query("token")            //用户鉴权token
	//toUserId := c.Query("to_user_id")    //对方用户id
	//actionType := c.Query("action_type") //1-关注，2-取消关注
	//TODO 关注操作
}

/*
GetFollowList 关注列表
GET /douyin/relation/follow/list/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902568
*/
func GetFollowList(c *gin.Context) {
	//userId := c.Query("user_id") //用户id
	//token := c.Query("token")    //用户鉴权token
	//TODO 关注列表
}

/*
GetFollowerList 粉丝列表
GET /douyin/relation/follower/list/
https://apifox.com/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902563
*/
func GetFollowerList(c *gin.Context) {
	//userId := c.Query("user_id") //用户id
	//token := c.Query("token")    //用户鉴权token
	//TODO 粉丝列表
}
