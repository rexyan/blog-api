package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type cComment struct{}

var Comment = cComment{}

func (c *cComment) GetCommentList(ctx context.Context, req *admin.GetCommentListReq) (res *admin.GetCommentListRes, err error) {
	commentList, err := service.Comment().GetCommentList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Structs(commentList, &res)
	return
}

func (c *cComment) UpdateCommentNotice(ctx context.Context, req *admin.UpdateCommentNoticeStatusReq) (res *admin.UpdateCommentNoticeStatusRes, err error) {
	err = service.Comment().UpdateCommentNotice(ctx, req.ID, req.IsNotice)
	return
}

func (c *cComment) UpdateCommentPublish(ctx context.Context, req *admin.UpdateCommentPublishStatusReq) (res *admin.UpdateCommentPublishStatusRes, err error) {
	err = service.Comment().UpdateCommentPublish(ctx, req.ID, req.IsPublished)
	return
}
