package serializer

const (
	CodeUserNotFound = 1000 + iota
	CodeUserAlreadyExists
	CodeUserRegisterFailed
	CodeUserNameInvalid
	CodeUserPwdInvalid
	CodeUserLoginFailed
	CodeUserIdInvalid
	CodeUserTokenInvalid
	CodeUserTokenGenerateFailed
)

var CodeUserMsg = map[int]string{
	CodeSuccess:                 "",
	CodeUserAlreadyExists:       "用户已存在",
	CodeUserRegisterFailed:      "用户注册失败",
	CodeUserNameInvalid:         "用户名不合法",
	CodeUserPwdInvalid:          "密码不合法",
	CodeUserLoginFailed:         "用户名或密码错误",
	CodeUserIdInvalid:           "用户id不合法",
	CodeUserTokenInvalid:        "用户token不合法",
	CodeUserTokenGenerateFailed: "用户token生成失败",
}

const (
	CodePublishCommonFailed = 2000 + iota
	CodePublishTitleInvalid
)

var CodePublishMsg = map[int]string{
	CodeSuccess:             "",
	CodePublishTitleInvalid: "用户token不合法",
}

const (
	CodeFeedNoMore = 3000 + iota
	CodeFeedGetFailed
)

var CodeFeedMsg = map[int]string{
	CodeSuccess:       "",
	CodeFeedNoMore:    "没有更多视频",
	CodeFeedGetFailed: "无法获取视频",
}

const (
	CodeFavoriteLikeFailed = 4000 + iota
)

var CodeFavoriteMsg = map[int]string{
	CodeSuccess:            "",
	CodeFavoriteLikeFailed: "点赞失败",
}

const (
	CodeCommentFaild = 5000 + iota
)

var CodeCommentMsg = map[int]string{
	CodeSuccess:      "",
	CodeCommentFaild: "评论失败",
}

const (
	CodeRelationLikeFailed = 6000 + iota
)

var CodeRelationMsg = map[int]string{
	CodeSuccess:            "",
	CodeRelationLikeFailed: "关注失败",
}
