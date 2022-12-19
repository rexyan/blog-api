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

type Moment struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	CreateTime  string `json:"createTime"`
	Likes       int    `json:"likes"`
	IsPublished bool   `json:"published"`
}

type GetMomentListReq struct {
	g.Meta        `path:"/moments" tags:"Moment" method:"get" summary:"动态列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
}

type GetMomentListRes struct {
	List []Moment `json:"list"`
	CommonPageHelper
}

type MomentDetailReq struct {
	g.Meta        `path:"/moment" tags:"Moment" method:"get" summary:"动态详情"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" dc:"动态ID" in:"query"`
}

type MomentDetailRes struct {
	Moment
}

type UpdateMomentReq struct {
	g.Meta        `path:"/moment" tags:"Moment" method:"put" summary:"修改动态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Moment
}

type UpdateMomentRes struct {
	Moment
}

type UpdateMomentPublishedReq struct {
	g.Meta        `path:"/moment/published" tags:"Moment" method:"put" summary:"修改动态发布状态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" dc:"动态ID"`
	Published     bool   `json:"published"`
}

type UpdateMomentPublishedRes struct {
}

type DeleteMomentReq struct {
	g.Meta        `path:"/moment" tags:"Moment" method:"delete" summary:"删除动态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" dc:"动态ID" in:"query"`
}

type DeleteMomentRes struct {
}
