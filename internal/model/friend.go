package model

import "blog-api/internal/model/entity"

type GetFriendInfoOutput struct {
	CommentEnabled bool
	Content        string
}

type GetFriendListOutput struct {
	CommonPageHelper
	List []*entity.Friend
}

type UpdateFriendInput struct {
	entity.Friend
}
