package model

//评论模型
type Comment struct {
	Basics
	VideoId     uint64 `json:"video_id"`     //视频id
	AuthorId    uint64 `json:"author_id"`    //发布人id
	CommentText string `json:"comment_text"` //评论内容
}
