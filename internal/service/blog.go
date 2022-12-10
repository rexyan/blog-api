
// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"blog-api/internal/model"
	"context"
)
type(
IBlog interface {
	GetBlogCount(ctx context.Context) (count int64)
	GetBlogCountByTagId(ctx context.Context, TagId int64) (count int64)
	GetBlogCountByCategoryId(ctx context.Context, CategoryID int64) (count int64)
	CreateBlog(ctx context.Context, in model.CreateBlogInput) (err error)
	BlogAssociationTag(ctx context.Context, BlogId int, TagId int) (res int, err error)
	GetBlogs(ctx context.Context, title string, page model.PageInput) (res *model.BlogsOutput, err error)
}
)
var(
localBlog IBlog
)
func Blog() IBlog {
	if localBlog == nil {
		panic("implement not found for interface IBlog, forgot register?")
	}
	return localBlog
}

func RegisterBlog(i IBlog) {
	localBlog = i
}

