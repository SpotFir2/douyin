package model

//Video 视频信息模型
type Video struct {
	Basics
	AuthorId      uint64 `json:"author_id"`      //作者id
	PlayUrl       string `json:"play_url"`       //播放地址
	CoverUrl      string `json:"cover_url"`      //封面地址
	FavoriteCount int    `json:"favorite_count"` //赞数目
	CommentCount  int    `json:"comment_count"`  //评论数目
	Title         string `json:"title"`          //标题
}
