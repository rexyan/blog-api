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

func (c *cComment) GetCommentTreeList(ctx context.Context, req *admin.GetCommentListReq) (res *admin.GetCommentListRes, err error) {
	commentList, err := service.Comment().GetCommentTreeList(ctx, model.PageInput{
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

func (c *cComment) UpdateComment(ctx context.Context, req *admin.UpdateCommentReq) (res *admin.UpdateCommentRes, err error) {
	err = service.Comment().UpdateComment(ctx, &model.CommentItem{
		ID:       req.ID,
		Nickname: req.Nickname,
		Email:    req.Email,
		Content:  req.Content,
		Avatar:   req.Avatar,
		Website:  req.Website,
		IP:       req.IP,
	})
	return
}

func (c *cComment) DeleteComment(ctx context.Context, req *admin.DeleteCommentReq) (res *admin.DeleteCommentRes, err error) {
	err = service.Comment().DeleteComment(ctx, req.Id)
	return
}
