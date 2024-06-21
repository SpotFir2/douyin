package model

//User 用户信息模型
type User struct {
	Basics
	Name          string `json:"name"`           //用户名
	Password      string `json:"password"`       //密码
	FollowCount   int    `json:"follow_count"`   //关注数目
	FollowerCount int    `json:"follower_count"` //粉丝数目
}
