package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/logic/middleware"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type cBlog struct{}

var Blog = cBlog{}

func (c *cBlog) CreateBlog(ctx context.Context, req *admin.CreateBlogReq) (res *admin.CreateBlogRes, err error) {
	userId := gconv.String(middleware.Auth().GetIdentity(ctx))
	err = service.Blog().CreateBlog(ctx, model.CreateBlogInput{
		Title:            req.Title,
		FirstPicture:     req.FirstPicture,
		Content:          req.Content,
		Description:      req.Description,
		IsPublished:      req.Published,
		IsRecommend:      req.Recommend,
		IsAppreciation:   req.Appreciation,
		IsCommentEnabled: req.CommentEnabled,
		Views:            req.Views,
		Words:            req.Words,
		ReadTime:         req.ReadTime,
		CategoryId:       req.Cate,
		IsTop:            req.Top,
		Password:         req.Password,
		UserId:           userId,
		TagList:          req.TagList,
	})
	return
}

func (c *cBlog) GetBlogList(ctx context.Context, req *admin.GetBlogListReq) (res *admin.GetBlogListRes, err error) {
	res = &admin.GetBlogListRes{}
	categoryList, err := service.Category().GetCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	blogs, err := service.Blog().GetBlogs(ctx, req.Title, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(blogs, &res.Blogs)
	for _, category := range categoryList {
		res.Categories = append(res.Categories, admin.Categories{
			ID:           gconv.Int(category.Id),
			CategoryName: category.CategoryName,
			Blogs:        []interface{}{},
		})
	}
	return
}

func (c *cBlog) GetBlogDetail(ctx context.Context, req *admin.BlogDetailReq) (res *admin.BlogDetailRes, err error) {
	blog, err := service.Blog().GetBlogDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(blog, &res)
	return
}

func (c *cBlog) UpdateBlog(ctx context.Context, req *admin.UpdateBlogReq) (res *admin.UpdateBlogRes, err error) {
	userId := gconv.String(middleware.Auth().GetIdentity(ctx))
	err = service.Blog().UpdateBlog(ctx, model.UpdateBlogInput{
		Id: req.Id,
		CreateBlogInput: model.CreateBlogInput{
			Title:            req.Title,
			FirstPicture:     req.FirstPicture,
			Content:          req.Content,
			Description:      req.Description,
			IsPublished:      req.Published,
			IsRecommend:      req.Recommend,
			IsAppreciation:   req.Appreciation,
			IsCommentEnabled: req.CommentEnabled,
			Views:            req.Views,
			Words:            req.Words,
			ReadTime:         req.ReadTime,
			CategoryId:       req.Cate,
			IsTop:            req.Top,
			Password:         req.Password,
			UserId:           userId,
			TagList:          req.TagList,
		},
	})
	return
}

func (c *cBlog) UpdateBlogTop(ctx context.Context, req *admin.BlogPlacedTopReq) (res *admin.BlogPlacedTopRes, err error) {
	err = service.Blog().UpdateBlogTop(ctx, &model.UpdateBlogTopInput{
		Id:  req.Id,
		Top: req.Top,
	})
	return
}

func (c *cBlog) UpdateBlogRecommend(ctx context.Context, req *admin.BlogRecommendReq) (res *admin.BlogRecommendRes, err error) {
	err = service.Blog().UpdateBlogRecommend(ctx, &model.UpdateBlogRecommendInput{
		Id:        req.Id,
		Recommend: req.Recommend,
	})
	return
}

func (c *cBlog) UpdateBlogVisibility(ctx context.Context, req *admin.BlogVisibilityReq) (res *admin.BlogVisibilityRes, err error) {
	err = service.Blog().UpdateBlogVisibility(ctx, &model.UpdateBlogVisibilityInput{
		Id:             req.Id,
		Appreciation:   req.Appreciation,
		CommentEnabled: req.CommentEnabled,
		Password:       req.Password,
		Published:      req.Published,
		Recommend:      req.Recommend,
		Top:            req.Top,
	})
	return
}

func (c *cBlog) DeleteBlog(ctx context.Context, req *admin.DeleteBlogReq) (res *admin.DeleteBlogRes, err error) {
	err = service.Blog().DeleteBlog(ctx, req.Id)
	return
}

func (c *cBlog) GetBlogIdAndTitle(ctx context.Context, req *admin.IdAndTitleReq) (res []*admin.IdAndTitleRes, err error) {
	IdAndTitleList, err := service.Blog().GetBlogIdAndTitle(ctx)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(IdAndTitleList, &res)
	if err != nil {
		return nil, err
	}
	return
}
