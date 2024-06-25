package model

//relation 关注关系模型
type Relation struct {
	Basics
	UserId     uint64 `json:"user_id"`     //用户id
	FollowedId uint64 `json:"followed_id"` //被关注用户id
}

//GetRelationByUserId 获取关注信息
//userId-用户id FollowedId-被关注用户id
//true-已关注 false-未关注
func GetRelationByUserId(userId uint64, FollowerId uint64) bool {
	relation := &Relation{}
	err := DB.First(relation, "user_id=? AND followed_id=?", userId, FollowerId)
	return err == nil
}
