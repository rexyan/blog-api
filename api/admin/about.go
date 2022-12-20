package admin

import "github.com/gogf/gf/v2/frame/g"

type AboutItem struct {
	MusicID        string `json:"musicId"`
	CommentEnabled string `json:"commentEnabled"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}

type GetAboutReq struct {
	g.Meta        `path:"/about" tags:"关于" method:"get" summary:"关于我"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type GetAboutRes struct {
	AboutItem
}

type UpdateAboutReq struct {
	g.Meta        `path:"/about" tags:"关于" method:"put" summary:"修改关于我"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	AboutItem
}

type UpdateAboutRes struct {
}
