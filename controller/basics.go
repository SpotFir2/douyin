package controller

//用户信息
type User struct {
	Id            uint64 `json:"id"`             //用户id
	Name          string `json:"name"`           //用户名称
	FollowCount   int    `json:"follow_count"`   // 关注总数
	FollowerCount int    `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

//视频信息
type Video struct {
	Id            uint64 `json:"id"`             //视频id
	Author        User   `json:"author"`         //视频作者
	PlayUrl       string `json:"play_url"`       //播放地址
	CoverUrl      string `json:"cover_url"`      //封面地址
	FavoriteCount int    `json:"favorite_count"` //点赞总数
	CommentCount  int    `json:"comment_count"`  //评论总数
	IsFavorite    bool   `json:"is_favorite"`    //true-已点赞，false-未点赞
	Title         string `json:"title"`          //视频标题
}

//评论信息
type Comment struct {
	Id         uint64 `json:"id"`          //用户id
	User       User   `json:"user"`        //评论用户信息
	Content    string `json:"content"`     //评论内容
	CreateDate string `json:"create_date"` //评论发布日期，格式 mm-dd
}

//基本响应信息
type Response struct {
	StatusCode uint64 `json:"status_code"`          //状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` //返回状态描述
}
