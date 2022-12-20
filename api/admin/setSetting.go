package admin

import "github.com/gogf/gf/v2/frame/g"

type Type struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}

type GetSiteSettingsReq struct {
	g.Meta        `path:"/siteSettings" tags:"站点设置" method:"get" summary:"站点设置列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type GetSiteSettingsRes struct {
	Type3 []Type `json:"type3"`
	Type2 []Type `json:"type2"`
	Type1 []Type `json:"type1"`
}

type UpdateSetSettingsReq struct {
	g.Meta        `path:"/siteSettings" tags:"站点设置" method:"post" summary:"更新站点设置"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	DeleteIds     []int  `json:"deleteIds"`
	Settings      []Type `json:"settings"`
}

type UpdateSetSettingsRes struct{}
