package serializer

import "github.com/SpotFir2/douyin/model"

type CommentActionResponse struct {
	Response
	Comment model.Comment `json:"comment,omitempty"` //评论成功返回评论内容，不需要重新拉取整个列表
}

type GetCommentListResponse struct {
	Response
	CommentList []model.Comment `json:"comment_list,omitempty"` //评论列表
}
