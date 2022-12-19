package model

import "blog-api/internal/model/entity"

type PageInput struct {
	PageNum  int
	PageSize int
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
	ID               int
	Title            string
	FirstPicture     string
	Content          interface{}
	Description      interface{}
	IsPublished      bool
	IsRecommend      bool
	IsAppreciation   bool
	IsCommentEnabled bool
	IsTop            bool
	CreateTime       string
	UpdateTime       string
	Views            interface{}
	Words            interface{}
	ReadTime         interface{}
	Password         string
	User             interface{}
	CategoryId       int64
	Category         struct {
		ID           int
		CategoryName string
		Blogs        []interface{}
	}
	Tags []interface{}
}

type CommonPageHelper struct {
	Total             int
	PageNum           int
	PageSize          int
	Size              int
	StartRow          int
	EndRow            int
	Pages             int
	PrePage           int
	NextPage          int
	IsFirstPage       bool
	IsLastPage        bool
	HasPreviousPage   bool
	HasNextPage       bool
	NavigatePages     int
	NavigatepageNums  []int
	NavigateFirstPage int
	NavigateLastPage  int
}

type BlogListOutput struct {
	List []BlogListOutPutItems
	CommonPageHelper
}

type BlogDetailOutput struct {
	BlogListOutPutItems
	Tags []entity.Tag
}

type UpdateBlogInput struct {
	Id int // 文章Id
	CreateBlogInput
}

type UpdateBlogTopInput struct {
	Id  int // 文章Id
	Top bool
}

type UpdateBlogRecommendInput struct {
	Id        int // 文章Id
	Recommend bool
}

type UpdateBlogVisibilityInput struct {
	Id             int
	Appreciation   bool
	CommentEnabled bool
	Password       string
	Published      bool
	Recommend      bool
	Top            bool
}
