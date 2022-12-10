package admin

import (
	"blog-api/api/admin"
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
