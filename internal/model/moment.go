package model

import "github.com/gogf/gf/v2/os/gtime"

type CreateMomentInput struct {
	Content     string
	CreateTime  *gtime.Time
	Likes       string
	IsPublished bool
}

type GetMomentListOutput struct {
	List []struct {
		ID int
		CreateMomentInput
	} `json:"list"`
	CommonPageHelper
}

type MomentDetailOutput struct {
	ID int
	CreateMomentInput
}

type UpdateMomentInput struct {
	ID int
	CreateMomentInput
}
