package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type cTag struct{}

var Tag = cTag{}

func (c *cTag) CreateTag(ctx context.Context, req *admin.CreateTagReq) (res *admin.CreateTagRes, err error) {
	err = service.Tag().CreateTag(ctx, &model.CreateTagInput{
		Color:   req.Color,
		TagName: req.TagName,
	})
	return
}

func (c *cTag) GetTagList(ctx context.Context, req *admin.GetTagListReq) (res *admin.GetTagListRes, err error) {
	tagList, err := service.Tag().GetTagPageList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(tagList, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cTag) UpdateTag(ctx context.Context, req *admin.UpdateTagReq) (res *admin.UpdateTagRes, err error) {
	err = service.Tag().UpdateTag(ctx, &model.UpdateTagInput{
		ID: req.Id,
		CreateTagInput: model.CreateTagInput{
			Color:   req.Color,
			TagName: req.TagName,
		},
	})
	return
}

func (c *cTag) DeleteTag(ctx context.Context, req *admin.DeleteTagReq) (res *admin.DeleteTagRes, err error) {
	err = service.Tag().DeleteTag(ctx, req.Id)
	return
}
