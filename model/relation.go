package model

//relation 关注关系模型
type Relation struct {
	Basics
	UserId     uint64 `json:"user_id"`     //用户id
	FollowedId uint16 `json:"followed_id"` //被关注用户id
}
