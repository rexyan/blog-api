package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateMomentReq struct {
	g.Meta        `path:"/moment" tags:"Moment" method:"post" summary:"新增动态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Content       string `json:"content"`
	CreateTime    string `json:"createTime"`
	Likes         string `json:"likes"`
	Published     bool   `json:"published"`
}

type CreateMomentRes struct{}
