package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type cCategory struct{}

var Category = cCategory{}

func (c *cCategory) CategoryAndTag(ctx context.Context, req *admin.CategoryAndTagReq) (res *admin.CategoryAndTagRes, err error) {
	res = &admin.CategoryAndTagRes{}
	categoryList, err := service.Category().GetCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	for _, category := range categoryList {
		res.Categories = append(res.Categories, admin.Categories{
			ID:           gconv.Int(category.Id),
			CategoryName: category.CategoryName,
			Blogs:        []interface{}{},
		})
	}
	tagList, err := service.Tag().GetTagList(ctx)
	if err != nil {
		return nil, err
	}
	for _, tag := range tagList {
		res.Tags = append(res.Tags, admin.Tags{
			ID:    gconv.Int(tag.Id),
			Name:  tag.TagName,
			Color: tag.Color,
			Blogs: []interface{}{},
		})
	}
	return
}

func (c *cCategory) CreateCategory(ctx context.Context, req *admin.CreateCategoryReq) (res *admin.CreateCategoryRes, err error) {
	err = service.Category().CreateCategory(ctx, req.Name)
	return
}

func (c *cCategory) GetCategoryList(ctx context.Context, req *admin.GetCategoryListReq) (res *admin.GetCategoryListRes, err error) {
	categoryPageList, err := service.Category().GetCategoryPageList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(categoryPageList, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cCategory) UpdateCategory(ctx context.Context, req *admin.UpdateCategoryReq) (res *admin.UpdateCategoryRes, err error) {
	err = service.Category().UpdateCategory(ctx, req.ID, req.CategoryName)
	return
}

func (c *cCategory) DeleteCategory(ctx context.Context, req *admin.DeleteCategoryReq) (res *admin.DeleteCategoryRes, err error) {
	err = service.Category().DeleteCategory(ctx, req.Id)
	return
}
