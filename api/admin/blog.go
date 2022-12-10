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

type CreateBlogRes struct {
}

type GetBlogListReq struct {
	g.Meta        `path:"/blogs" tags:"Blog" method:"get" summary:"文章列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Title         string `json:"title" dc:"文章标题" d:""`
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

type Blogs struct {
	Total             int    `json:"total"`
	List              []Blog `json:"list"`
	PageNum           int    `json:"pageNum"`
	PageSize          int    `json:"pageSize"`
	Size              int    `json:"size"`
	StartRow          int    `json:"startRow"`
	EndRow            int    `json:"endRow"`
	Pages             int    `json:"pages"`
	PrePage           int    `json:"prePage"`
	NextPage          int    `json:"nextPage"`
	IsFirstPage       bool   `json:"isFirstPage"`
	IsLastPage        bool   `json:"isLastPage"`
	HasPreviousPage   bool   `json:"hasPreviousPage"`
	HasNextPage       bool   `json:"hasNextPage"`
	NavigatePages     int    `json:"navigatePages"`
	NavigatepageNums  []int  `json:"navigatepageNums"`
	NavigateFirstPage int    `json:"navigateFirstPage"`
	NavigateLastPage  int    `json:"navigateLastPage"`
}

type GetBlogListRes struct {
	Blogs      Blogs        `json:"blogs"`
	Categories []Categories `json:"categories"`
}
