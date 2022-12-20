package model

type CommentItem struct {
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
	ReplyComments   []CommentItem `json:"replyComments"`
}

type GetCommentListOutput struct {
	CommonPageHelper
	List []CommentItem
}
