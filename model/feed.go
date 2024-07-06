package model

func GetFeed(latestTime uint64, userId ...uint64) ([]*VideoAuthorBundle, error) {
	var (
		res []*VideoAuthorBundle
		err error
	)
	if len(userId) != 0 { //用户登陆时需返回关注与点赞信息
		err = DB.
			Table("video v").
			Select("v.id as id,u.id as author_id,u.name as author_name,u.follow_count as author_follow_count,u.follower_count as author_follower_count,IF(r.id IS NULL,FALSE,TRUE) as author_is_follow,v.play_url as play_url,v.cover_url as cover_url,v.favorite_count as favorite_count,v.comment_count as comment_count,IF(f.id IS NULL,FALSE,TRUE) as is_favorite,v.title as title").
			Joins("LEFT JOIN user u on v.author_id=u.id").
			Joins("LEFT JOIN relation r ON v.author_id=r.followed_id AND r.user_id = ?", userId).
			Joins("LEFT JOIN favorite f ON v.id=f.video_id AND f.user_id = ?", userId).
			Where("WHERE v.created_at < ?", latestTime).
			Order("v.created_at").
			Limit(30).
			Scan(&res).
			Error
	} else { //用户未登录时关注与点赞默认为false
		err = DB.
			Table("video v").
			Select("v.id as id,u.id as author_id,u.name as author_name,u.follow_count as author_follow_count,u.follower_count as author_follower_count,FALSE as author_is_follow,v.play_url as play_url,v.cover_url as cover_url,v.favorite_count as favorite_count,v.comment_count as comment_count,FALSE as is_favorite,v.title as title").
			Joins("LEFT JOIN user u on v.author_id=u.id").
			Where("v.created_at < ?", latestTime).
			Order("v.created_at").
			Limit(30).
			Scan(&res).
			Error
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
