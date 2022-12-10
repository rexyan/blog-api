package admin

import "github.com/gogf/gf/v2/frame/g"

type CategoryAndTagReq struct {
	g.Meta        `path:"/categoryAndTag" tags:"Category" method:"get" summary:"CategoryAndTagReq"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type Categories struct {
	ID           int           `json:"id"`
	CategoryName string        `json:"name"`
	Blogs        []interface{} `json:"blogs"`
}
type Tags struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Color string        `json:"color"`
	Blogs []interface{} `json:"blogs"`
}
type CategoryAndTagRes struct {
	Categories []Categories `json:"categories"`
	Tags       []Tags       `json:"tags"`
}
