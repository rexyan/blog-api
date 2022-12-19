package model

import "blog-api/internal/model/entity"

type CommentItem struct {
	entity.Comment
	ReplyComments []CommentItem
}

type GetCommentListOutput struct {
	CommonPageHelper
	List []CommentItem
}
