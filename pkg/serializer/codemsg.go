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

const (
	CodePublishTitleInvalid = 2000 + iota
	CodePublishGetFailed
	CodePublishActionFailed
)

const (
	CodeFeedNoMore = 3000 + iota
	CodeFeedGetFailed
)

const (
	CodeFavoriteLikeFailed = 4000 + iota
)

const (
	CodeCommentFaild = 5000 + iota
)

const (
	CodeRelationLikeFailed = 6000 + iota
)

var CodeMsg = map[int]string{
	CodeSuccess:                 "",
	CodeUserAlreadyExists:       "用户已存在",
	CodeUserRegisterFailed:      "用户注册失败",
	CodeUserNameInvalid:         "用户名不合法",
	CodeUserPwdInvalid:          "密码不合法",
	CodeUserLoginFailed:         "用户名或密码错误",
	CodeUserIdInvalid:           "用户id不合法",
	CodeUserTokenInvalid:        "用户token不合法",
	CodeUserTokenGenerateFailed: "用户token生成失败",
	CodePublishTitleInvalid:     "视频标题不合法",
	CodePublishGetFailed:        "获取发布列表失败",
	CodePublishActionFailed:     "视频上传失败",
	CodeFeedNoMore:              "没有更多视频",
	CodeFeedGetFailed:           "无法获取视频",
	CodeFavoriteLikeFailed:      "点赞失败",
	CodeCommentFaild:            "评论失败",
	CodeRelationLikeFailed:      "关注失败",
}
