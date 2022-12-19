package admin

import "github.com/gogf/gf/v2/frame/g"

type Page struct {
	PageNum  int `json:"pageNum" d:"1"`
	PageSize int `json:"pageSize" d:"10"`
}

type CreateBlogReq struct {
	g.Meta         `path:"/blog" tags:"Blog" method:"post" summary:"新增文章"`
	Authorization  string `json:"Authorization" in:"header"  dc:"Authorization"`
	Title          string `json:"title" dc:"文章标题"`
	FirstPicture   string `json:"firstPicture" dc:"文章首图URL"`
	Description    string `json:"description" dc:"文章描述"`
	Content        string `json:"content" dc:"文章正文"`
	Cate           int    `json:"cate" dc:"分类"`
	TagList        []int  `json:"tagList" dc:"标签"`
	Words          string `json:"words" dc:"字数"`
	ReadTime       int    `json:"readTime" dc:"阅读时长(分钟)"`
	Views          int    `json:"views" dc:"浏览次数"`
	Appreciation   bool   `json:"appreciation" dc:"赞赏"`
	Recommend      bool   `json:"recommend" dc:"推荐"`
	CommentEnabled bool   `json:"commentEnabled" dc:"评论"`
	Top            bool   `json:"top" dc:"置顶"`
	Published      bool   `json:"published" dc:"是否公开"`
	Password       string `json:"password" dc:"私密密码"`
}

type CreateBlogRes struct{}

type GetBlogListReq struct {
	g.Meta        `path:"/blogs" tags:"Blog" method:"get" summary:"文章列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Title         string `json:"title" in:"query" dc:"文章标题" d:""`
	Page
}

type Blog struct {
	ID             int           `json:"id"`
	Title          string        `json:"title"`
	FirstPicture   string        `json:"firstPicture"`
	Content        interface{}   `json:"content"`
	Description    interface{}   `json:"description"`
	Published      bool          `json:"published"`
	Recommend      bool          `json:"recommend"`
	Appreciation   bool          `json:"appreciation"`
	CommentEnabled bool          `json:"commentEnabled"`
	Top            bool          `json:"top"`
	CreateTime     string        `json:"createTime"`
	UpdateTime     string        `json:"updateTime"`
	Views          interface{}   `json:"views"`
	Words          interface{}   `json:"words"`
	ReadTime       interface{}   `json:"readTime"`
	Password       string        `json:"password"`
	User           interface{}   `json:"user"`
	Category       Categories    `json:"category"`
	Tags           []interface{} `json:"tags"`
}

type CommonPageHelper struct {
	Total             int   `json:"total"`
	PageNum           int   `json:"pageNum"`
	PageSize          int   `json:"pageSize"`
	Size              int   `json:"size"`
	StartRow          int   `json:"startRow"`
	EndRow            int   `json:"endRow"`
	Pages             int   `json:"pages"`
	PrePage           int   `json:"prePage"`
	NextPage          int   `json:"nextPage"`
	IsFirstPage       bool  `json:"isFirstPage"`
	IsLastPage        bool  `json:"isLastPage"`
	HasPreviousPage   bool  `json:"hasPreviousPage"`
	HasNextPage       bool  `json:"hasNextPage"`
	NavigatePages     int   `json:"navigatePages"`
	NavigatepageNums  []int `json:"navigatepageNums"`
	NavigateFirstPage int   `json:"navigateFirstPage"`
	NavigateLastPage  int   `json:"navigateLastPage"`
}

type Blogs struct {
	List []Blog `json:"list"`
	CommonPageHelper
}

type GetBlogListRes struct {
	Blogs      Blogs        `json:"blogs"`
	Categories []Categories `json:"categories"`
}

type BlogDetailReq struct {
	g.Meta        `path:"/blog" tags:"Blog" method:"get" summary:"文章详情"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" dc:"文章ID" in:"query"`
}
type BlogDetailRes struct {
	Tags []Tags `json:"tags"`
	Blog
}

type UpdateBlogReq struct {
	g.Meta        `path:"/blog" tags:"Blog" method:"put" summary:"修改文章"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" dc:"文章ID"`
	CreateBlogReq
}

type UpdateBlogRes struct{}

type BlogPlacedTopReq struct {
	g.Meta        `path:"/blog/top" tags:"Blog" method:"put" summary:"文章置顶"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" dc:"文章ID"`
	Top           bool   `json:"top" dc:"是否置顶"`
}
type BlogPlacedTopRes struct{}

type BlogRecommendReq struct {
	g.Meta        `path:"/blog/recommend" tags:"Blog" method:"put" summary:"文章推荐"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" dc:"文章ID"`
	Recommend     bool   `json:"top" dc:"是否置顶"`
}
type BlogRecommendRes struct{}

type BlogVisibilityReq struct {
	g.Meta         `path:"/blog/{Id}/visibility" tags:"Blog" method:"put" summary:"文章可见性"`
	Authorization  string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id             int    `json:"id" dc:"文章ID" in:"path"`
	Appreciation   bool   `json:"appreciation" dc:"赞赏"`
	CommentEnabled bool   `json:"commentEnabled" dc:"评论"`
	Password       string `json:"password" dc:"私密密码"`
	Published      bool   `json:"published" dc:"是否公开"`
	Recommend      bool   `json:"recommend" dc:"推荐"`
	Top            bool   `json:"top" dc:"置顶"`
}
type BlogVisibilityRes struct{}

type DeleteBlogReq struct {
	g.Meta        `path:"/blog" tags:"Blog" method:"delete" summary:"文章可见性"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id" dc:"文章ID" in:"query"`
}

type DeleteBlogRes struct{}

type IdAndTitleReq struct {
	g.Meta        `path:"/blogIdAndTitle" tags:"Blog" method:"get" summary:"ID 和 Title"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type IdAndTitleRes struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
