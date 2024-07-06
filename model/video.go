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

//视频与作者绑定模型
type VideoAuthorBundle struct {
	Basics
	Author        UserAPI `json:"author"`
	PlayUrl       string  `json:"play_url"`       //播放地址
	CoverUrl      string  `json:"cover_url"`      //封面地址
	FavoriteCount int     `json:"favorite_count"` //赞数目
	CommentCount  int     `json:"comment_count"`  //评论数目
	IsFavorited   bool    `json:"is_favorited"`   //true-已点赞，false-未点赞
	Title         string  `json:"title"`          //标题
}

//GetVideoListByAuthorId 通过用户id获取视频列表
func GetVideoListByAuthorId(userId uint64) ([]*Video, error) {
	var videoList []*Video
	err := DB.Where("author_id = ?").Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	return videoList, nil
}

//GetVideoListById 通过视频id列表获取视频列表
func GetVideoListById(videoIdList []uint64) ([]*Video, error) {
	var videoList []*Video
	err := DB.Where("id IN (?)", videoIdList).Find(videoList).Error
	if err != nil {
		return nil, err
	}
	return videoList, nil
}

//Save 保存video记录,如果存在则更新,不存在则创建
func (v *Video) Save() error {
	return DB.Save(&v).Error
}
