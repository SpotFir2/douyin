package model

//GetPublishListByUserId
//oUserId-对象用户ID userId-发起用户ID
func GetPublishListByUserId(oUserId uint64, userId ...uint64) ([]*VideoAuthorBundle, error) {
	var (
		res []*VideoAuthorBundle
		err error
	)

	if len(userId) == 0 { //仅输入对象用户ID,无发起用户ID时,默认未关注未点赞
		err = DB.
			Table("video v").
			Select("v.ID AS id,u.id AS author_id,u.name AS author_name,u.follow_count AS author_follow_count,u.follower_count AS author_follower_count,FALSE AS author_is_follow,v.play_url AS play_url,v.cover_url AS cover_url,v.favorite_count AS favorite_count,v.comment_count AS comment_count,FALSE AS is_favorite,v.title AS title").Joins("LEFT JOIN users u ON v.author_id=u.id").
			Where("author_id = ?", oUserId).
			Order("v.created_at").
			Scan(&res).
			Error
	} else { //输入对象用户ID与发起用户ID时,在relation与favorite表中查询对应关注与点赞记录并返回
		err = DB.
			Table("video v").
			Select("v.id as id,u.id as author_id,u.name as author_name,u.follow_count as author_follow_count,u.follower_count as author_follower_count,IF(r.id IS NULL,FALSE,TRUE) as author_is_follow,v.play_url as play_url,v.cover_url as cover_url,v.favorite_count as favorite_count,v.comment_count as comment_count,IF(f.id IS NULL,FALSE,TRUE) as is_favorite,v.title as title").
			Joins("LEFT JOIN user u ON v.author_id=u.id").
			Joins("LEFT JOIN relation r ON v.author_id=r.followed_id AND r.user_id = ?", userId).
			Joins("LEFT JOIN favorite f ON v.id=f.video_id AND f.user_id = ?", userId).
			Where("v.author_id = ?", oUserId).
			Order("v.created_at").
			Scan(&res).
			Error
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
