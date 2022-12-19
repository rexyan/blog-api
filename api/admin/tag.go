package admin

import "github.com/gogf/gf/v2/frame/g"

type TagItem struct {
	Color   string `json:"color"`
	TagName string `json:"name"`
}
type CreateTagReq struct {
	g.Meta        `path:"/tag" tags:"Tag" method:"post" summary:"新增 Tag"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	TagItem
}

type CreateTagRes struct{}

type GetTagListReq struct {
	g.Meta        `path:"/tags" tags:"Tag" method:"get" summary:"Tag 列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
}

type GetTagListRes struct {
	List []struct {
		Id    int           `json:"id"`
		Blogs []interface{} `json:"blogs"`
		TagItem
	} `json:"list"`
	CommonPageHelper
}

type UpdateTagReq struct {
	g.Meta        `path:"/tag" tags:"Tag" method:"put" summary:"更新 Tag"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id"`
	TagItem
}

type UpdateTagRes struct {
}

type DeleteTagReq struct {
	g.Meta        `path:"/tag" tags:"Tag" method:"delete" summary:"删除 Tag"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" in:"query"`
}

type DeleteTagRes struct {
}
