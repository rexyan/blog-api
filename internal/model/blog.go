package model

import "blog-api/internal/model/entity"

type PageInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type CreateBlogInput struct {
	Title            string // 文章标题
	FirstPicture     string // 文章首图，用于随机文章展示
	Content          string // 文章正文
	Description      string // 描述
	IsPublished      bool   // 公开或私密
	IsRecommend      bool   // 推荐开关
	IsAppreciation   bool   // 赞赏开关
	IsCommentEnabled bool   // 评论开关
	Views            int    // 浏览次数
	Words            string // 文章字数
	ReadTime         int    // 阅读时长(分钟)
	CategoryId       int    // 文章分类
	IsTop            bool   // 是否置顶
	Password         string // 密码保护
	UserId           string // 文章作者
	TagList          []int  // 标签列表
}

type BlogListOutPutItems struct {
	ID             int         `json:"id"`
	Title          string      `json:"title"`
	FirstPicture   string      `json:"firstPicture"`
	Content        interface{} `json:"content"`
	Description    interface{} `json:"description"`
	Published      bool        `json:"published"`
	Recommend      bool        `json:"recommend"`
	Appreciation   bool        `json:"appreciation"`
	CommentEnabled bool        `json:"commentEnabled"`
	Top            bool        `json:"top"`
	CreateTime     string      `json:"createTime"`
	UpdateTime     string      `json:"updateTime"`
	Views          interface{} `json:"views"`
	Words          interface{} `json:"words"`
	ReadTime       interface{} `json:"readTime"`
	Password       string      `json:"password"`
	User           interface{} `json:"user"`
	CategoryId     int64       `json:"categoryId"`
	Category       struct {
		ID           int           `json:"id"`
		CategoryName string        `json:"name"`
		Blogs        []interface{} `json:"blogs"`
	} `json:"category"`
	Tags []interface{} `json:"tags"`
}

type BlogListOutput struct {
	Total             int                   `json:"total"`
	List              []BlogListOutPutItems `json:"list"`
	PageNum           int                   `json:"pageNum"`
	PageSize          int                   `json:"pageSize"`
	Size              int                   `json:"size"`
	StartRow          int                   `json:"startRow"`
	EndRow            int                   `json:"endRow"`
	Pages             int                   `json:"pages"`
	PrePage           int                   `json:"prePage"`
	NextPage          int                   `json:"nextPage"`
	IsFirstPage       bool                  `json:"isFirstPage"`
	IsLastPage        bool                  `json:"isLastPage"`
	HasPreviousPage   bool                  `json:"hasPreviousPage"`
	HasNextPage       bool                  `json:"hasNextPage"`
	NavigatePages     int                   `json:"navigatePages"`
	NavigatepageNums  []int                 `json:"navigatepageNums"`
	NavigateFirstPage int                   `json:"navigateFirstPage"`
	NavigateLastPage  int                   `json:"navigateLastPage"`
}

type BlogDetailOutput struct {
	BlogListOutPutItems
	Tags []entity.Tag `json:"tags"`
}
