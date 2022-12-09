package admin

import "github.com/gogf/gf/v2/frame/g"

// UserLoginReq 用户登录
type UserLoginReq struct {
	g.Meta        `path:"/login" tags:"用户" method:"post" summary:"用户登录"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type UserInfo struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Role       string `json:"role"`
}

type UserLoginRes struct {
	User  UserInfo `json:"user"`
	Token string   `json:"token"`
}
