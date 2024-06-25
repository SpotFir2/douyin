package model

//User 用户信息模型
type User struct {
	Basics
	Name          string `json:"name"`           //用户名
	Password      string `json:"password"`       //密码
	FollowCount   int    `json:"follow_count"`   //关注数目
	FollowerCount int    `json:"follower_count"` //粉丝数目
}

type UserAPI struct {
	ID            uint64 `json:"id"`             //用户id
	Name          string `json:"name"`           //用户名
	FollowCount   int    `json:"follow_count"`   //关注数目
	FollowerCount int    `json:"follower_count"` //粉丝数目
	IsFollow      bool   `json:"is_follow"`      //true-已关注，false-未关注
}

//GetUserByUsername 根据username获取user
func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	err := DB.First(user, "name=?", username).Error
	return user, err
}

//GetUserByUserid 根据userid获取user
func GetUserById(userid uint64) (*UserAPI, error) {
	user := &User{}
	err := DB.First(user, "id=?", userid).Error
	userapi := &UserAPI{
		ID:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	return userapi, err
}

//Save 在数据库中存储user
func (user *User) Save() error {
	return DB.Save(user).Error
}
