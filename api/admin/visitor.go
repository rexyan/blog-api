package admin

import "github.com/gogf/gf/v2/frame/g"

type Visitor struct {
	ID         int         `json:"id"`
	UUID       string      `json:"uuid"`
	IP         string      `json:"ip"`
	IPSource   string      `json:"ipSource"`
	Os         string      `json:"os"`
	Browser    string      `json:"browser"`
	CreateTime string      `json:"createTime"`
	LastTime   string      `json:"lastTime"`
	Pv         int         `json:"pv"`
	UserAgent  interface{} `json:"userAgent"`
}

type GetVisitorListReq struct {
	g.Meta        `path:"/visitors" tags:"访客" method:"get" summary:"访客列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
}

type GetVisitorListRes struct {
	CommonPageHelper
	List []Visitor `json:"list"`
}

type DeleteVisitorReq struct {
	g.Meta        `path:"/visitor" tags:"访客" method:"delete" summary:"删除访客记录"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	VisitorId     int    `json:"id"`
}

type DeleteVisitorRes struct {
}
