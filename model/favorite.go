package model

//favorite 视频点赞模型
type Favorite struct {
	Basics
	UserId  uint64 `json:"user_id"`  //用户id
	VideoId uint64 `json:"video_id"` //用户赞的视频id
}
