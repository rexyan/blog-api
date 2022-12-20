package admin

import "github.com/gogf/gf/v2/frame/g"

type GetFriendInfoReq struct {
	g.Meta        `path:"/friendInfo" tags:"友链" method:"get" summary:"友链信息"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type GetFriendInfoRes struct {
	CommentEnabled bool   `json:"commentEnabled"`
	Content        string `json:"content"`
}

type GetFriendListReq struct {
	g.Meta        `path:"/friends" tags:"友链" method:"get" summary:"友链列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
}

type Friend struct {
	ID          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Avatar      string `json:"avatar"`
	IsPublished bool   `json:"published"`
	Views       int    `json:"views"`
	CreateTime  string `json:"createTime"`
}

type GetFriendListRes struct {
	CommonPageHelper
	List []Friend `json:"list"`
}

type UpdateFriendPublishStatusReq struct {
	g.Meta        `path:"/friend/published" tags:"友链" method:"put" summary:"修改友链公开状态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	ID            int    `json:"id"`
	IsPublished   bool   `json:"published"`
}

type UpdateFriendPublishStatusRes struct {
}

type UpdateFriendReq struct {
	g.Meta        `path:"/friend" tags:"友链" method:"put" summary:"修改友链"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Friend
}

type UpdateFriendRes struct {
}

type UpdateFriendInfoReq struct {
	g.Meta        `path:"/friendInfo/content" tags:"友链" method:"put" summary:"修改友链页面信息"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Content       string `json:"content"`
}

type UpdateFriendInfoRes struct {
}
