package admin

import "github.com/gogf/gf/v2/frame/g"

type Comment struct {
	ID              int           `json:"id"`
	Nickname        string        `json:"nickname"`
	Email           string        `json:"email"`
	Content         string        `json:"content"`
	Avatar          string        `json:"avatar"`
	CreateTime      string        `json:"createTime"`
	Website         string        `json:"website"`
	IP              string        `json:"ip"`
	IsPublished     bool          `json:"published"`
	IsAdminComment  bool          `json:"adminComment"`
	Page            int           `json:"page"`
	IsNotice        bool          `json:"notice"`
	ParentCommentID int           `json:"parentCommentId"`
	Qq              interface{}   `json:"qq"`
	BlogId          interface{}   `json:"blog"`
	ReplyComments   []interface{} `json:"replyComments"`
}

type GetCommentListReq struct {
	g.Meta        `path:"/comments" tags:"Comment" method:"get" summary:"Comment 列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
}

type GetCommentListRes struct {
	CommonPageHelper
	List []Comment `json:"list"`
}

type UpdateCommentNoticeStatusReq struct {
	g.Meta        `path:"/comment/notice" tags:"Comment" method:"put" summary:"更新评论提醒状态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	ID            int    `json:"id"`
	IsNotice      bool   `json:"notice"`
}

type UpdateCommentNoticeStatusRes struct {
}

type UpdateCommentPublishStatusReq struct {
	g.Meta        `path:"/comment/published" tags:"Comment" method:"put" summary:"更新评论是否公开状态"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	ID            int    `json:"id"`
	IsPublished   bool   `json:"published"`
}

type UpdateCommentPublishStatusRes struct {
}
