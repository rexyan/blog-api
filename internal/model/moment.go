package model

import "github.com/gogf/gf/v2/os/gtime"

type CreateMomentInput struct {
	Content    string
	CreateTime *gtime.Time
	Likes      string
	Published  bool
}
