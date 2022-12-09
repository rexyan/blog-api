package admin

import "github.com/gogf/gf/v2/frame/g"

type DashboardReq struct {
	g.Meta        `path:"/dashboard" tags:"Dashboard" method:"get" summary:"Dashboard"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type DashboardRes struct {
	Uv           int           `json:"uv"`
	BlogCount    int           `json:"blogCount"`
	Pv           int           `json:"pv"`
	CityVisitor  []CityVisitor `json:"cityVisitor"`
	Tag          Tag           `json:"tag"`
	Category     Category      `json:"category"`
	VisitRecord  VisitRecord   `json:"visitRecord"`
	CommentCount int           `json:"commentCount"`
}

type CityVisitor struct {
	City string `json:"city"`
	Uv   int    `json:"uv"`
}
type Series struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type Tag struct {
	Legend []string `json:"legend"`
	Series []Series `json:"series"`
}
type Category struct {
	Legend []string `json:"legend"`
	Series []Series `json:"series"`
}
type VisitRecord struct {
	Date []string `json:"date"`
	Uv   []int    `json:"uv"`
	Pv   []int    `json:"pv"`
}
