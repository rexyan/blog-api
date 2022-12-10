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

// GetBlogCount 获取博客计数
func (s *sBlog) GetBlogCount(ctx context.Context) (count int64) {
	count, err := dao.Blog.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}

// GetBlogCountByTagId 根据 TagId 获取博客计数
func (s *sBlog) GetBlogCountByTagId(ctx context.Context, TagId int64) (count int64) {
	count, err := dao.BlogTag.Ctx(ctx).Count(dao.BlogTag.Columns().TagId, TagId)
	if err != nil {
		return 0
	}
	return
}

// GetBlogCountByCategoryId 根据 CateGoryId 获取博客计数
func (s *sBlog) GetBlogCountByCategoryId(ctx context.Context, CategoryID int64) (count int64) {
	count, err := dao.Blog.Ctx(ctx).Count(dao.Blog.Columns().CategoryId, CategoryID)
	if err != nil {
		return 0
	}
	return
}
