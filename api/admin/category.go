package admin

import "github.com/gogf/gf/v2/frame/g"

type CategoryAndTagReq struct {
	g.Meta        `path:"/categoryAndTag" tags:"Category" method:"get" summary:"获取 Category 和 Tag"`
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

type CreateCategoryReq struct {
	g.Meta        `path:"/category" tags:"Category" method:"post" summary:"创建 Category"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Name          string `json:"name"`
}

type CreateCategoryRes struct {
}

type GetCategoryListReq struct {
	g.Meta        `path:"/categories" tags:"Category" method:"get" summary:"Category 列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
}

type GetCategoryListRes struct {
	List []Categories `json:"list"`
	CommonPageHelper
}

type UpdateCategoryReq struct {
	g.Meta        `path:"/category" tags:"Category" method:"put" summary:"Category 列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Categories
}

type UpdateCategoryRes struct {
}

type DeleteCategoryReq struct {
	g.Meta        `path:"/category" tags:"Category" method:"delete" summary:"删除 Category"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" dc:"动态ID" in:"query"`
}

type DeleteCategoryRes struct {
}
