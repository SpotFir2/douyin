package model

//评论模型
type Comment struct {
	Basics
	VideoId  int64  `json:"video_id"`  //视频id
	AuthorId int64  `json:"author_id"` //发布人id
	Content  string `json:"content"`   //评论内容
}
