package blog

import (
	"blog-api/internal/dao"
	"blog-api/internal/service"
	"context"
)

type sBlog struct{}

func init() {
	service.RegisterBlog(New())
}
func New() *sBlog {
	return &sBlog{}
}

func (s *sBlog) GetBlogCount(ctx context.Context) (count int64) {
	count, err := dao.Blog.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}
