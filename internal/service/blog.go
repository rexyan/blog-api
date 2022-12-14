// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"context"
)

type IBlog interface {
	GetBlogCount(ctx context.Context) (count int64)
	GetBlogCountByTagId(ctx context.Context, TagId int64) (count int64)
	GetBlogCountByCategoryId(ctx context.Context, CategoryID int64) (count int64)
	CreateBlog(ctx context.Context, in model.CreateBlogInput) (err error)
	BlogAssociationTag(ctx context.Context, BlogId int, TagId int) (res int, err error)
	GetBlogs(ctx context.Context, title string, page model.PageInput) (res *model.BlogListOutput, err error)
	GetBlogDetail(ctx context.Context, BlogId int) (res *model.BlogDetailOutput, err error)
	UpdateBlog(ctx context.Context, in model.UpdateBlogInput) (err error)
	UpdateBlogTop(ctx context.Context, in *model.UpdateBlogTopInput) (err error)
	UpdateBlogRecommend(ctx context.Context, in *model.UpdateBlogRecommendInput) (err error)
	UpdateBlogVisibility(ctx context.Context, in *model.UpdateBlogVisibilityInput) (err error)
	DeleteBlog(ctx context.Context, BlogId int) (err error)
	GetBlogIdAndTitle(ctx context.Context) (res []entity.Blog, err error)
}

var localBlog IBlog

func Blog() IBlog {
	if localBlog == nil {
		panic("implement not found for interface IBlog, forgot register?")
	}
	return localBlog
}

func RegisterBlog(i IBlog) {
	localBlog = i
}
